export interface UserType {
	email: string;
	avatar: string;
	fullname: string;
}

export interface ProfileTypes {
	id: string;
	email: string;
	fullname: string;
	avatar: string;
	role: string;
	balance: string;
	joinedAt: string;
}

export interface EventTypes {
	id: string;
	title: string;
	image: string;
	description: string;
	location: string;
	isAvailable: boolean;
	startPrice: number;
	startTime: number;
	endTime: number;
	date: string;
	status: string;
	createdAt: string;
}

export interface EventDetailTypes {
	id: string;
	title: string;
	image: string;
	startPrice: number;
	description: string;
	location: string;
	status: string;
	startTime: number;
	date: string;
	endTime: number;
	createdAt: string;
	tickets: TicketTypes[];
}

export interface TicketTypes {
	id: string;
	eventId: string;
	name: string;
	price: number;
	quota: number;
	limit: number;
	sold: number;
	isRefundable: boolean;
}

export interface OrderTypes {
	id: string;
	eventId: string;
	eventName: string;
	eventImage: string;
	fullname: string;
	email: string;
	phone: string;
	totalPrice: number;
	status: string;
	createdAt: string;
}

export interface OrderDetailTypes {
	id: string;
	ticketName: string;
	ticketId: string;
	quantity: number;
	price: number;
	createdAt: string;
}

export interface UserTicketTypes {
	id: string;
	eventId: string;
	ticketId: string;
	ticketName: string;
	eventName: string;
	qrCode: string;
	isUsed: boolean;
	usedAt?: string;
}

export interface RefundOrderTypes {
	orderId: string;
	refundAmount: number;
	refundedAt: string;
	userBalance: number;
}

export interface WithdrawalTypes {
	id: string;
	userId: string;
	amount: number;
	status: string;
	reason: string;
	createdAt: string;
	reviewedBy?: string;
	approvedAt?: string;
}

export interface TransactionReportTypes {
	orderId: string;
	userName: string;
	userEmail: string;
	eventTitle: string;
	totalPaid: number;
	refundAmount?: number;
	netRevenue: number;
	status: string;
	method: string;
	paidAt?: string;
	refundedAt?: string;
}

export interface PaginationTypes {
	page: number;
	limit: number;
	totalRows: number;
	totalPages: number;
}

//  AUTH TYPES
export interface RegisterRequestTypes {
	email: string;
	password: string;
	fullname: string;
}

export interface LoginRequestTypes {
	email: string;
	password: string;
}

export interface SendOTPRequestTypes {
	email: string;
}

export interface VerifyOTPRequestTypes {
	email: string;
	otp: string;
}

// USER TYPES
export interface UpdateProfileRequestTypes {
	fullname: string;
	avatar?: File;
	avatarURL?: string;
}

export interface UserQueryParamsTypes {
	q?: string;
	role?: string;
	sort?: string;
	page?: number;
	limit?: number;
}

// EVENT TYPES

export interface CreateTicketRequestTypes {
	name: string;
	eventId: string;
	price: number;
	quota: number;
	limit?: number;
	isRefundable: boolean;
	refundPercent?: number;
}

export interface CreateEventRequestTypes {
	title: string;
	image: File;
	imageURL?: string;
	description: string;
	location: string;
	date: string;
	startTime: number;
	endTime: number;
	tickets: CreateTicketRequestTypes[];
	status: 'active' | 'ongoing' | 'done' | 'cancelled';
}

export interface UpdateEventRequestTypes {
	title: string;
	image: File;
	imageURL?: string;
	description: string;
	location: string;
	date: string;
	startTime: number;
	endTime: number;
	status: string;
	ticketCategories: TicketCategoryDetailTypes[];
}

export interface TicketCategoryDetailTypes {
	id: string;
	name: string;
	price: number;
	quota: number;
	isRefundable: boolean;
	refundAmount: number;
}

export interface EventQueryParamsTypes {
	q?: string;
	status?: string;
	startDate?: string;
	endDate?: string;
	location?: string;
	sort?: string;
	page?: number;
	limit?: number;
}

// TICKET TYPES
export interface UpdateTicketRequestTypes {
	name: string;
	eventId: string;
	price: number;
	quota: number;
	limit?: number;
	isRefundable: boolean;
}

// ORDER TYPES
export interface CreateOrderRequestTypes {
	eventId: string;
	orderDetails: OrderDetailRequestTypes[];
	fullname: string;
	email: string;
	phone: string;
}

export interface OrderDetailRequestTypes {
	ticketId: string;
	quantity: number;
}

export interface OrderQueryParamsTypes {
	q?: string;
	status?: string;
	page?: number;
	limit?: number;
	sort?: string;
}

export interface ValidateTicketRequestTypes {
	qrCode: string;
}

export interface RefundOrderRequestTypes {
	reason: string;
}

export interface CreateWithdrawalRequestTypes {
	amount: number;
	reason: string;
}

export interface TransactionReportQueryTypes {
	start_date?: string;
	end_date?: string;
	export?: boolean;
}
