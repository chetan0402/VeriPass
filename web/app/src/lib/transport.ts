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
import { timestampNow } from '@bufbuild/protobuf/wkt';

const MOCK = true;

function generateMockPasesForPage(req: ListPassesByUserRequest): Pass[] {
	const mockPasses: Pass[] = [];
	for (let i = 0; i < req.pageSize; i++) {
		let endtime;
		const idIdentifier = timestampToMs(req.pageToken) - i * 60 * 60 * 1000;
		if (idIdentifier % 3 == 0) endtime = msToTimestamp(idIdentifier - 4 * 60 * 60 * 1000);
		const id = 'pass' + idIdentifier;
		mockPasses.push({
			id: id,
			userId: req.userId,
			type: Pass_PassType.CLASS,
			startTime: msToTimestamp(idIdentifier - 60 * 60 * 1000),
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

	rpc(PassService.method.getLatestPassByUser, (req) => {
		return {
			id: 'pass' + req.userId,
			userId: req.userId,
			type: Pass_PassType.CLASS,
			startTime: msToTimestamp(timestampToMs(timestampNow()) - 4 * 60 * 60 * 1000),
			endTime: msToTimestamp(timestampToMs(timestampNow()) - 60 * 60 * 1000),
			$typeName: 'veripass.v1.Pass'
		};
	});
	rpc(PassService.method.getPass, (req) => {
		let endtime;
		const idIdentifier = Number(req.id.replace('pass', ''));
		if (idIdentifier % 3 == 0) endtime = msToTimestamp(idIdentifier - 4 * 60 * 60 * 1000);
		const id = 'pass' + idIdentifier;

		return {
			id: id,
			userId: '12345',
			type: Pass_PassType.CLASS,
			startTime: msToTimestamp(idIdentifier - 60 * 60 * 1000),
			endTime: endtime,
			$typeName: 'veripass.v1.Pass'
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
