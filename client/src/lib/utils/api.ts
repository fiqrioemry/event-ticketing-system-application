// lib/utils/api.ts
import type { ApiResponse, ErrorCode } from '$lib/types/api';

export const handleApiCall = async <T>(apiCall: () => Promise<any>): Promise<ApiResponse<T>> => {
	try {
		const response = await apiCall();
		return response.data;
	} catch (error: any) {
		if (error.response?.data) {
			return error.response.data;
		}
		return {
			success: false,
			message: error.message || 'Something went wrong',
			code: 'INTERNAL_SERVER_ERROR' as ErrorCode
		};
	}
};
