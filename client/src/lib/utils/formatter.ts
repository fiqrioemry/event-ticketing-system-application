export function formatDate(dateString: any) {
	return new Date(dateString).toLocaleDateString('id-ID', {
		year: 'numeric',
		month: 'long',
		day: 'numeric'
	});
}

export function getAvatarInitials(fullname?: string) {
	if (!fullname) return '';
	return fullname
		.split(' ')
		.map((name) => name.charAt(0))
		.join('')
		.toUpperCase()
		.substring(0, 2);
}

export function getQRCodeURL(qrData: any, size = 200) {
	return `https://api.qrserver.com/v1/create-qr-code/?size=${size}x${size}&data=${encodeURIComponent(qrData)}`;
}

export function formatPrice(price: number): string {
	return new Intl.NumberFormat('id-ID', {
		style: 'currency',
		currency: 'IDR'
	}).format(price);
}

export function formatOptions(
	options: { id: string; name: string }[]
): { label: string; value: string }[] {
	return (
		options.map((option) => ({
			label: option?.name,
			value: option?.id
		})) || []
	);
}

export const buildFormData = (data: Record<string, any>): FormData => {
	const formData = new FormData();

	Object.entries(data).forEach(([key, value]) => {
		if (value === undefined || value === null) return;

		if (Array.isArray(value)) {
			if (value.length > 0 && value[0] instanceof File) {
				value.forEach((file) => {
					formData.append(key, file);
				});
			} else {
				value.forEach((item) => {
					formData.append(`${key}`, item);
				});
			}
		} else if (value instanceof File) {
			formData.append(key, value);
		} else {
			formData.append(key, value);
		}
	});

	return formData;
};
