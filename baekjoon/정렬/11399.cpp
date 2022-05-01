#include <stdio.h>

int main() {
  int p[1001], n, i, j, accu=0;

  scanf("%d", &n);
  for (i=0; i<n; i++) {
    scanf("%d", &p[i]);
  }

  for (i=0; i<n; i++) {
    int tmp = p[i];
    for (j=i-1; j>=0; j--) {
      if (p[j] > tmp) p[j+1] = p[j];
      else break;
    }
    p[j+1] = tmp;
  }

  for (i=0; i<n; i++) {
    for (j=i; j>=0; j--) {
      accu += p[j];
    }
  }
  printf("%d\n", accu);
}
