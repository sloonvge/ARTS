package Search;

/**
 * Created by wanjx on 2019/3/23.
 */

public class BinarySearch {
    public int search(int[] nums, int t) {
        int lo = 0, hi = nums.length - 1;

        while (lo <= hi) {
            int mid = (lo + hi) / 2;

            if (nums[mid] == t) return mid;
            else if (nums[mid] > t) hi = mid -1;
            else lo = mid + 1;
        }

        return -1;
    }
}
