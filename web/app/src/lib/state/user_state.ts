import { createClient } from '@connectrpc/connect';
import { type User, UserService } from '$lib/gen/veripass/v1/user_pb';
import { transport } from '$lib/transport';

let user: User | undefined;
const userClient = createClient(UserService, transport);

let userprofile: string | undefined;

export async function getUserFromState(): Promise<User> {
	if (user) {
		return user;
	}
	user = await userClient.getUser({});
	return user;
}

export async function getUserProfileFromState() {
	if (userprofile) {
		return userprofile;
	}
	try {
		const userProfileResponse = await userClient.getPhoto({});
		const blob = new Blob([userProfileResponse.photo as BlobPart]);
		userprofile = URL.createObjectURL(blob);
		return userprofile;
	} catch (e) {
		console.error('Error fetching user image from server', e);
		throw e;
	}
}

export function invalidateUserSession() {
	user = undefined;
	return;
}
