// lib/utils/store.ts

// Unified store actions untuk semua kebutuhan
export const createStoreActions = <T extends { id: string }>(
	itemsKey: string = 'items',
	currentItemKey: string = 'currentItem'
) => ({
	// Loading states
	setLoading: (update: any, isLoading: boolean) => {
		update((state: any) => ({ ...state, isLoading }));
	},

	setCreating: (update: any, isCreating: boolean) => {
		update((state: any) => ({ ...state, isCreating }));
	},

	setUpdating: (update: any, isUpdating: boolean) => {
		update((state: any) => ({ ...state, isUpdating }));
	},

	setDeleting: (update: any, isDeleting: boolean) => {
		update((state: any) => ({ ...state, isDeleting }));
	},

	resetLoading: (update: any) => {
		update((state: any) => ({
			...state,
			isLoading: false,
			isCreating: false,
			isUpdating: false,
			isDeleting: false
		}));
	},

	// Error handling
	setError: (update: any, error: any) => {
		update((state: any) => ({
			...state,
			error,
			isLoading: false,
			isCreating: false,
			isUpdating: false,
			isDeleting: false
		}));
	},

	clearError: (update: any) => {
		update((state: any) => ({ ...state, error: null }));
	},

	// Success operations - auto handle loading states
	setItems: (update: any, items: T[], pagination?: any) => {
		update((state: any) => ({
			...state,
			[itemsKey]: items,
			pagination: pagination || state.pagination,
			error: null,
			isLoading: false
		}));
	},

	setCurrentItem: (update: any, item: T | null) => {
		update((state: any) => ({
			...state,
			[currentItemKey]: item,
			error: null,
			isLoading: false
		}));
	},

	addItem: (update: any, newItem: T) => {
		update((state: any) => ({
			...state,
			[itemsKey]: [newItem, ...state[itemsKey]],
			pagination: state?.pagination
				? {
						...state.pagination,
						totalItems: state.pagination.totalItems + 1,
						totalPages: Math.ceil((state.pagination.totalItems + 1) / state.pagination.limit)
					}
				: undefined,
			error: null,
			isCreating: false
		}));
	},

	updateItem: (update: any, updatedItem: T, id: string) => {
		update((state: any) => ({
			...state,
			[itemsKey]: state[itemsKey].map((item: T) => (item.id === id ? updatedItem : item)),
			[currentItemKey]: state[currentItemKey]?.id === id ? updatedItem : state[currentItemKey],
			error: null,
			isUpdating: false
		}));
	},

	removeItem: (update: any, id: string) => {
		update((state: any) => ({
			...state,
			[itemsKey]: state[itemsKey].filter((item: T) => item.id !== id),
			pagination: state.pagination
				? {
						...state.pagination,
						totalItems: Math.max(0, state.pagination.totalItems - 1),
						totalPages: Math.ceil(
							Math.max(0, state.pagination.totalItems - 1) / state.pagination.limit
						)
					}
				: undefined,
			[currentItemKey]: state[currentItemKey]?.id === id ? null : state[currentItemKey],
			error: null,
			isDeleting: false
		}));
	},

	// Reset state
	reset: (update: any, initialState: any) => {
		update(() => initialState);
	}
});
