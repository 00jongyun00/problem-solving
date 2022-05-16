package level2;

import java.util.ArrayList;
import java.util.HashMap;

public class OpenChatRoom {

    public String[] solution(String[] record) {
        ArrayList<String> arr = new ArrayList<>();
        HashMap<String, String> map = new HashMap<>();

        for (int i = 0; i < record.length; i++) {
           String[] command = record[i].split(" ");
           String event = command[0];

           if (event.equals("Enter")) {
               arr.add(command[1] + "님이 들어왔습니다.");
               map.put(command[1], command[2]);
           } else if (command[0].equals("Change")) {
               map.put(command[1], command[2]);
           } else {
               arr.add(command[1] + "님이 나갔습니다.");
           }
        }

        String[] answer = new String[arr.size()];
        for (int i = 0; i < arr.size(); i++) {
            int idx = arr.get(i).indexOf("님");
            String id = arr.get(i).substring(0, idx);
            String[] temp = arr.get(i).split(" ");
            answer[i] = map.get(id) + "님이 " + temp[1];
        }
        return answer;
    }
}
