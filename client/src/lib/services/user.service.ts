// lib/services/user.service.ts
import { authInstance } from '$lib/services/client';
import type { UpdateProfileRequest, ChangePasswordRequest } from '$lib/types/api';
import { buildFormData } from '$lib/utils/formatter';

// GET /api/user/me - Get my profile
export const getMyProfile = async () => {
	const res = await authInstance.get('/user/me');
	return res.data;
};

// PUT /api/user/me - Update my profile
export const updateProfile = async (profileData: UpdateProfileRequest) => {
	const formData = buildFormData(profileData);
	const res = await authInstance.put('/user/me', formData);
	return res.data;
};

// GET /api/user/change-password - Get my profile
export const changePassword = async (data: ChangePasswordRequest) => {
	const res = await authInstance.put('/user/change-password', data);
	return res.data;
};
