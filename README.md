# 🕌 Pray - Prayer Times CLI

A beautiful, production-quality Go CLI for Islamic prayer times with accurate calculations and gorgeous terminal output.

## ✨ Features

- 🌍 **Location-based** - Accurate prayer times for any city worldwide
- 🎨 **Beautiful output** - Colorful terminal interface with emojis and clean formatting
- ⏰ **Next prayer countdown** - Shows time remaining until the next prayer
- 🔧 **Configurable** - Different cities and calculation methods
- 📅 **Hijri calendar** - Shows both Gregorian and Islamic dates
- 🚀 **Single binary** - Zero runtime dependencies
- 🎯 **Saudi default** - Optimized for Saudi Arabia with Umm Al-Qura method

## 📦 Installation

### From Source
```bash
git clone https://github.com/isIbra/pray.git
cd pray
go build -o pray
```

### Download Binary
Download the latest release from [GitHub Releases](https://github.com/isIbra/pray/releases).

## 🚀 Quick Start

```bash
# Show today's prayer schedule (default: Riyadh)
pray

# Show next prayer with countdown
pray next

# Different city
pray --city Cairo

# Different calculation method
pray --method 2 --city London
```

## 📖 Usage

### Basic Commands

```bash
# Default view - show all prayers for today
pray
```

**Output:**
```
🕌 Prayer Times for Riyadh
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📅 31 Jan 2026 | 12 Shaʿbān, 1447 AH

    🌅 Fajr          05:15
    ☀️  Sunrise     06:35
    🌞 Dhuhr         12:06
    🌤️  Asr         15:14
    🌅 Maghrib       17:37
▶    🌙 Isha          19:07

⏰ Next prayer in 40m

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  📍 Method: Umm Al-Qura University, Makkah
```

### Next Prayer

```bash
pray next
```

**Output:**
```
🕌 Next Prayer
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  🌙 Isha at 19:07

⏰ In 40m

📍 Riyadh
```

### Different Cities

```bash
# Major cities work by name
pray --city "New York"
pray --city London
pray --city Dubai
pray --city Istanbul

# For better accuracy, include country
pray --city "Toronto" --country "CA"
```

### Calculation Methods

The `--method` flag controls the calculation methodology:

- `1` - University of Islamic Sciences, Karachi
- `2` - Islamic Society of North America (ISNA) 
- `3` - Muslim World League
- `4` - **Umm Al-Qura University, Makkah** (default - recommended for Saudi Arabia)
- `5` - Egyptian General Authority of Survey
- `7` - Institute of Geophysics, University of Tehran
- `8` - Gulf Region
- `9` - Kuwait
- `10` - Qatar
- `11` - Majlis Ugama Islam Singapura, Singapore
- `12` - Union Organization Islamic de France
- `13` - Diyanet İşleri Başkanlığı, Turkey
- `14` - Spiritual Administration of Muslims of Russia

**Example:**
```bash
# Use ISNA method for North America
pray --city Toronto --method 2

# Use Turkish method for Istanbul
pray --city Istanbul --method 13
```

## 🎨 Features

### Visual Highlights

- **Current/Next Prayer**: Highlighted with ▶ arrow and bright colors
- **Countdown Timer**: Shows time remaining in hours and minutes
- **Color Coding**: Different colors for prayer names, times, and information
- **Hijri Calendar**: Islamic date alongside Gregorian date
- **Method Information**: Shows which calculation method is being used

### Smart Time Handling

- **Automatic Next Prayer Detection**: Finds the next upcoming prayer
- **Tomorrow Handling**: When all today's prayers have passed, shows tomorrow's Fajr
- **Timezone Awareness**: Handles local timezone automatically
- **Sunrise Filtering**: Doesn't notify for sunrise in prayer mode (since it's not a prayer time)

## 🔧 Configuration

### Environment Variables

You can set default values using environment variables:

```bash
export PRAY_DEFAULT_CITY="Cairo"
export PRAY_DEFAULT_METHOD="3"
```

### Command Line Options

```bash
  --city string     City name for prayer times (default "Riyadh")
  --method int      Calculation method (4 = Umm Al-Qura) (default 4)
  -h, --help        Show help information
```

## 🌍 Supported Locations

The CLI works with thousands of cities worldwide. Examples:

### Major Cities
- **Saudi Arabia**: Riyadh, Jeddah, Mecca, Medina, Dammam
- **UAE**: Dubai, Abu Dhabi, Sharjah
- **Egypt**: Cairo, Alexandria
- **Turkey**: Istanbul, Ankara
- **Indonesia**: Jakarta, Surabaya
- **Pakistan**: Karachi, Lahore, Islamabad
- **Malaysia**: Kuala Lumpur, Johor Bahru
- **North America**: New York, Toronto, Los Angeles
- **Europe**: London, Paris, Berlin, Amsterdam

### Tips for Location Names
- Use English city names
- For common names, include country: `--city "London,UK"`
- Major cities usually work without country specification
- Use proper capitalization for best results

## 📡 API Information

This CLI uses the **Aladhan Prayer Times API** (https://aladhan.com/prayer-times-api), which provides:

- Accurate prayer time calculations
- Multiple calculation methods
- Global city database
- Hijri calendar integration
- No API key required

## 🔐 Privacy

- **No data collection**: All calculations are done via public API
- **No tracking**: No analytics or user behavior tracking
- **Local only**: No data stored locally except temporary cache

## 🛠️ Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/isIbra/pray.git
cd pray

# Download dependencies
go mod tidy

# Build the binary
go build -o pray

# Run tests
go test ./...

# Install globally (optional)
go install
```

### Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- Standard Go libraries for HTTP and JSON

## 📋 Requirements

- **Go 1.21+** for building from source
- **No runtime dependencies** - single binary
- **Internet connection** for fetching prayer times
- **Terminal with color support** for best experience

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes
4. Add tests if applicable
5. Commit: `git commit -m "Add amazing feature"`
6. Push: `git push origin feature/amazing-feature`
7. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Aladhan API](https://aladhan.com) for prayer time calculations
- [Charm](https://charm.sh) for beautiful terminal UI libraries
- The Muslim developer community for inspiration and feedback

## 🐛 Issues

Found a bug or have a feature request? Please create an issue on [GitHub](https://github.com/isIbra/pray/issues).

---

**Made with ❤️ for the Muslim community**

*"And establish prayer and give zakah and bow with those who bow [in worship and obedience]."* - Quran 2:43