export interface Pagination {
	count: number;
	rowsPerPage: number;
	currentPage: number;
	pageBuffer: number;
	maxPage: number;
	data: Array<any>;
	setPage: () => void;
	prevPage: () => void;
	nextPage: () => void;
	hasNextPage: () => boolean;
	hasPrevPage: () => boolean;
	displayPage: () => Array<any>;
}
