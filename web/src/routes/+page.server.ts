import type { Actions, PageServerLoad } from './$types';
import { fail, message, superValidate } from 'sveltekit-superforms';
import { z } from 'zod';
import { zod } from 'sveltekit-superforms/adapters';
import { error } from '@sveltejs/kit';

const inputSchema = z.object({
	siteId: z.coerce.number(),
	fileInfo: z.object({
		path: z.string(),
		type: z.enum(['csv', 'json']).default('csv')
	})
});

export const load: PageServerLoad = async ({ locals }) => {
	const form = await superValidate(zod(inputSchema));

	const siteRes = await locals.apiService.allSites();
	if (!siteRes.ok) {
		console.log(siteRes.error);
		error(siteRes.status ?? 500, siteRes.error);
	}

	console.log(siteRes.data);

	return { form, sites: siteRes.data };
};

export const actions: Actions = {};
