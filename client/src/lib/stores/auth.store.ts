import type {
	User,
	AuthState,
	LoginRequest,
	RegisterRequest,
	VerifyOTPRequest,
	ForgotPasswordRequest,
	ResetPasswordRequest,
	GoogleOAuthData
} from '$lib/types/api';
import { toast } from 'svelte-sonner';
import { goto } from '$app/navigation';
import { writable, derived } from 'svelte/store';
import * as auth from '$lib/services/auth.service';
import { createStoreActions } from '$lib/utils/store';

const initialAuthState: AuthState = {
	error: null,
	user: null,
	isLoading: false,
	isAuthenticated: false
};

function createAuthStore() {
	const { subscribe, update, set } = writable<AuthState>(initialAuthState);
	const actions = createStoreActions<User>('user', 'user');

	return {
		subscribe,

		async login(credentials: LoginRequest) {
			this.setLoading(true);

			try {
				const response = await auth.login(credentials);

				this.setUser(response.data);
				toast.success(response.message || 'Login successful');
				goto('/user/profile');
				return response;
			} catch (error: any) {
				this.setError(error, 'Login failed');
				throw error;
			}
		},

		async register(userData: RegisterRequest) {
			this.setLoading(true);

			try {
				const response = await auth.register(userData);
				this.setLoading(false);
				toast.success(response.message || 'Please check your email to verify your account');
				return response;
			} catch (error: any) {
				this.setError(error, 'Registration failed');
				throw error;
			}
		},

		async resendOtp(email: string) {
			this.setLoading(true);
			try {
				const response = await auth.resendOTP({ email });
				this.setLoading(false);
				toast.success(response.message || 'OTP sent successfully');
			} catch (error: any) {
				this.setError(error, 'Failed to resend OTP');
				throw error;
			}
		},

		async verifyOtp(otpRequest: VerifyOTPRequest) {
			this.setLoading(true);
			try {
				const response = await auth.verifyOTP(otpRequest);

				this.setUser(response.data);
				toast.success(response.message || 'Registration successful');
				goto('/user/profile');
			} catch (error: any) {
				this.setError(error, 'Failed to verify OTP');
				throw error;
			}
		},
		async logout(redirectTo: string = '/signin') {
			this.setLoading(true);
			try {
				const response = await auth.logout();
				toast.success(response.message || 'Logged out successfully');
			} catch (error: any) {
				console.error(error.response?.data?.message || 'Logout failed');
			} finally {
				this.setUser(null);
				goto(redirectTo);
			}
		},

		async forgotPassword(forgotPasswordData: ForgotPasswordRequest) {
			this.setLoading(true);

			try {
				const response = await auth.forgotPassword(forgotPasswordData);
				this.setLoading(false);
				toast.success(response.message || 'Password reset link sent to your email');
				return response;
			} catch (error: any) {
				this.setError(error, 'Failed to send reset link');
				toast.error('Failed to send reset link');
				throw error;
			}
		},

		async validateResetToken(token: string) {
			this.setLoading(true);

			try {
				const response = await auth.validateResetToken(token);
				this.setLoading(true);
				return response;
			} catch (error: any) {
				this.setError(error, 'Invalid or expired reset token');
				throw error;
			}
		},

		async resetPassword(resetPasswordData: ResetPasswordRequest) {
			this.setLoading(true);

			try {
				const response = await auth.resetPassword(resetPasswordData);

				this.setUser(response.data);
				toast.success(response.message || 'Password reset successful');
				goto('/user/profile');
				return response;
			} catch (error: any) {
				this.setError(error, 'Failed to reset password');
				throw error;
			}
		},

		async googleOAuthLogin(oauthData: GoogleOAuthData) {
			this.setLoading(true);

			try {
				const response = await auth.googleOAuthCallback(oauthData.code, oauthData.state);
				this.setUser(response.data);
				toast.success(response.message || 'Google login successful');
				goto('/user/profile');
			} catch (error: any) {
				this.setError(error, 'Google login failed');
			}
		},

		getGoogleOAuthUrl(): string {
			return auth.googleOAuthRedirect();
		},

		async checkAuth() {
			this.setLoading(true);

			try {
				const response = await auth.refreshToken();
				this.setUser(response.data);
				return response;
			} catch (error: any) {
				this.reset();
				throw error;
			}
		},

		async refreshSession() {
			try {
				const response = await auth.refreshToken();
				this.setUser(response.data);
				return response;
			} catch (error: any) {
				this.reset();
				throw error;
			}
		},

		setUser: (user: User | null) => {
			update((state: AuthState) => ({
				...state,
				user,
				isAuthenticated: !!user,
				error: null,
				isLoading: false
			}));
		},

		setError: (error: any, fallback: string) => {
			const message = error.response?.data?.message || fallback;
			actions.setError(update, error);
			console.error(message);
		},

		clearError: () => {
			actions.clearError(update);
		},

		setLoading: (isLoading: boolean) => {
			actions.setLoading(update, isLoading);
		},

		reset: () => {
			actions.reset(update, initialAuthState);
		}
	};
}

export const authStore = createAuthStore();
export const authUser = derived(authStore, ($state) => $state.user);
export const authError = derived(authStore, ($state) => $state.error);
export const authLoading = derived(authStore, ($state) => $state.isLoading);
export const isAuthenticated = derived(authStore, ($state) => $state.isAuthenticated);
