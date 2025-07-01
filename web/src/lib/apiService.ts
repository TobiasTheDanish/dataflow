import { baseUrl } from './env';
import type {
	CreateSiteData,
	ServiceResponse,
	Site,
	SiteConfig,
	SiteConnectionType,
	SiteDTO,
	TestConnectionData
} from './types';

type FetchFn = (input: RequestInfo | URL, init?: RequestInit) => Promise<Response>;

export type ApiService = ReturnType<typeof ApiServiceFactory>;

export const ApiServiceFactory = (fetch: FetchFn) => ({
	allSites: async function(): Promise<ServiceResponse<Site[]>> {
		const res = await fetch(`${baseUrl}/sites`);

		if (res.ok) {
			return {
				ok: true,
				data: await res
					.json()
					.then((json) => json.sites as SiteDTO[])
					.then((sites) =>
						sites.map((site) => ({
							...site,
							config: JSON.parse(site.config) as SiteConfig
						}))
					)
			};
		}

		return {
			ok: false,
			status: res.status,
			error: await res.text()
		};
	},
	testSiteConnection: async function(
		type: SiteConnectionType,
		data: TestConnectionData
	): Promise<ServiceResponse> {
		const res = await fetch(`${baseUrl}/sites/${type}/connect`, {
			method: 'post',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({
				...data.config,
				url: data.url
			})
		});

		if (res.ok) {
			return {
				ok: true,
				data: await res.json()
			};
		}

		return {
			ok: false,
			status: res.status,
			error: await res.text()
		};
	},
	createSite: async function(
		type: SiteConnectionType,
		data: CreateSiteData
	): Promise<ServiceResponse> {
		const res = await fetch(`${baseUrl}/sites/${type}`, {
			method: 'post',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({
				name: data.name,
				config: {
					...data.config,
					url: data.url
				}
			})
		});

		if (res.ok) {
			return {
				ok: true,
				data: await res.json()
			};
		}

		return {
			ok: false,
			status: res.status,
			error: await res.text()
		};
	}
});
