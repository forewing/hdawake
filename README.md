# HD Awake
Keep your hard drive awake

## Usage

```
Usage of hdawake:
  -path string
        path to target (default "./.awake")
  -time duration
        period of wake up action (default 1m0s)
```

## Install

- Get pre-built binary from [release page](https://github.com/forewing/hdawake/releases/latest)

- Or install using go

    ```
    # Go 1.16+
    go install github.com/forewing/hdawake@latest

    # Go ~1.15
    # go get github.com/forewing/hdawake@latest
    ```

## Build

```
git clone https://github.com/forewing/hdawake.git
cd hdawake
go build github.com/forewing/hdawake/cmd/hdawake
```
