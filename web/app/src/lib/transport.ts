import { createRouterTransport } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { UserService } from './gen/veripass/v1/user_pb';

const MOCK = true;

const mockRouter = createRouterTransport(({ rpc }) => {
	rpc(UserService.method.getUser, (req) => {
		return {
			id: req.id,
			name: 'Mock User',
			hostel: 'Mock Hostel',
			room: 'Mock Room',
			phone: '1234567890'
		};
	});
});

export const transport = createConnectTransport({
	baseUrl: '/api',
	interceptors: [
		(next) => async (req) => {
			if (!req.stream && MOCK) {
				return await mockRouter.unary(
					req.method,
					req.signal,
					undefined,
					req.header,
					req.message,
					req.contextValues
				);
			}
			return await next(req);
		}
	]
});
