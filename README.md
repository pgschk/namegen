# namegen
A very simple name generator, currently for Kerbal names

## Build
```bash
go build -o namegen ./cmd/namegen/
```

## Run
```bash
./namegen
```
```bash
# generate a different amount of names
./namegen -n 1
```
```bash
# generate a lot of name slowly, to read while generating
./namegen -n 100 --slowmode
```