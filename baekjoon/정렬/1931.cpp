#include <stdio.h>
#include <vector>
#include <algorithm>
using namespace std;

typedef pair<int, int> P;
vector<P> arr;

bool comp(P a, P b) {
    if (a.second == b.second) {
        return a.first < b.first;
    }
    return a.second < b.second;
}

int main() {
    int N;
    scanf("%d", &N);
    // left: start time right: end time
    int left, right;
    for (int i = 0; i < N; i++) {
        scanf("%d %d", &left, &right);
        arr.push_back(P(left, right));
    }
    sort(arr.begin(), arr.end(), comp);
    int now = arr[0].second;
    int cnt = 1;

    for (int i=1; i<N; i++) {
        if (now <= arr[i].first) {
            cnt++;
            now = arr[i].second;
        }
    }
    printf("%d\n", cnt);

    return 0;
}
