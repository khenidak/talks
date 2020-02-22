#include <stdio.h>
#include <unistd.h>

void fC(int i){
 	usleep(1000);
}

void fB(int i){
	fC(i);
}

void fA(int i){
	fB(i);
}

int main() {
	for(int i = 0; i < 100; i++)
		fA(i);

   return 0;
}

