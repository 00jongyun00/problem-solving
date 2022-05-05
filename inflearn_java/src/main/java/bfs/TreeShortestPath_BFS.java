package bfs;

import java.util.LinkedList;
import java.util.Queue;

public class TreeShortestPath_BFS {

    public static int Solution(int level, Node root) {
        Queue<Node> queue = new LinkedList<>();
        queue.offer(root);

        while (!queue.isEmpty()) {
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                Node target = queue.poll();
                if (target.lt == null && target.rt == null) {
                    return ++level;
                }
                if (target.lt != null) {
                    queue.offer(target.lt);
                }
                if (target.rt != null) {
                    queue.offer(target.rt);
                }
            }
            level++;
        }
        return level;
    }

    public static void main(String[] args) {
        Node node = new Node(1);
        node.lt = new Node(2);
        node.rt = new Node(3);
        node.lt.lt = new Node(4);
        node.lt.rt = new Node(5);
        node.rt.lt = new Node(6);
        node.rt.rt = new Node(7);
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
