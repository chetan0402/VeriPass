import { createClient } from '@connectrpc/connect';
import { type User, UserService } from '$lib/gen/veripass/v1/user_pb';
import { transport } from '$lib/transport';
import { NoUserSessionFound } from '$lib/errors';

let user: User | undefined;
const userClient = createClient(UserService, transport);

let userprofile: string | undefined;

export async function getUserFromState(): Promise<User> {
	if (user) {
		return user;
	}

	const userId = getSavedUserID();
	if (!userId) {
		throw new NoUserSessionFound();
	}

	try {
		user = await userClient.getUser({ id: userId });
		return user;
	} catch (e) {
		console.error('Error fetching user from server', e);
		throw e;
	}
}

export async function getUserProfileFromState() {
	if (userprofile) {
		return userprofile;
	}
	const userId = getSavedUserID();
	if (!userId) {
		throw new NoUserSessionFound();
	}
	try {
		const userProfileResponse = await userClient.getPhoto({ id: userId });
		const blob = new Blob([userProfileResponse.photo as BlobPart]);
		userprofile = URL.createObjectURL(blob);
		return userprofile;
	} catch (e) {
		console.error('Error fetching user image from server', e);
		throw e;
	}
}

export function invalidateUserSession() {
	//Reset all the session info
	localStorage.removeItem('user_id');
	user = undefined;
	return;
}

function getSavedUserID() {
	//To retrieve the saved user id after login; below is only for testing
	return localStorage.getItem('user_id');
}
