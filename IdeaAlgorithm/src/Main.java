import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Main {

/** 请完成下面这个函数，实现题目要求的功能 **/
    /** 当然，你也可以不按照这个模板来作答，完全按照自己的想法来 ^-^  **/
    static String calculate(int m, int k) {

        return String.valueOf(getPigId(m)) + "," + String.valueOf(+ getYear(m));
    }



    public static void main(String[] args){
        Scanner in = new Scanner(System.in);
        String[] line = in.nextLine().split(",");
        int m = Integer.valueOf(line[0]);
        int k = Integer.valueOf(line[1]);
        System.out.println(calculate(m, k));

    }

    static  int getYear(int m) {
        int n = 0;
        int year = 2019;
        if (m > 0 && m < 4) {
            return year;
        }
        if (m == 4) {
            return year + 1;
        }
        int f1 = 3;
        int f2 = 2;
        int f3 = 4;
        while (n < m) {
             n = f1 + f3;
             f1 = f2;
             f2 = f3;
             f3 = n;
             year++;
        }
        return year;
    }
    static int getPigId(int m) {
        ArrayList<String> Ids;

        int ret;


        //Ids = ArrayList<>();
        ret = 0;
        if (m == 1) {
            return 2;
        }

        if (m == 2) {
            return 3;
        }

        if (m == 3) {
            return 4;
        }

        int f2 = 3;
        int f3 = 2;
        int f1 = 4;
        for (int i = 3; i < m; i++) {
            ret = f2 + f3;
            //Ids.add(i, String.valueOf(ret));
            f3 = f2;
            f2 = f1;
            f1 = ret;
        }
        return ret;
    }

    static String reverse(String s) {
        String[] ss = s.split("");
        for (int i = 0; i < ss.length; i++) {
            String t = ss[i];
            ss[i] = ss[ss.length - 1];
            ss[ss.length - 1] = t;
        }
        s = String.valueOf(ss);
        return s;
    }

}
