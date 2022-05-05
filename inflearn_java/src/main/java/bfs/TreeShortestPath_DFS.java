package bfs;

public class TreeShortestPath_DFS {

    public static int Solution(int level, Node root) {
        if (root.lt == null && root.rt == null) {
            return level;
        } else {
            return Math.min(
                Solution(level + 1, root.lt),
                Solution(level + 1, root.rt)
            );
        }
    }

    public static void main(String[] args) {
        Node node = new Node(1);
        node.lt = new Node(2);
        node.rt = new Node(3);
        node.lt.lt = new Node(4);
        node.lt.rt = new Node(5);
        int result = Solution(0, node);
        System.out.println(result);
    }

    static class Node {

        int data;
        Node lt, rt;

        public Node(int val) {
            data = val;
            lt = rt = null;
        }
    }

}
