# carcat

Concatenate CAR files (v1).

## Usage

```
go install github.com/hsanjuan/carcat
carcat file1.car file2.car file3.car ... > merged.car
```

The roots of the resulting `merged.car` file will be set to those of the last
input file provided.

This application does not verify anything read from the CAR files.
