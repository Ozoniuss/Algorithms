class Solution {
    public int search(int[] nums, int target) {
        if (nums.length == 0) {
            return -1;
        }

        int left_index = 0;
        int right_index = nums.length;

        while (true) {
            if (left_index > right_index) {
                return -1;
            }
            int midpoint = (left_index + right_index) / 2;

            if (midpoint == nums.length) {
                return -1;
            }

            if (nums[midpoint] > target) {
                right_index = midpoint - 1;
            }

            if (nums[midpoint] < target) {
                left_index = midpoint + 1;
            }

            if (nums[midpoint] == target) {
                return midpoint;
            }
        }
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] nums = { -1, 0, 3, 5, 9, 12 };
        int x = s.search(nums, 13);
        System.out.println(x);
    }
}