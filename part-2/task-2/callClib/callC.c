#include <stdio.h>
#include "callC.h"

void helloFromC(){
	printf("Hello from C!\n");
}

void printMessage(char* message){
	printf("Go send me %s\n", message);
}