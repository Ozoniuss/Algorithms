function jumpdp(nums: Array<number>): number {

    if (nums.length == 0) {
        return -1;
    }

    let maxJumpIdx: Array<number> = [];
    maxJumpIdx.push(nums[0]);

    for (let i = 1; i<nums.length;i++) {
        maxJumpIdx[i] = Math.max(nums[i]+i, maxJumpIdx[i-1]);
    }

    let curri = 0;
    let ans = 0;

    while (curri < nums.length - 1) {
        ans++;
        curri = maxJumpIdx[curri];
    }
    return ans
}

console.log(jumpdp([2,3,1,1,4]));