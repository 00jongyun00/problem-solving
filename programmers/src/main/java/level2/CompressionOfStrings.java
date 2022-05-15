package level2;

import static java.lang.Math.min;

public class CompressionOfStrings {

    public static int Solution(String s) {
        int answer = s.length();

        for (int i = 1; i <= s.length() / 2 + 1; i++) {
            int zipLevel = 1;
            String zipStr = s.substring(0, i);
            StringBuilder sb = new StringBuilder();

            for (int j = i; j <= s.length(); j += i) {
                String next = s.substring(j, min(j + i, s.length()));
                if (zipStr.equals(next)) {
                    zipLevel++;
                } else {
                    sb.append(zipLevel != 1 ? zipLevel : "").append(zipStr);
                    zipStr = next;
                    zipLevel = 1;
                }
            }
            sb.append(zipStr);
            answer = min(sb.length(), answer);
        }

        return answer;
    }

    public static void main(String[] args) {
        String s = "aabbaccc";
        System.out.println(Solution(s));
    }
}
