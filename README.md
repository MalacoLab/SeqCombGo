# gocomb
sequence combination tool written in Golang

# Requirement

- go

# compile
```
git clone git@github.com:MalacoLab/gocomb.git
cd gocomb
go build
```

# how to use

## convert fas to nex
```
./gocomb -o output.nex import.fas
./gocomb import.fas
```

## combine serveral fas to single nex

```
./gocomb import1.fas import2.fas
./gocomb -o export.nex import1.fas import2.fas
```
