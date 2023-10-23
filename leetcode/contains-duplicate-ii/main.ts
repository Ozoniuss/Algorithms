function containsNearbyDuplicate(nums: number[], k: number): boolean {
	if (nums.length < 2) {
		return false;
	}
	for (let j = 1; j <= k; j++) {
		for (let i = 0; i < nums.length - j; i++) {
			if ((nums[i] - nums[i + j]) * (nums[i + j] - nums[i]) === 0) {
				return true;
			}
		}
	}
	return false;
}

function containsNearbyDuplicateMap(nums: number[], k: number): boolean {
	if (nums.length < 2) {
		return false;
	}
	const positions = new Map<number, number>();
	for (let idx = 0; idx < nums.length; idx++) {
		const num = nums[idx];
		const pos = positions.get(num);
		if (pos === undefined) {
			positions.set(num, idx);
		} else if (idx - pos <= k) {
			return true;
		} else {
			positions.set(num, idx);
		}
	}

	return false;
}

console.log(containsNearbyDuplicateMap([1, 2, 3, 1], 3));
