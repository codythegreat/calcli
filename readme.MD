## CALCLI

A command line equation parser and solver

## USAGE

Calcli can be used with or without flags -floor, -ceil, and -round.

Calcli can handle any spaces inside of the equation

./calcli -ceil "20000*1.1^20"
./calcli "(3.1415*5^2)+7"

## INSTALLATION

Calcli can be installed using go build calcli.go.

I recommend storing the executable in your home directory so that you can do this:

~/calcli "2+2"


