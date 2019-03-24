package Sort;

/**
 * Created by wanjx on 2019/3/23.
 */

public class MergeSort {
    public void merge(int[] nums, int lo, int mid, int hi) {
        // int lo = 0, hi = nums.length - 1;
        int i = lo, j = mid + 1;

        int[] aux = new int[nums.length];
        for (int k = 0; k < nums.length; k++) {
            aux[k] = nums[k];
        }

        for (int k = lo; k <= hi; k++) {
            if (i > mid) nums[k] = aux[j++];
            else if (j > hi) nums[k] = aux[i++];
            else if (aux[i] > aux[j]) nums[k] = aux[j++];
            else nums[k] = aux[i++];
        }
    }

    public void sort(int[] nums) {
        int lo = 0, hi = nums.length - 1;
        sort(nums, lo, hi);
    }
    public void sort(int[] nums, int lo, int hi) {
        if (lo >= hi) {
            return;
        }

        int mid = (lo + hi) / 2;
        sort(nums, lo, mid);
        sort(nums, mid + 1, hi);
        merge(nums, lo, mid, hi);
    }

    public void sortBU(int[] nums) {
        int n = nums.length;

        for (int len = 1; len < n; len *= 2) {
            for (int lo = 0; lo < n - len; lo += len + len) {
                int mid = lo + len - 1;
                int hi = Math.min(lo + len + len - 1, n - 1);
                merge(nums, lo, mid, hi);
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
        MergeSort M = new MergeSort();
        // M.merge(a, 2);
        M.string(a);
        // M.sort(a);
        M.sortBU(a);
        M.string(a);
    }
}
