export function resetAuthToken(): void {
	console.log('Reset Auth Token');
	document.cookie = 'token=; Path=/; Max-Age=0; Secure; SameSite=Strict';
}
