package level2;

public class CompressionOfStrings {

    public static int Solution(String s) {
        int answer = s.length();

        for (int i = 1; i <= s.length() / 2; i++) {
            int zipLevel = 1;
            String zipStr = s.substring(0, i);
            StringBuilder result = new StringBuilder();

            for (int j = i; j <= s.length(); j += i) {
                String next = s.substring(j, Math.min(j + i, s.length()));
                if (zipStr.equals(next)) {
                    zipLevel++;
                } else {
                    result.append(zipLevel != 1 ? zipLevel : "").append(zipStr);
                    zipStr = next;
                    zipLevel = 1;
                }
            }
            result.append(zipStr);
            answer = Math.min(answer, result.length());
        }
        return answer;
    }

    public static void main(String[] args) {
        String s = "aabbaccc";
        System.out.println(Solution(s));
    }
}
