async function addTwoPromises(
	promise1: Promise<number>,
	promise2: Promise<number>
): Promise<number> {
	const result1: number = await promise1;
	const result2: number = await promise2;
	return result1 + result2;
}

async function addTwoPromisesPromise(
	promise1: Promise<number>,
	promise2: Promise<number>
): Promise<number> {
	return Promise.all([promise1, promise2]).then((numbers) => {
		return numbers.reduce(
			(accumulator, currentValue) => accumulator + currentValue,
			0
		);
	});
}

addTwoPromises(Promise.resolve(2), Promise.resolve(2)).then(console.log); // 4
