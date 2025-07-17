// lib/services/auth.service.ts
import type {
	LoginRequest,
	RegisterRequest,
	ForgotPasswordRequest,
	ResetPasswordRequest
} from '$lib/types/api';
import { publicInstance, authInstance } from '$lib/services/client';

// POST /api/auth/login
export const login = async (credentials: LoginRequest) => {
	const res = await publicInstance.post('/auth/login', credentials);
	return res.data;
};

// POST /api/auth/register
export const register = async (userData: RegisterRequest) => {
	const res = await publicInstance.post('/auth/register', userData);
	return res.data;
};

// POST /api/auth/logout
export const logout = async () => {
	const res = await authInstance.post('/auth/logout');
	return res.data;
};

// POST /api/auth/refresh-token
export const refreshToken = async () => {
	const res = await publicInstance.post('/auth/refresh-token');
	return res.data;
};

// POST /api/auth/resend-otp
export const resendOTP = async (data: { email: string }) => {
	const res = await publicInstance.post('/auth/resend-otp', data);
	return res.data;
};

// POST /api/auth/verify-otp
export const verifyOTP = async (data: { email: string; otp: string }) => {
	const res = await publicInstance.post('/auth/verify-otp', data);
	return res.data;
};

// POST /api/auth/forgot-password
export const forgotPassword = async (data: ForgotPasswordRequest) => {
	const res = await publicInstance.post('/auth/forgot-password', data);
	return res.data;
};

// GET /api/auth/validate-reset-token
export const validateResetToken = async (token: string) => {
	const res = await publicInstance.get('/auth/validate-reset-token', {
		params: { token }
	});
	return res.data;
};

// POST /api/auth/reset-password
export const resetPassword = async (data: ResetPasswordRequest) => {
	const res = await publicInstance.post('/auth/reset-password', data);
	return res.data;
};

// GET /api/auth/google - Redirect to Google OAuth
export const googleOAuthRedirect = (): string => {
	return `${publicInstance.defaults.baseURL}/auth/google`;
};

// GET /api/auth/google/callback
export const googleOAuthCallback = async (code: string, state?: string) => {
	const res = await publicInstance.get('/auth/google/callback', {
		params: { code, state }
	});
	return res.data;
};
