// src/lib/stores/user.store.ts

import type {
	User,
	Profile,
	ProfileState,
	UpdateProfileRequest,
	ChangePasswordRequest
} from '$lib/types/api';
import { toast } from 'svelte-sonner';
import { authStore } from './auth.store';
import { writable, derived } from 'svelte/store';
import * as user from '$lib/services/user.service';
import { createStoreActions } from '$lib/utils/store';

const initialState: ProfileState = {
	error: null,
	profile: null,
	isLoading: false,
	isUpdating: false
};

function createUserStore() {
	const { subscribe, set, update } = writable<ProfileState>(initialState);
	const actions = createStoreActions<User>('user', 'user');

	return {
		subscribe,

		async changePassword(data: ChangePasswordRequest) {
			actions.setUpdating(update, true);

			try {
				const response = await user.changePassword(data);
				actions.setUpdating(update, false);
				toast.success(response.message || 'Password changed successfully');
				return response;
			} catch (error: any) {
				actions.setUpdating(update, false);
				return error;
			}
		},

		async getMyProfile() {
			actions.setLoading(update, true);

			try {
				const response: any = await user.getMyProfile();
				const userData = response.data;
				this.setProfile(userData);
				return response;
			} catch (error: any) {
				this.setError(error, 'Failed to fetch profile');
				throw error;
			}
		},

		async updateProfile(profileData: UpdateProfileRequest) {
			actions.setUpdating(update, true);

			try {
				const response: any = await user.updateProfile(profileData);
				const updatedUser = response.data;

				this.setProfile(updatedUser);
				authStore.setUser(updatedUser);
				toast.success(response.message || 'Profile updated successfully');
			} catch (error: any) {
				this.setError(error, 'Failed to update profile');
				throw error;
			}
		},

		setProfile: (profile: Profile | null) => {
			update((state) => ({
				...state,
				profile,
				error: null,
				isLoading: false,
				isUpdating: false
			}));
		},

		setError: (error: any, fallback: string) => {
			const message = error.response?.data?.message || fallback;
			actions.setError(update, error);
			console.error(message);
		},

		clearError: () => {
			actions.clearError(update);
		},

		reset: () => {
			actions.reset(update, initialState);
		}
	};
}

export const userStore = createUserStore();
export const userError = derived(userStore, ($userStore) => $userStore.error);
export const userProfile = derived(userStore, ($userStore) => $userStore.profile);
export const isLoading = derived(userStore, ($userStore) => $userStore.isLoading);
export const isUpdating = derived(userStore, ($userStore) => $userStore.isUpdating);
