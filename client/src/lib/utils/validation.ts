// lib/utils/validation.ts - Enhanced validators with dirty state
export const validators = {
	email: (email: string): string | null => {
		if (!email) return 'Email is required';
		const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
		if (!emailRegex.test(email)) return 'Please enter a valid email';
		return null;
	},

	password: (password: string): string | null => {
		if (!password) return 'Password is required';
		if (password.length < 6) return 'Password must be at least 6 characters';
		return null;
	},

	fullname: (name: string): string | null => {
		if (!name) return 'Full name is required';
		if (name.length < 2) return 'Name must be at least 2 characters';
		return null;
	}
};

export const shouldShowError = (
	error: string | null,
	touched: boolean,
	dirty: boolean
): boolean => {
	return error !== null && (touched || dirty);
};
