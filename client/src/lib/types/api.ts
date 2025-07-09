export interface ApiResponse<T = any> {
	success: boolean;
	message: string;
	data: T | null;
	meta: {
		timestamp: string;
		requestId: string;
	} | null;
	errors: Record<string, string[]> | null;
	pagination: {
		page: number;
		limit: number;
		total: number;
		totalPages: number;
	} | null;
}
export interface ErrorResponse {
	success: false;
	message: string;
	errors: Record<string, string[]> | null;
}
