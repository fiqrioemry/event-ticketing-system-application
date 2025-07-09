import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';
import { jwtDecode } from 'jwt-decode';
import { goto } from '$app/navigation';
import type { Writable } from 'svelte/store';

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

export function createFilterUpdater<T extends Record<string, any>>(
	store: Writable<T>,
	key: keyof T,
	callback?: () => void
) {
	return (val: T[keyof T]) => {
		store.update((prev) => ({ ...prev, [key]: val }));
		if (callback) callback();
	};
}

export function cleanFilters(obj: Record<string, unknown>) {
	const cleaned: Record<string, unknown> = {};
	for (const [key, value] of Object.entries(obj)) {
		// Hanya masukkan kalau value tidak kosong dan bukan default
		if (
			value !== null &&
			value !== undefined &&
			value !== '' &&
			!(key === 'sort' && value === 'date_asc') // nilai default, tidak usah masuk URL
		) {
			cleaned[key] = value;
		}
	}
	return cleaned;
}

export function formatRupiah(amount: number): string {
	return new Intl.NumberFormat('id-ID', {
		style: 'currency',
		currency: 'IDR',
		minimumFractionDigits: 0
	}).format(amount);
}

export function formatDate(dateStr: string) {
	return new Date(dateStr).toLocaleString('id-ID', {
		day: 'numeric',
		month: 'long',
		year: 'numeric'
	});
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type WithoutChild<T> = T extends { child?: any } ? Omit<T, 'child'> : T;
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type WithoutChildren<T> = T extends { children?: any } ? Omit<T, 'children'> : T;
export type WithoutChildrenOrChild<T> = WithoutChildren<WithoutChild<T>>;
export type WithElementRef<T, U extends HTMLElement = HTMLElement> = T & { ref?: U | null };

type JwtPayload = {
	userId: string;
	role: string;
	exp: number;
	iat: number;
};

/**
 * Decode a JWT token and validate its expiration.
 * @param token - The JWT string
 * @returns Decoded payload
 * @throws Error if token is expired or invalid
 */
export function decodeJwt(token: string): JwtPayload {
	const payload = jwtDecode<JwtPayload>(token);

	if (!payload || !payload.exp) {
		throw new Error('Invalid JWT payload');
	}

	const now = Math.floor(Date.now() / 1000);
	if (payload.exp < now) {
		throw new Error('Token expired');
	}

	return payload;
}

// src/lib/utils/urlSearchSync.ts

export function updateURLFromStore(storeValues: Record<string, string | number>) {
	const query = new URLSearchParams();
	for (const [key, value] of Object.entries(storeValues)) {
		if (value !== '' && value !== 'all') {
			query.set(key, String(value));
		}
	}
	goto(`?${query.toString()}`);
}
