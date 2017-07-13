#include <stdio.h>
#include <math.h>
#include <float.h>

float h()
{
	float num1 = 1e-06;

	return num1;
}

int main()
{
	double value001 = (1*powf(h(), 2)-powf(-h(), 2)) / (2 * h());
	printf("The derivate of x square at x=0 is %f.\n", value001);

	double value002 = (sin(h()) - sin(-h())) / (2 * h());
	printf("The derivate of sin x at x=0 is %f.\n", value002);

	double value003 = (exp(pow(10,-11)) - exp(-pow(10,-11))) / (2*pow(10,-11));
	printf("The derivate of exponential x at x=0 is %f.\n", value003);
	//if power of exp is smaller than -11, value003 gets far from 1.00000. so I can't insert h().


	printf("The positive minimum of float type is %e\n", h());
	return 0;
}