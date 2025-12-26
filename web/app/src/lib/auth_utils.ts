/**
 * Logout the user by deleting the token present in cookie and redirects him the required page
 * pass "/" for user and "/admin" for admin generally
 * @param redirect - the relative pass of the page where user should be redirected after logout
 */
export function resetAuthTokenAndLogout(redirect: string): void {
	window.location.href = `/api/logout?redirect=${encodeURIComponent(redirect)}`;
}
