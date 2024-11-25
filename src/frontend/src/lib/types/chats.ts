interface ChatChannelInfo {
	id: number;
	fname: string;
	lname: string;
	message: string;
	timesent: Date;
}

interface ChatMessage {
	sender: string;
	message: string;
	timesent: Date;
}

export type { ChatChannelInfo, ChatMessage };
