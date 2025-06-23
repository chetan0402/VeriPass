export class NoUserSessionFound extends Error {
	constructor(message = 'No user session found') {
		super(message);
		this.name = 'NoUserSessionFound';
	}
}
