class ListNode {
	val: number;
	next: ListNode | null;
	constructor(val?: number, next?: ListNode | null) {
		this.val = val === undefined ? 0 : val;
		this.next = next === undefined ? null : next;
	}
}

function hasCycleWithSet(head: ListNode | null): boolean {
	const explored = new Set();
	if (!head) {
		return false;
	}
	let current: ListNode = head;

	while (true) {
		// hack: use the "current" pointer instead of the value, which then
		// requires having to find a continuous cycle of values...
		if (explored.has(current)) {
			return true;
		}
		explored.add(current);
		if (!current.next) {
			return false;
		}
		current = current.next;
	}
}

function hasCycleWithRunnerAndCatcher(head: ListNode | null): boolean {
	if (!head) {
		return false;
	}

	let walker: ListNode = head;
	let runner: ListNode = head;

	while (true) {
		if (runner.next) {
			walker = walker.next!;
			runner = runner.next;
		}
		if (runner.next) {
			runner = runner.next;
		}
		if (!runner.next) {
			return false;
		}
		if (walker === runner) {
			return true;
		}
	}
}
