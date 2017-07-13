#include <stdio.h>
#include <math.h>
#include <time.h>

int main()
{
	int sum = 0;

	clock_t ttt;
	ttt = clock();
	for (int i = 0; i <= 10000; i++)
	{
		sum = sum + i;
	}

	ttt = clock() - ttt;
	printf("%d\n", sum);
	printf("It takes %e seconds.\n",((float)ttt)/CLOCKS_PER_SEC);

	return 0;
}