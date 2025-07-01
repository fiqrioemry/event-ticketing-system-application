// src/lib/hooks/use-toast.ts
import { toast as t } from 'sonner';

export const toast = {
	loading: (msg: string, opts?: any) => t.loading(msg, opts),
	success: (msg: string, opts?: any) => t.success(msg, opts),
	error: (msg: string, opts?: any) => t.error(msg, opts)
};
