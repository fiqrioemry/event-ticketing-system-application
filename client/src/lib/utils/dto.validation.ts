import { type ValidationRules } from './validation';

// Login Request
export const loginValidationRules: ValidationRules = {
	email: {
		required: true,
		email: true,
		message: 'Please enter a valid email address'
	},
	password: {
		required: true,
		minLength: 6,
		message: 'Password must be at least 6 characters'
	}
};

export const purchaseValidationRules: ValidationRules = {
	email: {
		required: true,
		email: true,
		message: 'Please enter a valid email address'
	},
	phone: {
		required: true,
		minLength: 10,
		maxLength: 15,
		pattern: /^\+?[0-9]\d{0,15}$/,
		message: 'Please enter a valid phone number'
	},
	fullname: {
		required: true,
		minLength: 6,
		message: 'Full name must be at least 6 characters'
	}
};

// Register Request
export const registerValidationRules: ValidationRules = {
	fullname: {
		required: true,
		minLength: 2,
		maxLength: 100,
		message: 'Full name must be between 2-100 characters'
	},
	email: {
		required: true,
		email: true,
		message: 'Please enter a valid email address'
	},
	password: {
		required: true,
		minLength: 6,
		message: 'Password must be at least 6 characters'
	}
};

// Update User Request
export const updateUserValidationRules: ValidationRules = {
	fullname: {
		required: true,
		minLength: 2,
		maxLength: 100,
		message: 'Full name must be between 2-100 characters'
	}
};

// Change Password Request
export const changePasswordValidationRules: ValidationRules = {
	currentPassword: {
		required: true,
		minLength: 6,
		message: 'Current password must be at least 6 characters'
	},
	newPassword: {
		required: true,
		minLength: 6,
		message: 'New password must be at least 6 characters'
	},
	confirmPassword: {
		required: true,
		minLength: 6,
		message: 'Confirm password must be at least 6 characters'
	}
};

// Forgot Password Request
export const forgotPasswordValidationRules: ValidationRules = {
	email: {
		required: true,
		email: true,
		message: 'Please enter a valid email address'
	}
};

// Reset Password Request
export const resetPasswordValidationRules: ValidationRules = {
	token: {
		required: true,
		message: 'Reset token is required'
	},
	newPassword: {
		required: true,
		minLength: 6,
		message: 'New password must be at least 6 characters'
	},
	confirmPassword: {
		required: true,
		minLength: 6,
		message: 'Confirm password must be at least 6 characters'
	}
};
