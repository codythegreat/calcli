## CALCLI

Calcli is a command line equation parser and solver

Calcli currently supports calculations using PEMDAS and square root. Equations are also automatically converted to LaTaX format (both inline and display modes.)

## USAGE

Calcli can be used with or without flags -floor, -ceil, -round, and -abs.

Currently you can only use one flag at a time.

Calcli can handle any spaces inside of the equation

```
./calcli -ceil "20000*1.1^20"
./calcli "(3.1415 * 5 ^ 2) + SQRT{7}"
```

## INSTALLATION

Calcli can be installed using go build calcli.go.

I recommend storing the executable in your home directory so that you can do this:

```
~/calcli "(2+2)^3"
```

