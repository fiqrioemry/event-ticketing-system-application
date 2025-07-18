// API SUCCESS AND ERROR RESPONSE TYPES
export enum ErrorCode {
	VALIDATION_ERROR = 'VALIDATION_ERROR',
	UNAUTHORIZED = 'UNAUTHORIZED',
	FORBIDDEN = 'FORBIDDEN',
	NOT_FOUND = 'NOT_FOUND',
	INTERNAL_SERVER_ERROR = 'INTERNAL_SERVER_ERROR',
	BAD_REQUEST = 'BAD_REQUEST',
	CONFLICT = 'CONFLICT'
}

export interface ErrorResponse {
	success: false;
	message: string;
	code: ErrorCode;
	errors?: Record<string, any>;
}

export interface SuccessResponse<T = any> {
	success: true;
	message: string;
	data: T;
	meta?: Meta;
}

export interface Pagination {
	page: number;
	limit: number;
	totalItems: number;
	totalPages: number;
}

export interface Meta {
	pagination?: Pagination;
	permissions?: Record<string, boolean>;
	flags?: Record<string, boolean>;
}

export interface Pagination {
	page: number;
	limit: number;
	total: number;
	totalPages: number;
}

export type ApiResponse<T = any> = SuccessResponse<T> | ErrorResponse;

// AUTHENTICATION & USER TYPES

export interface AuthState {
	isLoading: boolean;
	error: ErrorResponse | null;
	isAuthenticated: boolean;
	user: User | null;
}

export interface RegisterRequest {
	email: string;
	password: string;
	fullname: string;
}

export interface LoginRequest {
	email: string;
	password: string;
}

export interface AuthResponse {
	user: User;
	accessToken: string;
	refreshToken: string;
}

export interface ResendOTPRequest {
	email: string;
}

export interface VerifyOTPRequest {
	email: string;
	otp: string;
}

export interface ChangePasswordRequest {
	currentPassword: string;
	newPassword: string;
	confirmPassword: string;
}

export interface GoogleOAuthData {
	code: string;
	state?: string;
}

export interface ForgotPasswordRequest {
	email: string;
}

export interface ResetPasswordRequest {
	token: string;
	newPassword: string;
	confirmPassword: string;
}

export interface ForgotPasswordResponse {
	Message: string;
	Email: string;
}

export interface ResetTokenData {
	UserID: string;
	Email: string;
	ExpiresAt: Date;
	CreatedAt: Date;
}

export interface ForgotPasswordRequest {
	email: string;
}

export interface UpdateProfileRequest {
	fullname: string;
	avatar?: File;
	avatarURL?: string;
}

export interface UserQueryParams {
	search?: string;
	sort?: string;
	page?: number;
	limit?: number;
}

export interface User {
	id: string;
	email: string;
	role: string;
	fullname: string;
	balance?: number;
	avatar: string;
	joinedAt: string;
}

export interface UserProfile {
	id: string;
	email: string;
	role: string;
	fullname: string;
	balance?: number;
	avatar: string;
	joinedAt: string;
}

export interface Profile {
	id: string;
	email: string;
	fullname: string;
	balance: 0;
	avatar: string;
	joinedAt: string;
}

export interface ProfileState {
	isLoading: boolean;
	isUpdating: boolean;
	error: ErrorResponse | null;
	profile: Profile | null;
}

// EVENTS TYPES
export interface EventQueryParams {
	search?: string;
	status?: string;
	startDate?: string;
	endDate?: string;
	location?: string;
	sort?: string;
	page?: number;
	limit?: number;
}

export interface EventResponse {
	id: string;
	title: string;
	image: string;
	description: string;
	location: string;
	isAvailable: boolean;
	startPrice: number;
	startTime: number;
	endTime: number;
	date: string; // ISO string format
	status: string;
	createdAt: string; // ISO string format
}

export interface EventDetail {
	id: string;
	title: string;
	image: string;
	description: string;
	location: string;
	isAvailable: boolean;
	startPrice: number;
	startTime: number;
	endTime: number;
	date: string; // ISO string format
	status: string;
	tickets: Tiket[];
	createdAt: string; // ISO string format
}

export interface Tiket {
	id: string;
	name: string;
	price: number;
	sold: number;
	quota: number;
	limit?: number | null;
	isRefundable: boolean;
	refundPercent?: number;
	createdAt: string; // ISO string format
}

export interface UpdateEventRequest {
	title: string;
	description: string;
	location: string;
	date: string;
	startTime: number;
	endTime: number;
	status: 'active' | 'ongoing' | 'done' | 'cancelled';
	image?: File;
	imageURL?: string;
}

export interface CreateEventRequest {
	title: string;
	description: string;
	location: string;
	date: string;
	startTime: number;
	endTime: number;
	status?: 'active' | 'ongoing' | 'done' | 'cancelled';
	tickets: CreateTicketRequest[];
	image: File; // required saat create
	imageURL?: string;
}

export interface CreateTicketRequest {
	name: string;
	price: number;
	quota: number;
	limit?: number;
	refundable: boolean;
	refundPercent?: number;
}

// ORDERS RESPONSE TYPES

export interface CreateOrderRequest {
	eventId: string;
	orderDetails: {
		ticketId: string;
		quantity: number;
	}[];
	fullname: string;
	email: string;
	phone: string;
}

export interface Order {
	id: string;
	eventId: string;
	eventName: string;
	eventImage: string;
	fullname: string;
	email: string;
	phone: string;
	paymentUrl?: string;
	totalPrice: number;
	status: string;
	createdAt: string; // ISO
}

export interface OrderDetail {
	reduce(arg0: (sum: any, item: any) => any, arg1: number): any;
	id: string;
	ticketId: string;
	ticketName: string;
	quantity: number;
	price: number;
	createdAt: string; // ISO
}

export interface CheckoutSessionResponse {
	paymentId: string;
	sessionId: string;
	url: string;
}

export interface UserTicket {
	id: string;
	ticketName: string;
	qrCode: string;
	price: number;
	eventName: string;
	quantity: number;
	isUsed: boolean;
	isPrinted: boolean;
}

export interface RefundRequest {
	reason: string;
}

export interface OrderQueryParams {
	search?: string;
	status?: string;
	page?: number;
	limit?: number;
	sort?: string;
}
