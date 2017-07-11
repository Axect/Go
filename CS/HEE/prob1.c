#define _CRT_SECURE_NO_WARNINGS
#include <stdio.h>

float hbar()
{
	return 6.582119514e-16;
}

int main()
{

	printf("Insert non-nagative integer:\n");

	float num1;
	float num2;
	float num3;

	scanf("%f", &num1);

	for (;;)
	{
		if (num1!=(int)num1||num1<0)
		{
			printf("Incorrect.\n Please insert the correct number.\n");
			scanf("%f", &num1);

		}
		else
		{
			num2 = num1 + ((float)1 / 2);
			num3 = num2*hbar();
			printf("E_%d = %eh\n", (int)num1, num2);
			printf("E_%d = %eeV\n", (int)num1, num3);
			break;
		}

	}
	return 0;
}