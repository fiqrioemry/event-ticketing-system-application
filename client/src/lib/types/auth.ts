// lib/types/auth.ts - Auth specific types
import type { ApiResponse } from '$lib/types/api';

export interface UserResponse {
	id: string;
	fullname: string;
	email: string;
	avatar: string | null;
	role: 'admin' | 'user' | 'moderator';
	joinedAt: string;
}

export interface LoginRequest {
	email: string;
	password: string;
	rememberMe?: boolean;
}

export interface RegisterRequest {
	email: string;
	fullname: string;
	password: string;
}

export type AuthResponse = ApiResponse<{ user: UserResponse }>;
