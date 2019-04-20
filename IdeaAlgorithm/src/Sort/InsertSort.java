package Sort;

/**
 * Created by wanjx on 2019/3/23.
 */

public class InsertSort {
    public void sort(int[] nums) {
        for (int i = 1; i < nums.length; i++) {
            for(int j = i; j > 0 && nums[j] < nums[j - 1]; j--) {
                int tmp = nums[j];
                nums[j] = nums[j-1];
                nums[j-1] = tmp;
            }
        }
    }

    public void string(int[] a) {
        if (a == null) {
            return;
        }
        for (int i = 0; i < a.length; i++) {
            System.out.print(a[i] + "\t");
        }
        System.out.print("\n");
    }

    public static void main(String[] args) {
        int[] a = {4, 12, 60, 1, 233, 3};
        InsertSort S = new InsertSort();
        // S.merge(a, 2);
        S.string(a);
        S.sort(a);
        S.string(a);
    }
}
