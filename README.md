# GoPassGen

A flexible command-line password generator written in Go that allows you to customize the length and character composition of generated passwords.

## Features

- Generate secure random passwords with configurable length
- Control the minimum and maximum number of numeric characters
- Control the minimum and maximum number of special characters
- Automatic balancing of character types to meet constraints

## Installation

### Prerequisites

- Go 1.16 or later

### Building from source

1. Clone the repository:
```
git clone https://github.com/yourusername/gopassgen.git
cd gopassgen
```

2. Build the application:
```
go build -o gopassgen
```

3. (Optional) Install the application:
```
go install
```

## Usage

```
gopassgen [options]
```

### Command-line Options

| Flag        | Description                                   | Default |
|-------------|-----------------------------------------------|---------|
| -length     | Total password length                         | 20      |
| -min-nums   | Minimum number of numeric characters          | 1       |
| -max-nums   | Maximum number of numeric characters          | 10      |
| -min-spec   | Minimum number of special characters          | 1       |
| -max-spec   | Maximum number of special characters          | 10      |
| -help       | Display help information                      | false   |

### Examples

Generate a password with default settings (20 characters, at least 1 number and 1 special character):
```
gopassgen
```

Generate a 16-character password with at least 2 numbers and 2 special characters:
```
gopassgen -length 16 -min-nums 2 -min-spec 2
```

Generate a password with exactly 3 numbers and 3 special characters:
```
gopassgen -min-nums 3 -max-nums 3 -min-spec 3 -max-spec 3
```

Generate a password with only letters (no numbers or special characters):
```
gopassgen -min-nums 0 -max-nums 0 -min-spec 0 -max-spec 0
```

## Testing

Run the tests with:
```
go test
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
