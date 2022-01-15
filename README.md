# SeqCombGo
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
SeqCombGo -o output.nex import.fas
SeqCombGo import.fas
```

## Combine serveral fas to single nex

```
SeqCombGo import1.fas import2.fas
SeqCombGo -o export.nex import1.fas import2.fas
```
