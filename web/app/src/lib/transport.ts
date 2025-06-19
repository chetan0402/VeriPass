import { createRouterTransport } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { UserService } from './gen/veripass/v1/user_pb';
import {
	PassService,
	Pass_PassType,
	type Pass,
	type ListPassesByUserRequest
} from '$lib/gen/veripass/v1/pass_pb';
import { msToTimestamp, timestampToMs } from '$lib/timestamp_utils';

const MOCK = true;

function generateMockPasesForPage(req: ListPassesByUserRequest): Pass[] {
	const mockPasses: Pass[] = [];
	for (let i = 0; i < req.pageSize; i++) {
		let endtime;
		if (i % 3 == 0) endtime = msToTimestamp(timestampToMs(req.pageToken) - i * 92880000);
		mockPasses.push({
			id: 'pass' + i + timestampToMs(req.pageToken),
			userId: req.userId,
			type: Pass_PassType.CLASS,
			startTime: msToTimestamp(timestampToMs(req.pageToken) - i * 828800000),
			endTime: endtime,
			$typeName: 'veripass.v1.Pass'
		});
	}
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
		let pageToken = msToTimestamp(0);
		const pageTokenMs = timestampToMs(req.pageToken);
		if (pageTokenMs > new Date().getTime() - 500000000) {
			pageToken = msToTimestamp(pageTokenMs - 100000000);
		}

		return {
			passes: generateMockPasesForPage(req),
			nextPageToken: pageToken
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
