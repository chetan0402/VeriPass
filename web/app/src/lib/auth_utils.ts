export function resetAuthToken(redirect: string): void {
	window.location.href = `/api/logout?redirect=${encodeURIComponent(redirect)}`;
}
