import { get } from 'svelte/store';
import { goto } from '$app/navigation';
import { AuthAPI } from '$lib/api/auth';
import { loginForm, authStore } from '$lib/stores/auth.store';

export const useLogin = async (): Promise<void> => {
	authStore.clearError();
	authStore.setLoading(true);
	try {
		const formData = get(loginForm);
		const response = await AuthAPI.login(formData);
		console.log('Login response:', response);
		authStore.setUser(response.data?.user);
		authStore.resetForms();
		await goto('/');
	} catch (err: any) {
		authStore.setError(err.message || 'Login failed. Please try again.');
	} finally {
		authStore.setLoading(false);
	}
};
