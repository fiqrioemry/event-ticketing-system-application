// lib/stores/auth.store.ts - Enhanced with better state management
import { writable, derived } from 'svelte/store';
import type { LoginRequest, RegisterRequest, UserResponse } from '$lib/types/auth';

// Form states
export const loginForm = writable<LoginRequest>({
	email: '',
	password: '',
	rememberMe: false
});

export const registerForm = writable<RegisterRequest>({
	email: '',
	fullname: '',
	password: ''
});

export const isLoading = writable<boolean>(false);
export const error = writable<string | null>(null);
export const user = writable<UserResponse | null>(null);

// Helper functions
export const authStore = {
	resetForms: () => {
		loginForm.set({ email: '', password: '', rememberMe: false });
		registerForm.set({ email: '', fullname: '', password: '' });
	},

	setUser: (userData: UserResponse) => {
		user.set(userData);
	},

	clearError: () => {
		error.set(null);
	},

	setLoading: (loading: boolean) => {
		isLoading.set(loading);
	},

	setError: (errorMessage: string) => {
		error.set(errorMessage);
	}
};
