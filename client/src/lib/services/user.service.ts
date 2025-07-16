// lib/services/user.service.ts
import { authInstance } from '$lib/services/client';
import type { UpdateProfileRequest, User, ApiResponse } from '$lib/types/api';

// GET /api/user/ - Get all users (admin only)
export const getAllUsers = async (): Promise<ApiResponse<User[]>> => {
	const res = await authInstance.get('/user/');
	return res.data;
};

// GET /api/user/:id - Get user detail by ID (admin only)
export const getUserDetail = async (id: string | number): Promise<ApiResponse<User>> => {
	const res = await authInstance.get(`/user/${id}`);
	return res.data;
};

// GET /api/user/me - Get my profile
export const getMyProfile = async (): Promise<ApiResponse<User>> => {
	const res = await authInstance.get('/user/me');
	return res.data;
};

// PUT /api/user/me - Update my profile
export const updateProfile = async (
	profileData: UpdateProfileRequest
): Promise<ApiResponse<User>> => {
	const res = await authInstance.put('/user/me', profileData);
	return res.data;
};
