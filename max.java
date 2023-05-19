public class Max {
    public static int max(int a, int b, int c) {
        return Math.max(Math.max(a, b), c);
    }

    public static void main(String[] args) {
        System.out.println(max(5, 10, 3));
    }
}

