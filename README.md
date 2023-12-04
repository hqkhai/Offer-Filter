# Offers

## Prerequisite

You should install golang version 1.21.0 on your pc

## How to run?

```
go run main.go
```

After successfully executing the command, the program will prompt you to enter a date string in the format 'YYYY-MM-DD'. If you input the wrong format, the program will report an error. After filtering, the program will write to an 'output.json' file. You can stop the program by entering 'x' letter.

1. Using the functional option pattern, the program provides a safe mechanism for users to define certain values for the 'offerFilter.' For example:

```
offerFilter := offer.NewOfferFilter(
    offer.WithInputFileName("input"),
    offer.WithOutputFileName("output"),
    offer.WithMaxDate(5),
    offer.WithNotEligibleCategories("Hotel"),
)
```

**Options**
| Option | Description | Default Values |
| - | - | - |
| **WithInputFileName** | Define the input file name| input |
| **WithOutputFileName** | Define the output file name | output |
| **WithMaxDate** | The number of days used to check offer validity | 5 |
| **WithNotEligibleCategories** | List of invalid categories | Hotel |

2. With the Template Method Pattern, in the future, we can easily expand the program to read from files in various formats such as .csv, .txt, and more.
