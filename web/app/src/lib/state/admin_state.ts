import { createClient } from '@connectrpc/connect';
import { transport } from '$lib/transport';
import { NoAdminSessionFound } from '$lib/errors';
import { AdminService, type Admin } from '$lib/gen/veripass/v1/admin_pb';

let admin: Admin | undefined;
const adminClient = createClient(AdminService, transport);

export async function getAdminFromState(): Promise<Admin> {
	if (admin) {
		return admin;
	}

	const adminEmail = getSavedAdminEmail();
	if (!adminEmail) {
		throw new NoAdminSessionFound();
	}

	try {
		admin = await adminClient.getAdmin({ email: adminEmail });
		return admin;
	} catch (e) {
		console.error('Error fetching admin details from server', e);
		throw e;
	}
}

function getSavedAdminEmail() {
	//To retrieve the saved user id after login; below is only for testing
	return localStorage.getItem('admin_email');
}

export function invalidateAdminSession() {
	//Reset all the session info
	localStorage.removeItem('admin_email');
	admin = undefined;
	return;
}
