import { createRouterTransport } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { UserService } from './gen/veripass/v1/user_pb';
import { PassService, Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';
import type { Timestamp } from '@bufbuild/protobuf/wkt';

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

	rpc(PassService.method.listPassesByUser, (req) => {
		return {
			passes: [
				{
					id: 'pass1',
					userId: req.userId,
					type: Pass_PassType.CLASS,
					startTime: unixToTimestamp(1450008583),
					$typeName: 'veripass.v1.Pass'
				},
				{
					id: 'pass4',
					userId: req.userId,
					type: Pass_PassType.MARKET,
					startTime: unixToTimestamp(1750008583),
					$typeName: 'veripass.v1.Pass'
				},
				{
					id: 'pass2',
					userId: req.userId,
					type: Pass_PassType.HOME,
					startTime: unixToTimestamp(1750002483),
					endTime: unixToTimestamp(1750008983),
					$typeName: 'veripass.v1.Pass'
				}
			],
			nextPageToken: ''
		};
	});
});

function unixToTimestamp(unixSeconds: number): Timestamp {
	const date = new Date(unixSeconds * 1000);
	const seconds = Math.floor(date.getTime() / 1000);
	const nanos = (date.getTime() % 1000) * 1_000_000; // Convert milliseconds to nanoseconds
	return { $typeName: 'google.protobuf.Timestamp', seconds: BigInt(seconds), nanos };
}

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
