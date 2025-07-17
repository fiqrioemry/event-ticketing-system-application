export const pagination = {
	page: 1,
	limit: 10,
	total: 0,
	totalPages: 1
};

export const locationOptions = [
	{ value: '', label: 'All Locations' },
	{ value: 'jakarta', label: 'Jakarta' },
	{ value: 'surabaya', label: 'Surabaya' },
	{ value: 'bandung', label: 'Bandung' },
	{ value: 'yogyakarta', label: 'Yogyakarta' },
	{ value: 'bali', label: 'Bali' }
];

export const sortOptions = [
	{ value: 'date_asc', label: 'Created Date' },
	{ value: 'date_desc', label: 'Created Date (Desc)' },
	{ value: 'title_asc', label: 'Title A-Z' },
	{ value: 'title_desc', label: 'Title Z-A' }
];

export const statusOptions = [
	{ value: 'pending', label: 'Pending' },
	{ value: 'paid', label: 'Paid' },
	{ value: 'failed', label: 'Failed' },
	{ value: 'refunded', label: 'Refunded' }
];
