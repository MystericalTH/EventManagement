export function createPagination(data, maxItemsPerRow: number) {
	let requestList = $state(data);
	let count = $derived(requestList.length);
	let currentPage = $state(1);
	let rowsPerPage = $state(maxItemsPerRow);
	let maxPage = $derived(Math.ceil(count / maxItemsPerRow));
	let pageBuffer = $state(1);

	let hasNextPage = () => {
		return currentPage < maxPage;
	};
	let hasPrevPage = () => {
		return currentPage > 1;
	};

	let nextPage = () => {
		if (hasNextPage()) {
			currentPage += 1;
			pageBuffer = currentPage;
		}
	};

	let prevPage = () => {
		if (hasPrevPage()) {
			currentPage -= 1;
			pageBuffer = currentPage;
		}
	};

	let displayPage = () => {
		const start = (currentPage - 1) * rowsPerPage;
		const end = start + rowsPerPage;
		let data = requestList.slice(start, end);

		return data;
	};

	return {
		get count() {
			return count;
		},
		get rowsPerPage() {
			return rowsPerPage;
		},
		get itemsInPage() {
			return itemsInPage;
		},
		get currentPage() {
			return currentPage;
		},
		set currentPage(value: number) {
			currentPage = value;
		},
		get pageBuffer() {
			return pageBuffer;
		},
		set pageBuffer(value: number) {
			pageBuffer = value;
		},

		get maxPage() {
			return maxPage;
		},
		prevPage,
		nextPage,
		hasNextPage,
		hasPrevPage,
		displayPage
	};
}
