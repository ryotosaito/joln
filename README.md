# joln
Join lines with a specific separator

## Installation
```bash
$ go get -u github.com/ryotosaito/joln
```

## Usage
### Basic
```bash
$ seq 10 | joln -s ,
```
Output:
```
1,2,3,4,5,6,7,8,9,10
```

### Folding by specific number of columns
```bash
$ yes 'hello world' | head -10 | joln -s '|' -w 3
```
Output:
```
hello world|hello world|hello world
hello world|hello world|hello world
hello world|hello world|hello world
hello world
```
