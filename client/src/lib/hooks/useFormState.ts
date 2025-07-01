import { writable, derived, type Readable } from 'svelte/store';
import { ZodSchema } from 'zod';

/** === TYPE === */
export type FormState<T> = {
	values: T;
	errors: Partial<Record<keyof T, string>>;
	touched: Partial<Record<keyof T, boolean>>;
	isValid: boolean;
	isDirty: boolean;
	isLoading?: boolean;
};

/** === STORE CREATOR === */
export function createFormStore<T extends Record<string, any>>(
	schema: ZodSchema<T>,
	defaultValues: T
) {
	const { subscribe, update, set } = writable<FormState<T>>({
		values: { ...defaultValues },
		errors: {},
		touched: {},
		isValid: false,
		isDirty: false,
		isLoading: false
	});

	return {
		subscribe,
		setValues(values: T) {
			const result = schema.safeParse(values);
			const errors: Partial<Record<keyof T, string>> = {};
			const touched: Partial<Record<keyof T, boolean>> = {};

			if (!result.success) {
				for (const err of result.error.errors) {
					const key = err.path[0] as keyof T;
					errors[key] = err.message;
				}
			}

			for (const key in values) {
				if (values[key] !== defaultValues[key]) {
					touched[key as keyof T] = true;
				}
			}

			const isDirty = JSON.stringify(values) !== JSON.stringify(defaultValues);

			update((prev) => ({
				...prev,
				values,
				errors,
				touched,
				isValid: Object.keys(errors).length === 0,
				isDirty
			}));
		},
		setLoading(value: boolean) {
			update((prev) => ({ ...prev, isLoading: value }));
		},
		reset() {
			set({
				values: { ...defaultValues },
				errors: {},
				touched: {},
				isValid: false,
				isDirty: false,
				isLoading: false
			});
		}
	};
}

/** === STATE SELECTOR === */
export function useFormState<T extends Record<string, any>>(formStore: {
	subscribe: Readable<FormState<T>>['subscribe'];
	setValues: (values: T) => void;
}) {
	return derived(formStore, ($form) => {
		const proxyValues = new Proxy(structuredClone($form.values), {
			set(target, prop: string | symbol, value: any) {
				const key = prop.toString() as keyof T;
				target[key] = value;
				formStore.setValues(target);
				return true;
			}
		});

		return {
			form: proxyValues,
			errors: $form.errors,
			touched: $form.touched,
			isValid: $form.isValid,
			isDirty: $form.isDirty,
			isLoading: $form.isLoading
		};
	});
}
