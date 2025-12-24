import { formatDateString, formatTimeStringLocal, timestampToDate } from './time_utils';
import type { Timestamp } from '@bufbuild/protobuf/wkt';
import { type Pass, Pass_PassType } from '$lib/gen/veripass/v1/pass_pb';

/**
 * Converts a timestamp into a formatted local time string.
 * @param timeStamp - The timestamp object to be formatted, or undefined.
 * @returns The local time string if valid, otherwise a placeholder '----'.
 */
export function getFormattedTime(timeStamp?: Timestamp) {
	if (timeStamp) {
		const date = timestampToDate(timeStamp);
		return formatTimeStringLocal(date);
	}
	return '----';
}

/**
 * Converts a timestamp into a formatted date string.
 * @param timeStamp - The timestamp object to be formatted, or undefined.
 * @returns The date string if valid, otherwise a placeholder '----'.
 */
export function getFormattedDate(timeStamp?: Timestamp) {
	if (timeStamp) {
		const date = timestampToDate(timeStamp);
		return formatDateString(date);
	}
	return '----';
}

/**
 * @param passItem - The pass object containing the type definition, or undefined.
 * @returns The string representation of the pass type (e.g., 'Class', 'Home'), or 'Not specified' if undefined or unknown.
 */
export function getPassType(passItem?: Pass) {
	switch (passItem?.type) {
		case Pass_PassType.CLASS:
			return 'Class';
		case Pass_PassType.HOME:
			return 'Home';
		case Pass_PassType.EVENT:
			return 'Event';
		case Pass_PassType.MARKET:
			return 'Market';
		default:
			return 'Not specified';
	}
}
