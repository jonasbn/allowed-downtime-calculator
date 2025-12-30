# Calculate Allowed Downtime

## Introduction

I read a LinkedIn post on the subject and I wanted to implement this in Go, just for the fun of it.

In retrospect I should never have started, since descrepancies occurred between my implementation and the other sources. I outlined this in a [blog post](https://dev.to/jonasbn/title-calculating-allowed-downtime-for-human-consumption-99-to-9999999-availability-3aid).

The other day, I found out by accident that there are various lengths to a year based on the definition of a year. I had implemented handling of leap years in my implementation, but there are two year definitions that I had not taken into account.

- tropical year of approx. `365.2425` days
- common year of `365` days

I found a nice explanation on [Wikipedia](https://en.wikipedia.org/wiki/Year) and I have updated my implementation to handle this. I also found another [article][SIBE] explaining the difference between a calculated year (calculated) and the the actual year length based on earts revolution around the sun. [Blog post](https://dev.to/jonasbn/calculating-allowed-downtime-for-human-consumption-a-follow-up-5dni) for the follow up.

A common year is:

```text
24 hours/day * 60 minutes/hour * 60 seconds/minute * 365 days/year = 31,536,000 seconds/year
```

This is the Gregorian calendar year length (common year), but the actual year length (tropical) is:

```text
24 hours/day * 60 minutes/hour * 60 seconds/minute * 365.2425 days/year = 31,556,952 seconds/year
```

My implementation calculates the allowed downtime for a given year based on the a range of uptime percentages for easy human consumption based on a gregorian calendar, which is the default and similar to tropical year length for the calculations in this implementation.

You can specify a calendar using: `--calendar`, the default however is the gregorian year length. The [post from LinkedIn][LINKEDIN] was based on the gregorian/tropical year length, which explains my initial descrepancies (and confusion), since my first implementation was based on the common year length.

These are the percentages I have based my calculations on:

- `99%` (2 nines), one can actually have a life or at least a weekend
- `99.9%` (3 nines), one can get _some_ sleep
- `99.99%` (4 nines), you are getting serious and on call
- `99.999%` (5 nines), the de facto _expectancy_ for most services
- `99.9999%` (6 nines), ambitious
- `99.99999%` (7 nines), which should is somewhat the same as 100% uptime

### Allowed Downtime Matrix

This I have based on the numbers from:

- [LinkedIn post][LINKEDIN], which is based on the gregorian or tropcial year length
- [uptime.is][UPTIMEIS] _source of calculation unknown to me_
- and finally my own initial calculations from this implementations which is a common year
- and the updated implementation which can handle both common and tropical year lengths

## 99% Availability: Calculated allowed downtime

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| gregorian | `3`  | `15`  | `39`    | `29`    |
| uptime.is | `3`  | `14`  | `56`    | `18`    |
| common    | `3`  | `15`  | `36`    | `0`     |
| tropical  | `3`  | `15`  | `39`    | `29`    |

## 99.9% Availability: Calculated allowed downtime

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| gregorian | `0`  | `8`   | `45`    | `56`    |
| uptime.is | `0`  | `8`   | `41`    | `38`    |
| common    | `0`  | `8`   | `45`    | `35`    |
| tropical  | `0`  | `8`   | `45`    | `56`    |

## 99.99% Availability: Calculated allowed downtime

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| gregorian | `0`  | `0`   | `52`    | `35`    |
| uptime.is | `0`  | `0`   | `52`    | `9.8`   |
| common    | `0`  | `0`   | `52`    | `33`    |
| tropical  | `0`  | `0`   | `52`    | `35`    |

## 99.999% Availability: Calculated allowed downtime

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| gregorian | `0`  | `0`   | `5`     | `15`    |
| uptime.is | `0`  | `0`   | `5`     | `13`    |
| common    | `0`  | `0`   | `5`     | `15`    |
| tropical  | `0`  | `0`   | `5`     | `15`    |

## 99.9999% Availability: Calculated allowed downtime

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| gregorian | `0`  | `0`   | `0`     | `31`    |
| uptime.is | `0`  | `0`   | `0`     | `31`    |
| common    | `0`  | `0`   | `0`     | `31`    |
| tropical  | `0`  | `0`   | `0`     | `31`    |

## 99.99999% Availability: Calculated allowed downtime

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| gregorian | `0`  | `0`   | `0`     | `3`     |
| uptime.is | `0`  | `0`   | `0`     | `3.1`   |
| common    | `0`  | `0`   | `0`     | `3`     |
| tropical  | `0`  | `0`   | `0`     | `3`     |

## Implementation

The implementation started out being based on Integers, but I quickly realized that I needed to use floating point numbers to get the precision I believe I need.

In addition the implementation used the Go package: `math` and it's `Mod` function.

- [golangbyexample.com: "Remainder or Modulus in Go (Golang)"](https://golangbyexample.com/remainder-modulus-go-golang/)
- [golang.org: "Package math"](https://pkg.go.dev/math)
- [golang.org: "Package math: Mod function"](https://pkg.go.dev/math#Mod)

### Leap Years

The implementation also supports handling of leap years for common years, since the number of days in a year is not a constant and it varies by a day.

The leap year calculation is based on the Exercism exercise "leap" and my own solution to this exercise.:

- [Exercism: "leap" exercise](https://exercism.org/tracks/go/exercises/leap)

#### Regular Year vs Leap Year for the Common Year

| Availability | Year       | Days | Hours | Minutes | Seconds |
|--------------|------------|------|-------|---------|---------|
| `99.000000%` | Regular    | `3`  | `15`  | `36`    | `0`     |
| `99.000000%` | Leap       | `3`  | `15`  | `50`    | `24`    |
| `99.900000%` | Regular    | `0`  | `8`   | `45`    | `35`    |
| `99.900000%` | Leap       | `0`  | `8`   | `47`    | `2`     |
| `99.990000%` | Regular    | `0`  | `0`   | `52`    | `33`    |
| `99.990000%` | Leap       | `0`  | `0`   | `52`    | `42`    |
| `99.999000%` | Regular    | `0`  | `0`   | `5`     | `15`    |
| `99.999000%` | Leap       | `0`  | `0`   | `5`     | `16`    |
| `99.999900%` | Regular    | `0`  | `0`   | `0`     | `31`    |
| `99.999900%` | Leap       | `0`  | `0`   | `0`     | `31`    |
| `99.999990%` | Regular    | `0`  | `0`   | `0`     | `3`     |
| `99.999990%` | Leap       | `0`  | `0`   | `0`     | `3`     |

As you can read from the table above, the difference is not that big, but it is there.

- `99.000000%` difference of `14` minutes and `24` seconds
- `99.900000%` difference of `1` minute and `27` seconds
- `99.990000%` difference of `9` seconds
- `99.999000%` difference of `1` second
- `99.999900%` no difference
- `99.999990%` no difference

## Installation

To install the CLI application, clone the repository and navigate to the project directory:

```bash
git clone git@github.com:jonasbn/allowed-downtime-calculator.git
cd allowed-downtime-calculator
```

Then, run the following command to install the required dependencies:

```bash
go mod tidy
```

## Usage

To run the CLI application, use the following command:

```bash
go run cmd/main.go [options]
Calculated allowed downtime for uptime requirement in year: 2025 (365.242500 days) as per gregorian year length:
   99.000000% is: 3 days 15 hours 39 minutes 29 seconds
   99.900000% is: 0 days 8 hours 45 minutes 56 seconds
   99.990000% is: 0 days 0 hours 52 minutes 35 seconds
   99.999000% is: 0 days 0 hours 5 minutes 15 seconds
   99.999900% is: 0 days 0 hours 0 minutes 31 seconds
   99.999990% is: 0 days 0 hours 0 minutes 3 seconds
```

### Options

- `year`: The year for which you want to calculate the allowed downtime. Defaults to current year.
- `calendar`:
  - `common`: Use the common year length of: `365` days/year.
  - `tropical`: Use the tropical year length of: `365,2422` days/year.
  - `gregorian`: Use the gregorian calendar year length of: `365,2425` days/year.
  - Defaults to `gregorian`
- `debug`: Enable debug mode to print additional information (currently very limited).

As mentioned earlier, the implementation supports handling of leap years, so you can provide any year you want and get the proper calculation.

Since there are only two outcomes and you are unsure or just want a regular year `2025` was a regular year and `2024` was a leap year.

### Parameters

In addition to the options, you can also provide data as parameters to the CLI application:

```bash
go run cmd/main.go 0 50 100
```

Do note parameters can be provided as integers or floats.

### Diagnostics

#### invalid parameter type

If a provided parameter cannot be converted to a float for use in the calculations, the application will print an error message and fallback to the defaults.

#### invalid parameter value

If a provided parameter is not representing a percentile between `0` and `100`, the application print an error message and fallback to the defaults.

## Testing

To run the tests for the CLI application, use the following command:

```bash
go test -v ./pkg/cli
```

## Resources and References

- [Post from LinkedIn][LINKEDIN]
- [uptime.is][UPTIMEIS]
- [Wikipedia: "Year"][WIKIPEDIA]
- [Wikipedia: "Tropical year"][TROPICAL]
- [Wikipedia: "Common year"][COMMON]
- [Wikipedia: "Leap year"][LEAP]
- [SibeNotes: "How many seconds are in a year?"][SIBE]

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

[LINKEDIN]: https://www.linkedin.com/feed/update/urn:li:activity:7286283676743540736/
[UPTIMEIS]: https://uptime.is/
[WIKIPEDIA]: https://en.wikipedia.org/wiki/Year
[SIBE]: https://sibenotes.com/maths/how-many-seconds-are-in-a-year/
[TROPICAL]: https://en.wikipedia.org/wiki/Tropical_year
[COMMON]: https://en.wikipedia.org/wiki/Common_year
[LEAP]: https://en.wikipedia.org/wiki/Leap_year
