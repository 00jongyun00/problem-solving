#include <stdio.h>
#include <vector>
#include <algorithm>

using namespace std;

bool comp(int a, int b) {
    if (a > b) return true;
    return false;
}

int main() {
    int N, S=0;
    int a[50];
    int b[50];
    scanf("%d", &N);

    for (int i = 0; i < N; i++) scanf("%d", &a[i]);
    for (int i = 0; i < N; i++) scanf("%d", &b[i]);

    sort(a, a+N);
    sort(b, b+N, comp);
    for (int i = 0; i < N; i++) {
        S = S + a[i] * b[i];
    }
    printf("%d", S);

    return 0;
}
