import { type Timestamp, timestampFromDate, timestampMs } from '@bufbuild/protobuf/wkt';

export function timestampToMs(timestamp: Timestamp | undefined): number {
	if (!timestamp) return 0;
	return timestampMs(timestamp);
}

export function msToTimestamp(ms: number): Timestamp {
	const date = new Date(ms);
	return timestampFromDate(date);
}

export function timestampToDate(startTime: Timestamp) {
	const startMillis = Number(startTime.seconds) * 1000 + Math.floor(startTime.nanos / 1e6);
	return new Date(startMillis);
}

export function formatTimeStringLocal(date: Date): string {
	const hours = date.getHours();
	const minutes = date.getMinutes();
	const hour12 = hours % 12 || 12;
	const minuteStr = minutes.toString().padStart(2, '0');
	const hoursStr = hour12.toString().padStart(2, '0');
	return `${hoursStr}:${minuteStr}`;
}

export function getFormattedTimeSuffixLocal(timeStamp?: Timestamp) {
	if (timeStamp) {
		const startDate = timestampToDate(timeStamp);
		return startDate.getHours() < 12 ? 'AM' : 'PM';
	}
	return '';
}

export function formatDateString(date: Date) {
	return date.toLocaleDateString('en-In', {
		day: '2-digit',
		month: 'short',
		year: 'numeric'
	});
}

export function msToDurationString(ms: number): string {
	const totalSeconds = Math.floor(ms / 1000);
	const days = Math.floor(totalSeconds / (3600 * 24));
	const hours = Math.floor((totalSeconds % (3600 * 24)) / 3600);
	const minutes = Math.floor((totalSeconds % 3600) / 60);
	const seconds = totalSeconds % 60;
	if (totalSeconds < 60) return `${seconds.toString()}s`;
	if (days > 0)
		return `${days.toString()} day${days > 1 ? 's' : ''} ${hours.toString()} hr${hours !== 1 ? 's' : ''}`;
	return `${hours.toString()} hr${hours !== 1 ? 's' : ''} ${minutes.toString()} min${minutes !== 1 ? 's' : ''}`;
}

export function get12oClockDate(date: Date): Date {
	date.setHours(0, 0, 0, 0);
	return date;
}

export function toISTDateStringFull(date: Date) {
	return date.toLocaleString('en-In', {
		hour12: false
	});
}
