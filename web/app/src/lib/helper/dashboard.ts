import { Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';

export const purposeOptions: { value: Pass_PassType; name: string }[] = [
	{ value: Pass_PassType.UNSPECIFIED, name: 'All' },
	{ value: Pass_PassType.CLASS, name: 'Class' },
	{ value: Pass_PassType.MARKET, name: 'Market' },
	{ value: Pass_PassType.HOME, name: 'Home' },
	{ value: Pass_PassType.EVENT, name: 'Event' }
];

export enum PassStatus {
	All = 2,
	Open = 0,
	Closed = 1
}

export const statusOptions: { value: PassStatus; name: string }[] = [
	{ value: PassStatus.All, name: 'All' },
	{ value: PassStatus.Open, name: 'Open' },
	{ value: PassStatus.Closed, name: 'Closed' }
];
