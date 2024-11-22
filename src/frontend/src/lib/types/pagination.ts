export interface Pagination<T> {
	count: number;
	rowsPerPage: number;
	currentPage: number;
	pageBuffer: number;
	maxPage: number;
	data: Array<T>;
	setPage: () => void;
	prevPage: () => void;
	nextPage: () => void;
	hasNextPage: () => boolean;
	hasPrevPage: () => boolean;
	displayPage: () => Array<T>;
}
