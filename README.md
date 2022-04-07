# Dependency
```
go get github.com/xh-dev-go/luhnId@latest
```

# Installation
```
go install github.com/xh-dev-go/luhnId@latest
```

# Demo

```
luhnId generate --starting-digit 2 --digit 20
# output: 20458242474622402056

luhnId validate --code 20458242474622402056
# output: Valid code
```