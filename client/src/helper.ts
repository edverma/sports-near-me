import type {DbIssue} from "./models";

const formatNumber = (number: number): string => {
	const formattedNumber = number.toLocaleString('en-US', {
		minimumIntegerDigits: 2,
		useGrouping: false
	})
	return formattedNumber;
}

export const formatDateNumerical = (dateString: string) => {
	const date = new Date(dateString)
	const month = date.getUTCMonth() + 1; //months from 1-12
	const day = date.getUTCDate();
	const year = date.getUTCFullYear();

	return formatNumber(year) + "-" + formatNumber(month) + "-" + formatNumber(day);
}

export const formatDateReadable = (dateString: string) => {
	const date = new Date(dateString);
	const month = date.toLocaleString('default', { month: 'long', timeZone:'UTC' });
	const day = date.getUTCDate();
	const year = date.getUTCFullYear();

	return month + ' ' + day + ', ' + year
}

export const getDates = (issues: [DbIssue]) => issues.map(issue => new Date(issue.send_date))