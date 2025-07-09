// lib/stores/form.store.ts - Generic form dirty state management
import { writable, derived } from 'svelte/store';

export interface FormField {
	value: any;
	touched: boolean;
	dirty: boolean;
	error: string | null;
}

export interface FormState<T> {
	values: T;
	touched: Record<keyof T, boolean>;
	dirty: Record<keyof T, boolean>;
	errors: Record<keyof T, string | null>;
	isSubmitting: boolean;
}

export function createFormStore<T>(initialValues: T) {
	const initialState: FormState<T> = {
		values: { ...initialValues },
		touched: {} as Record<keyof T, boolean>,
		dirty: {} as Record<keyof T, boolean>,
		errors: {} as Record<keyof T, string | null>,
		isSubmitting: false
	};

	const { subscribe, set, update } = writable(initialState);

	// Derived states
	const isFormDirty = derived({ subscribe }, (state) => Object.values(state.dirty).some(Boolean));

	const isFormTouched = derived({ subscribe }, (state) =>
		Object.values(state.touched).some(Boolean)
	);

	const hasErrors = derived({ subscribe }, (state) =>
		Object.values(state.errors).some((error) => error !== null)
	);

	const isFormValid = derived(
		{ subscribe },
		(state) => !Object.values(state.errors).some((error) => error !== null)
	);

	return {
		subscribe,

		// Derived stores
		isFormDirty,
		isFormTouched,
		hasErrors,
		isFormValid,

		// Actions
		setFieldValue: (field: keyof T, value: any) => {
			update((state) => ({
				...state,
				values: { ...state.values, [field]: value },
				dirty: { ...state.dirty, [field]: true }
			}));
		},

		setFieldTouched: (field: keyof T, touched = true) => {
			update((state) => ({
				...state,
				touched: { ...state.touched, [field]: touched }
			}));
		},

		setFieldError: (field: keyof T, error: string | null) => {
			update((state) => ({
				...state,
				errors: { ...state.errors, [field]: error }
			}));
		},

		validateField: (field: keyof T, validator: (value: any) => string | null) => {
			update((state) => {
				const error = validator(state.values[field]);
				return {
					...state,
					errors: { ...state.errors, [field]: error }
				};
			});
		},

		reset: () => {
			set({
				values: { ...initialValues },
				touched: {} as Record<keyof T, boolean>,
				dirty: {} as Record<keyof T, boolean>,
				errors: {} as Record<keyof T, string | null>,
				isSubmitting: false
			});
		},

		setSubmitting: (isSubmitting: boolean) => {
			update((state) => ({ ...state, isSubmitting }));
		}
	};
}
