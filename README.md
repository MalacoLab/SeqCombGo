# GoComb
sequence combination tool written in Golang

# Requirement

- go [for compile]

# Compile
```
git clone git@github.com:MalacoLab/gocomb.git
cd gocomb
go build
```

# How to use

## Convert fas to nex
```
./gocomb -o output.nex import.fas
./gocomb import.fas
```

## Combine serveral fas to single nex

```
./gocomb import1.fas import2.fas
./gocomb -o export.nex import1.fas import2.fas
```
