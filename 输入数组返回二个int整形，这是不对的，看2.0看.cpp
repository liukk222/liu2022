# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>

//void aff(int arr[],int sz,int*px,int*py)
//{
//	int a = 0;
//	int j = 0;
//	int as[1] = { 0 };
//	for (a = 0; a <= sz; a++)
//	{
//		for (j = 0; j < sz - 1; j++)
//		{if
//		(arr[j] > arr[j + 1])
//		{
//			as[0] = arr[j + 1];
//			arr[j + 1] = arr[j];
//			arr[j] = as[0];
//		}
//		}
//}
//	//int b = 0;
//	//printf("%d %d\n", arr[sz - 2], arr[sz - 1]);
//	for (a= 0; a < sz - 1; a= a+ 2)
//	{if(arr[a]!=arr[a+1])
//		/*printf("%d\n %d\n", arr[a], arr[a+ 1]);*/
//	{
//		*px = arr[a];
//		*py = arr[a + 1];
//	}
//	}
//	return;
//}
//
//int main()
//{
//	int arr[] = { 1,2,3,4,5,6,1,2,3,4 };
//	int sz = sizeof(arr) / sizeof(arr[0]);
//	int x = 0;
//	int y = 0;
//	aff(arr, sz, &x, &y);
//	//int a = 0;
//	
//
//	printf("%d %d", x, y);
//return 0;
//}

int main()
{
	int a = 0;
	for (a = 100; a <= 200; a++)
	{
		int b = 0;
		for (b = 2; b < a; b++)
		{
			if (a % b == 0) { break; }
		}
		if(a==b)
		printf("%d\n", a);

	}
	return 0;
}
