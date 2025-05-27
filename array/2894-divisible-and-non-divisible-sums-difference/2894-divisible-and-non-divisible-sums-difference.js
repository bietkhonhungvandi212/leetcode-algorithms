/**
 * @param {number} n
 * @param {number} m
 * @return {number}
 */
 var differenceOfSums = function(n, m) {
	let sum = 0
	for (let i = 1; i <= n; i++) {
		if (i % m !== 0) {
			sum += i;
		}
	}

	for (let i = 1; i <= n; i++) {
		if (i % m === 0) {
			sum -= i;
		}
	}
    
	return sum
};