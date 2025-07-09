export interface ResponseState {
	success: boolean;
	data?: any; // Use 'any' for flexibility, or define a more specific type if needed
	error?: string; // Optional error message
	message: string;
}
