# Calculate Allowed Downtime

## Introduction

I read a LinkedIn post on the subject and I wanted to implement this in Go, just for the fun of it.

In retrospect I should never have started, since descrepancies occurred between my implementation and the other sources.

My implementation calculates the allowed downtime for a given year based on the a range of uptime percentages for easy human consumption.

These are the percentages I have based my calculations on:

- `99%` (2 nines), one can actually have a life or at least a weekend
- `99.9%` (3 nines), one can get _some_ sleep
- `99.99%` (4 nines), you are getting serious and on call
- `99.999%` (5 nines), the de facto _expectancy_ for most services
- `99.9999%` (6 nines), ambitious
- `99.99999%` (7 nines), which should is somewhat the same as 100% uptime

### Allowed Downtime Matrix

This I have based on the numbers from the [LinkedIn post][LINKEDIN] and the numbers copied from [uptime.is][UPTIMEIS] and finally my own calculations from this implementation.

## 99% Availability: Calculated allowed downtime

Here we all really disagree

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| LinkedIn  | `3`  | `15`  | `39`    | `29`    |
| uptime.is | `3`  | `14`  | `56`    | `18`    |
| _mine_    | `3`  | `15`  | `36`    | `0`     |

## 99.9% Availability: Calculated allowed downtime

Here we all disagree, but we are closer

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| LinkedIn  | `0`  | `8`   | `45`    | `56`    |
| uptime.is | `0`  | `8`   | `41`    | `38`    |
| _mine_    | `0`  | `8`   | `45`    | `35`    |

## 99.99% Availability: Calculated allowed downtime

Here we all disagree, but we are even closer

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| LinkedIn  | `0`  | `0`   | `52`    | `35`    |
| uptime.is | `0`  | `0`   | `52`    | `9.8`   |
| _mine_    | `0`  | `0`   | `52`    | `33`    |

## 99.999% Availability: Calculated allowed downtime

Linked post and I agree on this one, but not with uptime.is

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| LinkedIn  | `0`  | `0`   | `5`     | `15`    |
| uptime.is | `0`  | `0`   | `5`     | `13`    |
| _mine_    | `0`  | `0`   | `5`     | `15`    |

## 99.9999% Availability: Calculated allowed downtime

All is good

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| LinkedIn  | `0`  | `0`   | `0`     | `31`    |
| uptime.is | `0`  | `0`   | `0`     | `31`    |
| _mine_    | `0`  | `0`   | `0`     | `31`    |

## 99.99999% Availability: Calculated allowed downtime

Minor descrepancy where uptime.is has an extra 0.1 second

| Source    | Days | Hours | Minutes | Seconds |
|-----------|------|-------|---------|---------|
| LinkedIn  | `0`  | `0`   | `0`     | `3`     |
| uptime.is | `0`  | `0`   | `0`     | `3.1`   |
| _mine_    | `0`  | `0`   | `0`     | `3`     |

## Implementation

The implementation started out being based on Integers, but I quickly realized that I needed to use floating point numbers to get the precision I believe I need.

In addition the implementation used the Go package: `math` and it's `Mod` function.

- [golangbyexample.com: "Remainder or Modulus in Go (Golang)"](https://golangbyexample.com/remainder-modulus-go-golang/)
- [golang.org: "Package math"](https://pkg.go.dev/math)
- [golang.org: "Package math: Mod function"](https://pkg.go.dev/math#Mod)

### Leap Years

The implementation also suupports handling of leap years, since the number of days in a year is not a constant and it varies by a day.

The leap year calculation is based on the Exercism exercise "leap" and my own solution to this exercise.:

- [Exercism: "leap" exercise](https://exercism.org/tracks/go/exercises/leap)

#### Regular Year vs Leap Year

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
git clone <repository-url>
cd my-cli-app
```

Then, run the following command to install the required dependencies:

```bash
go mod tidy
```

## Usage

To run the CLI application, use the following command:

```bash
go run cmd/main.go [options]
```

### Options

- `year`: The year for which you want to calculate the allowed downtime. Defaults to current year.

As mentioned earlier, the implementation supports handling of leap years, so you can provide any year you want and get the proper calculation.

Since there are only two outcomes and you are unsure or just want a regular year `2025` was a regular year and `2024` was a leap year.

## Resources and References

- [Post from LinkedIn][LINKEDIN]
- [uptime.is][UPTIMEIS]

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

[LINKEDIN]: https://www.linkedin.com/feed/update/urn:li:activity:7286283676743540736/
[UPTIMEIS]: https://uptime.is/
