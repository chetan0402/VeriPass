import { createClient } from '@connectrpc/connect';
import { transport } from '$lib/transport';
import { AdminService, type Admin } from '$lib/gen/veripass/v1/admin_pb';

let admin: Admin | undefined;
const adminClient = createClient(AdminService, transport);

export async function getAdminFromState(): Promise<Admin> {
	if (admin) {
		return admin;
	}
	admin = await adminClient.getAdmin({});
	return admin;
}

export function invalidateAdminSession() {
	admin = undefined;
	return;
}
