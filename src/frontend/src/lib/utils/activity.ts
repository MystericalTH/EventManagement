import type { Activity } from '$lib/types';

export function formatActivityDateTime(activity: Activity): string {
	const formatDate = (date: string) => {
		const d = new Date(date);
		return `${d.getDate()}/${d.getMonth() + 1}/${d.getFullYear()}`;
	};
	const formatTime = (time: string) => time;

	const { startDate, startTime, endDate, endTime } = activity;

	if (endDate) {
		if (startDate === endDate) {
			if (startTime && endTime) {
				return `${formatDate(startDate)} ${formatTime(startTime)} - ${formatTime(endTime)}`;
			} else {
				return `${formatDate(startDate)}`;
			}
		} else {
			if (startTime && endTime) {
				return `${formatDate(startDate)} ${formatTime(startTime)} - ${formatDate(endDate)} ${formatTime(endTime)}`;
			} else {
				return `${formatDate(startDate)} - ${formatDate(endDate)}`;
			}
		}
	} else {
		if (startTime) {
			return `${formatDate(startDate)} ${formatTime(startTime)}`;
		} else {
			return formatDate(startDate);
		}
	}
}
