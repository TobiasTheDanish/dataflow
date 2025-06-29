<script lang="ts">
	import { Loader2 } from '@lucide/svelte';
	import { superForm } from 'sveltekit-superforms/client';
	import { Label } from './ui/label';
	import { Select, SelectContent, SelectItem, SelectTrigger } from './ui/select';
	import { Input } from './ui/input';
	import FtpConfig from './ftp-config.svelte';
	import HttpConfig from './http-config.svelte';
	import { Button } from './ui/button';

	const {
		data
	}: {
		data: {
			type: 'ftp' | 'http';
			name: string;
			ftpConfig: {
				port: number;
				authenticationRequired: boolean;
				username?: string | undefined;
				password?: string | undefined;
			};
			address: string;
			httpConfig: {
				headers: Record<string, string>;
			};
		};
	} = $props();

	const { form, enhance } = superForm(data, {
		dataType: 'json'
	});

	const {
		submit,
		submitting,
		enhance: testEnhance
	} = superForm(data, {
		dataType: 'json',
		invalidateAll: false,
		applyAction: false,
		onSubmit: ({ jsonData }) => {
			jsonData($form);
		}
	});
</script>

<form method="POST" action="/sites?/create" class="flex flex-col gap-4" use:enhance>
	<div class="grid gap-2">
		<Label for="type">Type</Label>
		<Select type="single" name="type" bind:value={$form.type}>
			<SelectTrigger id="type" class="uppercase">{$form.type}</SelectTrigger>
			<SelectContent>
				<SelectItem value="ftp">FTP</SelectItem>
				<SelectItem value="http">HTTP</SelectItem>
			</SelectContent>
		</Select>
	</div>

	<div class="grid gap-2">
		<Label for="name">Name</Label>
		<Input id="name" name="name" bind:value={$form.name} type="text" />
	</div>

	<div class="grid gap-2">
		<Label for="address">Address</Label>
		<Input id="address" name="address" bind:value={$form.address} type="text" inputmode="url" />
	</div>
	{#if $form.type == 'ftp'}
		<FtpConfig {form} />
	{:else if $form.type == 'http'}
		<HttpConfig {form} />
	{:else}
		<p>NOT IMPLEMENTED YET!!</p>
	{/if}

	<Button form="test-connection" type="button" onclick={submit} disabled={$submitting}>
		{#if $submitting}
			<Loader2 class="animate-spin" />
		{/if}
		Test connection
	</Button>

	<Button type="submit">Create site</Button>
</form>

<form id="test-connection" method="POST" action="/sites?/testConnection" use:testEnhance></form>
