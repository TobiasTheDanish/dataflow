export type SiteConnectionType = 'ftp' | 'http';

export type ErrorServiceResponse<E extends string = string> = {
	ok: false;
	status: number;
	error: E;
};

export type SuccesServiceResponse<T = unknown> = {
	ok: true;
	data: T;
};

export type ServiceResponse<T extends unknown = unknown, E extends string = string> =
	| SuccesServiceResponse<T>
	| ErrorServiceResponse<E>;

export interface SiteDTO {
	id: number;
	name: string;
	type: SiteConnectionType;
	config: string;
}

export interface Site {
	id: number;
	name: string;
	type: SiteConnectionType;
	config: SiteConfig;
}

export type HttpConfig = {
	url: string;
	headers: Record<string, string>;
};
export type FtpConfig = {
	url: string;
	port: number;
	authenticationRequired: boolean;
	username: string;
	password: string;
};

export type SiteConfig = HttpConfig | FtpConfig;

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
