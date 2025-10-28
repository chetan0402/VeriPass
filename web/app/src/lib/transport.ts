import { Code, ConnectError, createRouterTransport } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { ExitRequest_ExitType, UserService } from './gen/veripass/v1/user_pb';
import { PassService, Pass_PassType, type Pass } from '$lib/gen/veripass/v1/pass_pb';
import { msToTimestamp, timestampToMs } from '$lib/timestamp_utils';
import { timestampNow } from '@bufbuild/protobuf/wkt';
import {
	type Admin,
	AdminService,
	type GetAllPassesByHostelRequest,
	type GetAllPassesByHostelResponse_InfoIncludedPass
} from '$lib/gen/veripass/v1/admin_pb';

const MOCK = true;

const mockPasses: {
	[id: string]: Pass;
} = {};

function generateMockPasesForPage() {
	const newMockPasses: Pass[] = [];
	for (let i = 0; i < 30; i++) {
		let endtime;
		const idIdentifier = timestampToMs(timestampNow()) - i * 60 * 60 * 1000;
		if (idIdentifier % 3 == 0) endtime = msToTimestamp(idIdentifier - 4 * 60 * 60 * 1000);
		const id = 'pass' + idIdentifier;
		mockPasses[id] = {
			id: id,
			userId: '12345',
			type: Pass_PassType.CLASS,
			startTime: msToTimestamp(idIdentifier - 60 * 60 * 1000),
			endTime: endtime,
			$typeName: 'veripass.v1.Pass',
			qrCode: 'https://www.google.com'
		};
		newMockPasses.push(mockPasses[id]);
	}
	return newMockPasses;
}

function getPassType(selected: ExitRequest_ExitType): Pass_PassType {
	const map: Record<ExitRequest_ExitType, Pass_PassType> = {
		[ExitRequest_ExitType.CLASS]: Pass_PassType.CLASS,
		[ExitRequest_ExitType.MARKET]: Pass_PassType.MARKET,
		[ExitRequest_ExitType.HOME]: Pass_PassType.HOME,
		[ExitRequest_ExitType.EVENT]: Pass_PassType.EVENT,
		[ExitRequest_ExitType.UNSPECIFIED]: Pass_PassType.UNSPECIFIED
	};
	return map[selected] ?? ExitRequest_ExitType.UNSPECIFIED;
}

function generateMockPasesForHostel(
	req: GetAllPassesByHostelRequest
): GetAllPassesByHostelResponse_InfoIncludedPass[] {
	const newMockPasses: GetAllPassesByHostelResponse_InfoIncludedPass[] = [];
	for (let i = 0; i < req.pageSize; i++) {
		let endtime;
		let mockStartTime = req.startTime;
		if (timestampToMs(timestampNow()) - timestampToMs(req.pageToken) > 60 * 60 * 1000) {
			mockStartTime = req.pageToken;
		}
		const idIdentifier = timestampToMs(mockStartTime) + (i + 1) * 60 * 1000;
		const id = 'pass' + idIdentifier;
		if (!req.passIsOpen) {
			endtime = msToTimestamp(idIdentifier + Math.random() * 60 * 60 * 1000);
		}
		let passType = Math.floor(Math.random() * 3) + 1;
		if (req.type !== Pass_PassType.UNSPECIFIED) {
			passType = req.type;
		}
		const infoPass: GetAllPassesByHostelResponse_InfoIncludedPass = {
			$typeName: 'veripass.v1.GetAllPassesByHostelResponse.InfoIncludedPass',
			pass: {
				id: id,
				userId: '12345',
				type: passType,
				startTime: msToTimestamp(idIdentifier),
				endTime: endtime,
				$typeName: 'veripass.v1.Pass',
				qrCode: 'https://www.google.com'
			},
			studentRoom: 'C' + Math.floor(Math.random() * 100),
			studentName: 'Mock Student' + Math.floor(Math.random() * 100)
		};
		newMockPasses.push(infoPass);
	}
	return newMockPasses;
}

const mockRouter = createRouterTransport(({ rpc }) => {
	rpc(UserService.method.getUser, (req) => {
		if (req.id === '12345') {
			return {
				id: req.id,
				name: 'Mock User',
				hostel: 'Mock Hostel',
				room: 'Mock Room',
				phone: '1234567890'
			};
		} else {
			throw new ConnectError('User not found', Code.NotFound);
		}
	});

	rpc(AdminService.method.getAdmin, (req) => {
		return {
			email: req.email,
			name: 'Mock Admin',
			hostel: 'Mock Hostel',
			canAddPass: true,
			$typeName: 'veripass.v1.Admin'
		} satisfies Admin;
	});

	rpc(PassService.method.getLatestPassByUser, (req) => {
		const sortedPasses = Object.values(mockPasses)
			.filter((p) => p.userId === req.userId)
			.sort((a, b) => timestampToMs(b.startTime) - timestampToMs(a.startTime));
		if (!sortedPasses[0]) {
			throw new ConnectError('Pass not found', Code.NotFound);
		}
		return sortedPasses[0];
	});

	rpc(UserService.method.exit, (req) => {
		console.log('exit', req.id);

		const userId = String(req.id);
		const idIdentifier = userId + timestampToMs(timestampNow());
		const id = 'pass' + idIdentifier;

		mockPasses[id] = {
			id: id,
			userId: userId,
			type: getPassType(req.type),
			startTime: timestampNow(),
			$typeName: 'veripass.v1.Pass',
			qrCode: 'https://www.google.com'
		};

		return {
			passId: id
		};
	});

	rpc(UserService.method.entry, (req) => {
		const mockPass = mockPasses[req.passId];
		if (!mockPass) {
			throw new ConnectError('Pass not found', Code.NotFound);
		}
		mockPass.endTime = timestampNow();
		mockPasses[req.passId] = mockPass;
		return {};
	});

	rpc(PassService.method.getPass, (req) => {
		const pass = mockPasses[req.id];
		if (!pass) {
			throw new ConnectError('Pass not found', Code.NotFound);
		}
		return pass;
	});

	rpc(PassService.method.listPassesByUser, (req) => {
		const pageSize = req.pageSize;
		const pageTokenMs = timestampToMs(req.pageToken);
		if (Object.values(mockPasses).length < 10) {
			generateMockPasesForPage();
		}
		const sortedPasses = Object.values(mockPasses)
			.filter((p) => p.userId === req.userId)
			.sort((a, b) => timestampToMs(b.startTime) - timestampToMs(a.startTime));
		const paginated = sortedPasses.filter((p) => timestampToMs(p.startTime) < pageTokenMs);
		const passes = paginated.slice(0, pageSize);
		const nextPageToken =
			paginated.length > pageSize ? passes[passes.length - 1].startTime : msToTimestamp(0);
		return {
			passes,
			nextPageToken
		};
	});

	rpc(AdminService.method.getAllPassesByHostel, (req) => {
		const infoIncludedPasses: GetAllPassesByHostelResponse_InfoIncludedPass[] =
			generateMockPasesForHostel(req);
		console.log(infoIncludedPasses);
		return {
			passes: infoIncludedPasses,
			nextPageToken: infoIncludedPasses[infoIncludedPasses.length - 1].pass?.startTime
		};
	});

	rpc(UserService.method.getPhoto, async () => {
		const response = await fetch(
			'https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/User_icon_2.svg/1200px-User_icon_2.svg.png'
		);
		const arrayBuffer = await response.arrayBuffer();
		return {
			photo: new Uint8Array(arrayBuffer),
			$typeName: 'veripass.v1.GetPhotoResponse'
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
