import { toast } from 'svelte-sonner';
import { goto } from '$app/navigation';
import { writable, derived } from 'svelte/store';
import * as auth from '$lib/services/auth.service';
import { createStoreActions } from '$lib/utils/store';
import type { User, LoginRequest, RegisterRequest, AuthState } from '$lib/types/api';

const initialAuthState: AuthState = {
	error: null,
	user: null,
	isLoading: false,
	isAuthenticated: false
};

const setUserInStore = (update: any, user: User | null) => {
	update((state: AuthState) => ({
		...state,
		user,
		isAuthenticated: !!user,
		error: null,
		isLoading: false
	}));
};

function createAuthStore() {
	const { subscribe, update, set } = writable<AuthState>(initialAuthState);
	const actions = createStoreActions<User>('user', 'user');

	return {
		subscribe,

		async login(credentials: LoginRequest, redirectTo: string = '/') {
			actions.clearError(update);
			actions.setLoading(update, true);

			try {
				const response = await auth.login(credentials);
				const user = response.data;
				setUserInStore(update, user);
				toast.success(response.message || 'Login successful');
				goto(redirectTo);
				return response;
			} catch (error: any) {
				const message = error.response?.data?.message || error.message || 'Login failed';
				actions.setError(update, { message });
				toast.error(message);

				throw error; // Re-throw for component handling if needed
			}
		},

		async register(userData: RegisterRequest, redirectTo: string = '/dashboard') {
			actions.clearError(update);
			actions.setLoading(update, true);

			try {
				const response = await auth.register(userData);
				const user = response.data;
				setUserInStore(update, user);
				toast.success(response.message || 'Registration successful');
				goto(redirectTo);
				return response;
			} catch (error: any) {
				const message = error.response?.data?.message || error.message || 'Registration failed';
				actions.setError(update, { message });
				toast.error(message);
				throw error;
			}
		},
		async logout(redirectTo: string = '/signin') {
			actions.setLoading(update, true);
			try {
				const response = await auth.logout();
				toast.success(response.message || 'Logged out successfully');
			} catch (error: any) {
				toast.warning(error.response);
			} finally {
				actions.reset(update, initialAuthState);
				goto(redirectTo);
			}
		},

		async checkAuth() {
			actions.setLoading(update, true);

			try {
				const response = await auth.refreshToken();
				const user = response.data;
				setUserInStore(update, user);
				return response;
			} catch (error: any) {
				actions.reset(update, initialAuthState);
				throw error;
			}
		},

		async refreshSession() {
			try {
				const response = await auth.refreshToken();
				const user = response.data;
				setUserInStore(update, user);
				return response;
			} catch (error: any) {
				actions.reset(update, initialAuthState);
				throw error;
			}
		},

		setUser: (user: User | null) => {
			setUserInStore(update, user);
		},

		clearError: () => {
			actions.clearError(update);
		},

		reset: () => {
			actions.reset(update, initialAuthState);
		}
	};
}

export const authStore = createAuthStore();

// Derived stores
export const currentUser = derived(authStore, ($state) => $state.user);
export const authError = derived(authStore, ($state) => $state.error);
export const authLoading = derived(authStore, ($state) => $state.isLoading);
export const isAuthenticated = derived(authStore, ($state) => $state.isAuthenticated);
