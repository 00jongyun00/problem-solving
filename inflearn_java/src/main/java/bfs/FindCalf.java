package bfs;

import java.util.LinkedList;
import java.util.Queue;

public class FindCalf {

    public static int[] distance = {1, -1, 5};
    public static int[] check = new int[10001];

    public static int Solution(int s, int e) {

        int level = 1;
        Queue<Integer> queue = new LinkedList<>();
        queue.add(s);

        while (!queue.isEmpty()) {
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                int target = queue.poll();
                for (int dis : distance) {
                    int nextTarget = target + dis;
                    if (nextTarget == e) {
                        return level;
                    }
                    if (nextTarget >= 1 && nextTarget <= 10000 && check[nextTarget] == 0) {
                        check[nextTarget] = 1;
                        queue.offer(nextTarget);
                    }
                }
            }
            level++;
        }

        return level;
    }


    public static void main(String[] args) {
        int result = Solution(5, 14);
        System.out.println(result);
    }
}
