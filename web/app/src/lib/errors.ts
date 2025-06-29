export class NoUserSessionFound extends Error {
	constructor(message = 'No user session found') {
		super(message);
		this.name = 'NoUserSessionFound';
	}
}
export class NoAdminSessionFound extends Error {
	constructor(message = 'No admin session found') {
		super(message);
		this.name = 'NoAdminSessionFound';
	}
}
