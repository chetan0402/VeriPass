import { createRouterTransport } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { UserService } from './gen/veripass/v1/user_pb';
import {
	PassService,
	Pass_PassType,
	type Pass,
	type ListPassesByUserRequest
} from '$lib/gen/veripass/v1/pass_pb';
import type { Timestamp } from '@bufbuild/protobuf/wkt';

const MOCK = true;

function generateMockPasesForPage(req: ListPassesByUserRequest): Pass[] {
	const mockPasses: Pass[] = [];
	for (let i = 0; i < req.pageSize; i++) {
		let endtime;
		if (i % 3 == 0) endtime = unixToTimestamp(1450008583 + i * 9288);
		mockPasses.push({
			id: 'pass' + i + req.pageToken,
			userId: req.userId,
			type: Pass_PassType.CLASS,
			startTime: unixToTimestamp(1450008583 + i * 8288),
			endTime: endtime,
			$typeName: 'veripass.v1.Pass'
		});
	}
	if (Number(req.pageToken) > 3) return [];
	return mockPasses;
}

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
			passes: generateMockPasesForPage(req),
			nextPageToken: String(Number(req.pageToken) + 1)
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
