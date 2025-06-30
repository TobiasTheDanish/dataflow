import type { ApiService } from '$lib/apiService';

declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			apiService: ApiService;
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export { };
