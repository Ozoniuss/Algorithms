function twoSum(nums: number[], target: number): number[] {
    for (let i = 0; i < nums.length - 1; i++) {
        for (let j = i + 1; j < nums.length; j++) {
            if (nums[i] + nums[j] === target) {
                return [i, j]
            }
        }
    }
    return []
};


function twoSumSet(nums: number[], target: number): number[] {
    let numValues = new Map<number, number[]>()

    for (const [index, num] of nums.entries()) {
        if (numValues.has(num)) {
            let existing = numValues.get(num)!
            existing.push(index)
        } else {
            numValues.set(num, [index])
        }
    }


    for (const [key, val] of numValues) {
        if (numValues.has(target - key)) {
            if (val.length === 2) {
                return val
            } else if (target - key !== key) {
                return [val[0], numValues.get(target - key)![0]]
            }
        }
    }
    return [0, 0]
};

console.log(twoSumSet([3, 2, 4], 6))

