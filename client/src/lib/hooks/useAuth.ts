import { get } from 'svelte/store';
import { toast } from 'svelte-sonner';
import { goto } from '$app/navigation';
import { useFormState } from './useFormState';
import { authMe, login, logout } from '$lib/api/auth';
import { loginForm, userSession } from '$lib/stores/auth.store';

export async function useLogin() {
	const $state = get(useFormState(loginForm));
	const { form, isValid } = $state;

	if (!isValid) return;

	loginForm.setLoading(true);

	try {
		const res = await login(form);
		await authMe();
		toast.success(res.message);

		goto('/');
	} catch (err: any) {
		toast.error(err?.response?.data?.message || err.message || 'Login failed');
	} finally {
		loginForm.setLoading(false);
	}
}

export async function useLogout() {
	try {
		await logout();
		userSession.set(null);
		toast.success('Logout successful');
		goto('/');
	} catch (error) {
		console.error('Logout error:', error);
	}
}
