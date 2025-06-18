import { type Timestamp, timestampFromDate, timestampMs } from '@bufbuild/protobuf/wkt';

export function timestampToMs(timestamp: Timestamp | undefined): number {
	if (!timestamp) return 0;
	return timestampMs(timestamp);
}

export function msToTimestamp(ms: number): Timestamp {
	const date = new Date(ms);
	return timestampFromDate(date);
}
