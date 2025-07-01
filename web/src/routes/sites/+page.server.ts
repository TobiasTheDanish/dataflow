import { z } from 'zod';
import { zod } from 'sveltekit-superforms/adapters';
import type { Actions, PageServerLoad } from './$types';
import { fail, message, superValidate } from 'sveltekit-superforms';

const siteSchema = z
	.object({
		type: z.enum(['ftp', 'http']),
		name: z.string(),
		address: z.string(),
		ftpConfig: z.object({
			port: z.coerce.number().default(21),
			authenticationRequired: z.coerce.boolean().default(false),
			username: z.string().optional(),
			password: z.string().optional()
		}),
		httpConfig: z.object({
			headers: z.record(z.string()).default({})
		})
	})
	.superRefine((val, ctx) => {
		if (val.type == 'http' && val.httpConfig == undefined) {
			ctx.addIssue({
				code: 'invalid_type',
				expected: 'object',
				received: 'undefined',
				message: 'When type is http, httpConfig cannot be missing'
			});
		}

		if (val.type == 'ftp' && val.ftpConfig == undefined) {
			ctx.addIssue({
				code: 'invalid_type',
				expected: 'object',
				received: 'undefined',
				message: 'When type is ftp, ftpConfig cannot be missing'
			});
		}

		if (val.type == 'ftp' && val.ftpConfig && val.ftpConfig.authenticationRequired) {
			if (val.ftpConfig.username == undefined) {
				ctx.addIssue({
					code: 'invalid_type',
					expected: 'string',
					received: 'undefined',
					message: 'When authentication is required, username cannot be missing'
				});
			}
			if (val.ftpConfig.password == undefined) {
				ctx.addIssue({
					code: 'invalid_type',
					expected: 'string',
					received: 'undefined',
					message: 'When authentication is required, password cannot be missing'
				});
			}
		}
	});

export const load: PageServerLoad = async () => {
	const form = await superValidate(zod(siteSchema));

	return { form };
};

export const actions: Actions = {
	testConnection: async ({ request, locals }) => {
		const form = await superValidate(request, zod(siteSchema));
		if (!form.valid) {
			return fail(400, { form });
		}

		console.log(form.data);

		const config = form.data.type == 'http' ? form.data.httpConfig : form.data.ftpConfig;

		const res = await locals.apiService.testSiteConnection(form.data.type, {
			url: form.data.address,
			config
		});

		if (res.ok) {
			return message(form, res.data);
		}

		return message(form, { status: res.status, error: JSON.parse(res.error) });
	},
	create: async ({ request, locals }) => {
		const form = await superValidate(request, zod(siteSchema));
		if (!form.valid) {
			return fail(400, { form });
		}

		const config = form.data.type == 'http' ? form.data.httpConfig : form.data.ftpConfig;

		const res = await locals.apiService.createSite(form.data.type, {
			name: form.data.name,
			url: form.data.address,
			config
		});

		if (!res.ok) {
			return fail(res.status ?? 500, { form });
		}

		return message(form, 'Site created');
	}
};
