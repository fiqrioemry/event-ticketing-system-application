// lib/services/auth.service.ts
import { publicInstance, authInstance } from '$lib/services/client';
import type { LoginRequest, RegisterRequest } from '$lib/types/api';

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

// POST /api/auth/send-otp
export const sendOTP = async (data: { email: string }) => {
	const res = await publicInstance.post('/auth/send-otp', data);
	return res.data;
};

// POST /api/auth/verify-otp
export const verifyOTP = async (data: { email: string; otp: string }) => {
	const res = await publicInstance.post('/auth/verify-otp', data);
	return res.data;
};
