import { createClient } from '@connectrpc/connect';
import { type User, UserService } from '$lib/gen/veripass/v1/user_pb';
import { transport } from '$lib/transport';

let user: User | undefined;
const userClient = createClient(UserService, transport);

let userprofile: string | undefined;
/**
 * Returns the currently cached user details.
 *
 * If the user is already present in state, it is returned immediately.
 * Otherwise, the user details are fetched from the backend and cached
 * for future calls.
 * @returns A promise that resolves to the user details.
 * @throws {NotFound} If user details are not found in database
 * @throws {InvalidArgument} If token is not present
 */
export async function getUserFromState(): Promise<User> {
	if (user) {
		return user;
	}
	user = await userClient.getUser({});
	return user;
}

/**
 * Returns the currently caches user profile photo
 *
 * If the photo is already present in state, it is returned immediately.
 * Otherwise, the profile photo is fetched from the backend and cached
 * for future calls.
 * @returns A promise that resolves to user profile url;
 * @throws {NotFound} If user profile not found in database
 * @throws {InvalidArgument} If token is not present
 */
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

/**
 * Clears the cached user session details
 */
export function invalidateUserSession() {
	user = undefined;
	return;
}
