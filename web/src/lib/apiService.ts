import { baseUrl } from './env';
import type { CreateSiteData, SiteConnectionType, TestConnectionData } from './types';

type FetchFn = (input: RequestInfo | URL, init?: RequestInit) => Promise<Response>;

export type ApiService = ReturnType<typeof ApiServiceFactory>;

export const ApiServiceFactory = (fetch: FetchFn) => ({
	allSites: async function() {
		const res = await fetch(`${baseUrl}/sites`);

		if (res.ok) {
			return {
				ok: true,
				data: (await res.json()).sites
			};
		}

		return {
			ok: false,
			status: res.status,
			error: await res.text()
		};
	},
	testSiteConnection: async function(type: SiteConnectionType, data: TestConnectionData) {
		const res = await fetch(`http://localhost:8080/${type}/connect`, {
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
	createSite: async function(type: SiteConnectionType, data: CreateSiteData) {
		const res = await fetch(`http://localhost:8080/${type}/sites`, {
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
