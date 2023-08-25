export type DbIssue = {
	heading: string;
	subheadings: [string];
	content: [string];
	send_date: Date;
	username: string;
	created_at: Date;
	columns: number;
};

export type DbCredential = {
	username: string;
	hash: string;
}