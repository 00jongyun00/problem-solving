package PS_2022_05_15;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class NQueen {

    public static int[] check;


    public static int n;
    public static int count = 0;

    public static boolean Possibility(int col) {
        // 해당 열의 행과 i 열의 행이 일차할 경우 (같은 행에 존재할 경우)
        for (int i = 0; i < col; i++) {
            if (check[col] == check[i]) {
                return false;
            }
            if (Math.abs(col - i) == Math.abs(check[col] - check[i])) {
                return false;
            }
        }
        return true;
    }

    public static void Solution(int level) {
        if (level == n) {
            count++;
            return;
        }

        for (int i = 0; i < n; i++) {
            check[level] = i;
            if (Possibility(level)) {
                Solution(level + 1);
            }
        }
    }

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader((System.in)));
        n = Integer.parseInt(br.readLine());
        check = new int[n];
        Solution(0);
        System.out.println(count);
    }
}
