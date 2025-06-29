import { z } from 'zod';
import type { Actions, PageServerLoad } from './$types';
import { fail, message, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';

const inputSchema = z
	.object({
		type: z.enum(['ftp', 'http']),
		address: z.string(),
		ftpConfig: z.object({
			port: z.coerce.number().default(21),
			authenticationRequired: z.coerce.boolean().default(false),
			username: z.string().optional(),
			password: z.string().optional()
		}),
		httpConfig: z.object({
			headers: z.record(z.string()).default({})
		}),
		fileInfo: z.object({
			path: z.string(),
			type: z.enum(['csv', 'json']).default('csv')
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
	const form = await superValidate(
		{
			type: 'ftp',
			ftpConfig: {
				authenticationRequired: false,
				port: 21
			}
		},
		zod(inputSchema)
	);

	return { form };
};

export const actions: Actions = {
	default: async ({ request, fetch }) => {
		const form = await superValidate(request, zod(inputSchema));
		if (!form.valid) {
			return fail(400, form);
		}

		console.log(form.data);

		const config = form.data.type == 'http' ? form.data.httpConfig : form.data.ftpConfig;

		const res = await fetch(`http://localhost:8080/${form.data.type}/connect`, {
			method: 'post',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({
				url: form.data.address,
				...config
			})
		});

		const response = await res.text();
		console.log(response);

		return message(form, response);
	}
};
