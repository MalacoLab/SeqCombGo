# SeqCombGo
**Seq**uence **Comb**ination tool written in **Go**lang

# Requirement

- go [for compile]

# Compile
```
git clone git@github.com:MalacoLab/SeqCombGo.git
cd SeqCombGo
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

# Cite

```
@Article{An2022,
  author = {An, G; Zhang, G},
  title  = {SeqCombGo: Sequence Combination tool written in Golang},
  year   = {2022},
  url    = {https://github.com/MalacoLab/SeqCombGo},
}
```

# Detail

- [ ] datatype recognize
