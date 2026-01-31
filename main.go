package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	// Color scheme
	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#04B575")).
		Bold(true).
		PaddingLeft(1)

	prayerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		PaddingLeft(2)

	nextPrayerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true).
		PaddingLeft(2)

	timeStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#50C878")).
		Bold(true)

	cityStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#87CEEB")).
		Bold(true)

	countdownStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF6B6B")).
		Bold(true).
		Align(lipgloss.Center)

	emojiStyle = lipgloss.NewStyle().
		PaddingRight(1)
)

// Prayer time data structures
type PrayerTimesResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	Timings Timings `json:"timings"`
	Date    Date    `json:"date"`
	Meta    Meta    `json:"meta"`
}

type Timings struct {
	Fajr    string `json:"Fajr"`
	Sunrise string `json:"Sunrise"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Sunset  string `json:"Sunset"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

type Date struct {
	Readable string `json:"readable"`
	Hijri    Hijri  `json:"hijri"`
}

type Hijri struct {
	Date    string `json:"date"`
	Format  string `json:"format"`
	Day     string `json:"day"`
	Weekday Weekday `json:"weekday"`
	Month   Month  `json:"month"`
	Year    string `json:"year"`
}

type Weekday struct {
	En string `json:"en"`
	Ar string `json:"ar"`
}

type Month struct {
	Number int    `json:"number"`
	En     string `json:"en"`
	Ar     string `json:"ar"`
}

type Meta struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Method    Method  `json:"method"`
}

type Method struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Params map[string]interface{} `json:"params"`
	Location Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Prayer names with emojis
var prayerNames = map[string]string{
	"Fajr":    "🌅 Fajr",
	"Sunrise": "☀️  Sunrise",
	"Dhuhr":   "🌞 Dhuhr", 
	"Asr":     "🌤️  Asr",
	"Maghrib": "🌅 Maghrib",
	"Isha":    "🌙 Isha",
}

// Prayer order for iteration
var prayerOrder = []string{"Fajr", "Sunrise", "Dhuhr", "Asr", "Maghrib", "Isha"}

func main() {
	var city string
	var method int

	var rootCmd = &cobra.Command{
		Use:   "pray",
		Short: "🕌 Prayer times in your terminal",
		Long:  "A beautiful CLI tool to display Islamic prayer times with accurate calculations based on your location.",
		Run: func(cmd *cobra.Command, args []string) {
			showPrayerTimes(city, method)
		},
	}

	var nextCmd = &cobra.Command{
		Use:   "next",
		Short: "Show the next prayer time with countdown",
		Run: func(cmd *cobra.Command, args []string) {
			showNextPrayer(city, method)
		},
	}

	rootCmd.AddCommand(nextCmd)
	
	rootCmd.PersistentFlags().StringVar(&city, "city", "Riyadh", "City name for prayer times")
	rootCmd.PersistentFlags().IntVar(&method, "method", 4, "Calculation method (4 = Umm Al-Qura)")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func fetchPrayerTimes(city string, method int) (*PrayerTimesResponse, error) {
	url := fmt.Sprintf("http://api.aladhan.com/v1/timingsByCity?city=%s&country=SA&method=%d", city, method)
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prayer times: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d for URL: %s", resp.StatusCode, url)
	}

	var prayerData PrayerTimesResponse
	if err := json.NewDecoder(resp.Body).Decode(&prayerData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &prayerData, nil
}

func parseTime(timeStr string) (time.Time, error) {
	// Remove timezone info if present
	timeStr = strings.Split(timeStr, " ")[0]
	
	today := time.Now()
	parsed, err := time.Parse("15:04", timeStr)
	if err != nil {
		return time.Time{}, err
	}
	
	// Set the date to today
	return time.Date(today.Year(), today.Month(), today.Day(), 
		parsed.Hour(), parsed.Minute(), 0, 0, today.Location()), nil
}

func findNextPrayer(timings Timings) (string, time.Time, error) {
	now := time.Now()
	
	prayerTimes := map[string]string{
		"Fajr":    timings.Fajr,
		"Dhuhr":   timings.Dhuhr,
		"Asr":     timings.Asr,
		"Maghrib": timings.Maghrib,
		"Isha":    timings.Isha,
	}

	for _, prayer := range []string{"Fajr", "Dhuhr", "Asr", "Maghrib", "Isha"} {
		prayerTime, err := parseTime(prayerTimes[prayer])
		if err != nil {
			continue
		}
		
		if now.Before(prayerTime) {
			return prayer, prayerTime, nil
		}
	}
	
	// If no prayer found today, return tomorrow's Fajr
	tomorrow := now.AddDate(0, 0, 1)
	fajrTime, err := parseTime(timings.Fajr)
	if err != nil {
		return "", time.Time{}, err
	}
	
	tomorrowFajr := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(),
		fajrTime.Hour(), fajrTime.Minute(), 0, 0, tomorrow.Location())
	
	return "Fajr", tomorrowFajr, nil
}

func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}

func showPrayerTimes(city string, method int) {
	data, err := fetchPrayerTimes(city, method)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Header
	header := titleStyle.Render(fmt.Sprintf("🕌 Prayer Times for %s", cityStyle.Render(city)))
	dateInfo := fmt.Sprintf("📅 %s | %s %s, %s AH", 
		data.Data.Date.Readable,
		data.Data.Date.Hijri.Day,
		data.Data.Date.Hijri.Month.En,
		data.Data.Date.Hijri.Year)
	
	fmt.Println(header)
	fmt.Println(strings.Repeat("━", 50))
	fmt.Println(cityStyle.Render(dateInfo))
	fmt.Println()

	// Find next prayer
	nextPrayer, nextTime, err := findNextPrayer(data.Data.Timings)
	var nextPrayerName string
	if err == nil {
		nextPrayerName = nextPrayer
	}

	// Display prayers
	timings := map[string]string{
		"Fajr":    data.Data.Timings.Fajr,
		"Sunrise": data.Data.Timings.Sunrise,
		"Dhuhr":   data.Data.Timings.Dhuhr,
		"Asr":     data.Data.Timings.Asr,
		"Maghrib": data.Data.Timings.Maghrib,
		"Isha":    data.Data.Timings.Isha,
	}

	for _, prayer := range prayerOrder {
		timeStr := strings.Split(timings[prayer], " ")[0] // Remove timezone
		prayerName := prayerNames[prayer]
		
		if prayer == nextPrayerName && prayer != "Sunrise" {
			line := fmt.Sprintf("%s %s", emojiStyle.Render("▶"), nextPrayerStyle.Render(fmt.Sprintf("%-15s %s", prayerName, timeStyle.Render(timeStr))))
			fmt.Println(line)
		} else {
			line := fmt.Sprintf("  %s %s", prayerStyle.Render(fmt.Sprintf("%-15s", prayerName)), timeStyle.Render(timeStr))
			fmt.Println(line)
		}
	}

	// Show countdown to next prayer
	if err == nil && nextPrayerName != "Sunrise" {
		duration := time.Until(nextTime)
		if duration > 0 {
			fmt.Println()
			countdown := fmt.Sprintf("⏰ Next prayer in %s", formatDuration(duration))
			fmt.Println(countdownStyle.Render(countdown))
		}
	}

	// Footer with method info
	fmt.Println()
	fmt.Println(strings.Repeat("━", 50))
	methodInfo := fmt.Sprintf("📍 Method: %s", data.Data.Meta.Method.Name)
	fmt.Println(prayerStyle.Render(methodInfo))
}

func showNextPrayer(city string, method int) {
	data, err := fetchPrayerTimes(city, method)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	nextPrayer, nextTime, err := findNextPrayer(data.Data.Timings)
	if err != nil {
		fmt.Printf("Error finding next prayer: %v\n", err)
		os.Exit(1)
	}

	// Skip sunrise for prayer notifications
	if nextPrayer == "Sunrise" {
		// Find the prayer after sunrise
		now := time.Now()
		timings := map[string]string{
			"Dhuhr":   data.Data.Timings.Dhuhr,
			"Asr":     data.Data.Timings.Asr,
			"Maghrib": data.Data.Timings.Maghrib,
			"Isha":    data.Data.Timings.Isha,
		}

		for _, prayer := range []string{"Dhuhr", "Asr", "Maghrib", "Isha"} {
			prayerTime, parseErr := parseTime(timings[prayer])
			if parseErr != nil {
				continue
			}
			
			if now.Before(prayerTime) {
				nextPrayer = prayer
				nextTime = prayerTime
				break
			}
		}
	}

	duration := time.Until(nextTime)
	
	// Header
	fmt.Println(titleStyle.Render("🕌 Next Prayer"))
	fmt.Println(strings.Repeat("━", 30))
	fmt.Println()

	// Prayer info
	prayerName := prayerNames[nextPrayer]
	timeStr := nextTime.Format("15:04")
	
	fmt.Println(nextPrayerStyle.Render(fmt.Sprintf("%s at %s", prayerName, timeStyle.Render(timeStr))))
	fmt.Println()
	
	// Countdown
	if duration > 0 {
		countdown := fmt.Sprintf("⏰ In %s", formatDuration(duration))
		fmt.Println(countdownStyle.Render(countdown))
	} else {
		fmt.Println(countdownStyle.Render("🔔 Prayer time has arrived!"))
	}
	
	fmt.Println()
	fmt.Println(cityStyle.Render(fmt.Sprintf("📍 %s", city)))
}