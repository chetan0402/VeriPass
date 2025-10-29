import { type Timestamp, timestampFromDate, timestampMs } from '@bufbuild/protobuf/wkt';

export function timestampToMs(timestamp: Timestamp | undefined): number {
	if (!timestamp) return 0;
	return timestampMs(timestamp);
}

export function msToTimestamp(ms: number): Timestamp {
	const date = new Date(ms);
	return timestampFromDate(date);
}
export function msToDurationString(ms: number): string {
	const totalSeconds = Math.floor(ms / 1000);
	const days = Math.floor(totalSeconds / (3600 * 24));
	const hours = Math.floor((totalSeconds % (3600 * 24)) / 3600);
	const minutes = Math.floor((totalSeconds % 3600) / 60);
	const seconds = totalSeconds % 60;

	if (totalSeconds < 60) return `${seconds}s`;
	if (days > 0) return `${days} day${days > 1 ? 's' : ''} ${hours} hr${hours !== 1 ? 's' : ''}`;
	return `${hours} hr${hours !== 1 ? 's' : ''} ${minutes} min${minutes !== 1 ? 's' : ''}`;
}
export function get12OClockDate(date: Date): Date {
	date.setHours(0, 0, 0, 0);
	return date;
}
