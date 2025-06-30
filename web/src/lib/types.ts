export type SiteConnectionType = 'ftp' | 'http';

export interface CreateSiteData {
	name: string;
	url: string;
	config:
	| {
		port: number;
		authenticationRequired: boolean;
		username?: string | undefined;
		password?: string | undefined;
	}
	| {
		headers: Record<string, string>;
	};
}

export interface TestConnectionData {
	url: string;
	config:
	| {
		port: number;
		authenticationRequired: boolean;
		username?: string | undefined;
		password?: string | undefined;
	}
	| {
		headers: Record<string, string>;
	};
}
