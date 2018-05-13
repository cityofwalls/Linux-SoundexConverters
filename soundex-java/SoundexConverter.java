import java.util.HashMap;

public class SoundexConverter {
    public static void main(String[] args) {
        HashMap<String, String> conversions = new HashMap<String, String>();
        conversions.put("bfpv", "1"); conversions.put("cgjkqsxz", "2"); conversions.put("dt", "3"); 
        conversions.put("l", "4"); conversions.put("mn", "5"); conversions.put("r", "6");
        conversions.put ("aeiouyhw", "0");
        
        for (int i = 0; i < args.length; i++) {
            args[i] = nameConverter(args[i], conversions);
            System.out.print(args[i] + " ");
        }
        System.out.println();
    }
    
    public static String nameConverter(String name, HashMap<String, String> conversions) {
        char firstLetter = name.charAt(0);
        String noHW = name.replaceAll("[wh]", "");
        String convertedName = convertConsonantes(noHW, conversions);
        String noDupes = removeDuplicateLetters(convertedName, conversions);
        noDupes = noDupes.replaceAll("[0]", "");
        
        String result = "";
        int tooSmall = 4 - noDupes.length();
        if (tooSmall <= 0) {
            result = noDupes.substring(0, 4);
        } else {
            result = noDupes;
            for (int i = 0; i < tooSmall; i++) {
                result += "0";
            }
        }
        return result;
    }
    
    public static String convertConsonantes(String letters, HashMap<String, String> conversions) {
        String numbers = letters;
        for (String key : conversions.keySet()) {
            String regex = "[" + key + "]";
            String soundexNum = conversions.get(key);
            numbers = numbers.replaceAll(regex, soundexNum);
        }
        return numbers;
    }
    
    public static String removeDuplicateLetters(String withDupes, HashMap<String, String> conversions) {
        String firstLetter = "" + withDupes.charAt(0);
        String firstLetterConvert = firstLetter.toLowerCase();
        for (String key : conversions.keySet()) {
            firstLetterConvert = firstLetterConvert.replaceAll("[" + key + "]", conversions.get(key));
        }
        
        String noDupes = "";
        for (int i = 1; i < withDupes.length(); i++) {
            String currentLetter = "" + withDupes.charAt(i);
            if (i == 1 && !(currentLetter.equals(firstLetter))) {
                noDupes += firstLetter;
            } else {
                String previousLetter = "" + withDupes.charAt(i - 1);
                if (!currentLetter.equals(previousLetter)) {
                    noDupes += currentLetter;
                }
            }
        }
        return noDupes;
    }
}