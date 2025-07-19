// src/routes/+layout.server.ts (create this file)
export async function load() {
	console.log('🚀 ROOT LAYOUT SERVER EXECUTED');
	return {
		serverTime: new Date().toISOString()
	};
}

export const ssr = true;
