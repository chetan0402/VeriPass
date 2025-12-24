import { createClient } from '@connectrpc/connect';
import { transport } from '$lib/transport';
import { AdminService, type Admin } from '$lib/gen/veripass/v1/admin_pb';

let admin: Admin | undefined;
const adminClient = createClient(AdminService, transport);

/**
 * Returns the currently cached admin details.
 *
 * If the admin is already present in state, it is returned immediately.
 * Otherwise, the admin details are fetched from the backend and cached
 * for future calls.
 * @returns A promise that resolves to the admin details.
 * @throws {NotFound} If admin is not found in database
 * @throws {InvalidArgument} If token is not present
 */
export async function getAdminFromState(): Promise<Admin> {
	if (admin) {
		return admin;
	}
	admin = await adminClient.getAdmin({});
	return admin;
}

/**
 * Clears the cached admin session details
 */
export function invalidateAdminSession() {
	admin = undefined;
	return;
}
