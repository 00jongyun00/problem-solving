#include <stdio.h>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
  int n, m, i, left, right, pivot;
  scanf("%d", &n);
  vector<int> a(n);
  for (i=0; i<n; i++) scanf("%d", &a[i]);

  scanf("%d", &m);
  vector<int> b(m);
  for (i=0; i<m; i++) scanf("%d", &b[i]);

  sort(a.begin(), a.end());

  for (i=0; i<m; i++) {
    left = 0;
    right = n;
    pivot = (left + right) / 2;
    while (left <= right) {
      if (a[pivot] == b[i]) {
        printf("1\n");
        break;
      }
      if (a[pivot] < b[i]) left = pivot + 1;
      else right = pivot - 1;
      pivot = (left + right) / 2;
    }
    if (left > right) printf("0\n");
  }

  return 0;
}
