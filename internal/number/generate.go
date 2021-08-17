package number


//go:generate gcc -c -o number.o number.c
//go:generate ar rsc libnumber.a number.o
