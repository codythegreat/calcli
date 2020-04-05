# CALCLI

Calcli is a command line equation parser and solver

Calcli currently supports calculations using PEMDAS sin, cos, tan, and square root. Equations are also automatically converted to LaTaX format (both inline and display modes.)

If you see something in the code that can be improved or fixed, please feel free to fork the repo and make your changes. Help wanted.

## INSTALLATION

use the following command to download the repo into your go directory:
```
go get github.com/codythegreat/calcli
```
Once you have the source code you can cd into the direcotry `~/go/src/github.com/codythegreat/calcli`
Once inside the directory you can build the program
```
go build calcli.go
```
I recommend storing the executable in your home directory so that you can do this:
```
~/calcli "(2+2)^{3}"
```

## USAGE

Calcli can be used with or without flags -floor, -ceil, -round, -abs, -latexI and -latexD.

The -debug prints each function call with pertinent inputs and outputs

Currently you can only use one flag at a time.

```
./calcli -ceil "20000*1.1^20"
./calcli "(3.1415*5^{2})+sqrt{7}"
```

Note the following syntax used by Calcli:

```
./calcli "3^{3}"
./calcli "sqrt{3}"
./calcli "sin{3}"
./calcli "cos{3}"
./calcli "tan{3}"
```

## CONTRIBUTORS

Myself and [alexruf](https://github.com/alexruf)
