Alfred - The butler
===================
Alfred is a command-line tool that craws [shopicruit.myshopify.com/products.json](shopicruit.myshopify.com/products.json) and produces a result 
that meets Alice's requirements:

- An equal number of computers and keyboards.
- As many different computer and keyboard variants the store has to offer, while buying no duplicate variants.
- Spend the least amount of money possible.
- Print out total weight of all computers and keyboards.

### Install & Run
```
godep restore && go build && ./alfred
```

### Usage

```
Usage of ./alfred:
  -pages int
        Number of pages to craw (default 5)
  -protocol string
        Communication protocol (default "https")
  -url string
        Shop's url (default "shopicruit.myshopify.com/products.json")
```

### Result is displayed in ASCII table 

```
+--------------------------------+-----------------------+-------------------+
|              NAME              |         PRICE         | WEIGHT (IN GRAMS) |
+--------------------------------+-----------------------+-------------------+
| Aerodynamic Cotton Keyboard    |                  6.00 |               905 |
| (Blue)                         |                       |                   |
| Ergonomic Cotton Computer      |                  0.41 |              7143 |
| (Lime)                         |                       |                   |
| Heavy Duty Concrete Keyboard   |                 14.20 |              3030 |
| (Magenta)                      |                       |                   |
| Awesome Cotton Computer        |                  2.05 |              1906 |
| (Black)                        |                       |                   |
| Incredible Bronze Keyboard     |                 26.07 |              6623 |
| (Sky blue)                     |                       |                   |
| Awesome Cotton Computer        |                  2.91 |              4933 |
| (Orange)                       |                       |                   |
| Incredible Bronze Keyboard     |                 43.13 |               512 |
| (White)                        |                       |                   |
| Ergonomic Copper Computer      |                  3.51 |              3187 |
| (Black)                        |                       |                   |
| Incredible Silk Keyboard       |                 47.21 |              5140 |
| (Orchid)                       |                       |                   |
| Awesome Cotton Computer (Gold) |                  8.24 |               194 |
| Incredible Bronze Keyboard     |                 50.41 |              4183 |
| (Plum)                         |                       |                   |
| Awesome Cotton Computer        |                 10.78 |              4970 |
| (Orchid)                       |                       |                   |
| Incredible Silk Keyboard       |                 51.05 |              3622 |
| (Lime)                         |                       |                   |
| Rustic Wool Computer (Gold)    |                 11.37 |              5423 |
| Incredible Silk Keyboard       |                 53.73 |              5822 |
| (Orange)                       |                       |                   |
| Ergonomic Copper Computer      |                 12.45 |              7110 |
| (Red)                          |                       |                   |
| Incredible Bronze Keyboard     |                 58.86 |              4854 |
| (Green)                        |                       |                   |
| Awesome Bronze Computer        |                 18.76 |               922 |
| (Maroon)                       |                       |                   |
| Incredible Bronze Keyboard     |                 93.37 |              9973 |
| (Grey)                         |                       |                   |
| Ergonomic Cotton Computer      |                 18.99 |              5537 |
| (Magenta)                      |                       |                   |
+--------------------------------+-----------------------+-------------------+
|                                  TOTAL WEIGHT IN GRAMS |       85989       |
+--------------------------------+-----------------------+-------------------+
```
### Dependencies
- godep: `go get github.com/tools/godep` 
- tablewriter: `go get  github.com/olekukonko/tablewriter`
- httpmock: `go get github.com/jarcoal/httpmock`
