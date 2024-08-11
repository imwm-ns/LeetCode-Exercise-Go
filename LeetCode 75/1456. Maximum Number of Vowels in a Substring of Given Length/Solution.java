public class Solution {
    public static void main(String[] args) {
        String s = "abciiidef";
        int k = 3;

        System.out.println(maxVowels(s, k));
    }

    public static int maxVowels(String s, int k) {
        int mx = 0, cnt = 0;

        for (int i = 0; i < s.length(); i++) {
            if (isVowel(s.charAt(i))) cnt++;
            if (i >= k && isVowel(s.charAt(i - k))) cnt--;

            mx = Math.max(cnt, mx);
        }

        return mx;
    }

    private static boolean isVowel(char c) {
        if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
            return true;
        }
        return false;
    }
}