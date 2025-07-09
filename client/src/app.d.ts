import type { UserResponse } from '$lib/types/auth';

// src/app.d.ts
declare global {
	namespace App {
		interface Locals {
			user: UserResponse | null;
			safeGetSession?: () => Promise<{
				user: UserResponse | null;
			}>;
		}
	}
}

export {};
