import type { Pagination } from '$lib/types';
export function createPagination<T>(input: Array<T>, maxItemsPerRow: number): Pagination<T> {
	let data = $state(input);
	let count = $derived(data.length);
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

	let setPage = () => {
		if (pageBuffer > maxPage) {
			pageBuffer = maxPage;
		} else if (pageBuffer < 1) {
			pageBuffer = 1;
		}
		currentPage = pageBuffer;
	};

	let displayPage = () => {
		const start = (currentPage - 1) * rowsPerPage;
		const end = start + rowsPerPage;
		let pageData = data.slice(start, end);

		return pageData;
	};

	return {
		get count() {
			return count;
		},
		get rowsPerPage() {
			return rowsPerPage;
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
		get data() {
			return data;
		},
		setPage,
		prevPage,
		nextPage,
		hasNextPage,
		hasPrevPage,
		displayPage
	};
}
