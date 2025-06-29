<script lang="ts">
	import SuperDebug, { superForm } from 'sveltekit-superforms';
	import type { PageProps } from './$types';
	import { Label } from '$lib/components/ui/label';
	import { Select, SelectContent, SelectItem, SelectTrigger } from '$lib/components/ui/select';
	import { Input } from '$lib/components/ui/input';
	import { Button } from '$lib/components/ui/button';
	import { Separator } from '$lib/components/ui/separator';
	import FtpConfig from './ftp-config.svelte';
	import HttpConfig from './http-config.svelte';

	const { data }: PageProps = $props();
	const { form, enhance } = superForm(data.form, {
		dataType: 'json'
	});
</script>

<form method="POST" class=" m-6 flex h-[90vh] gap-6 rounded-lg border p-6 shadow-sm" use:enhance>
	<div class="flex w-[350px] flex-col gap-4">
		<p class="text-xl font-bold">Site information</p>
		<div class="grid gap-2">
			<Label for="type">Type</Label>
			<Select type="single" name="type" bind:value={$form.type}>
				<SelectTrigger class="uppercase">{$form.type}</SelectTrigger>
				<SelectContent>
					<SelectItem value="ftp">FTP</SelectItem>
					<SelectItem value="http">HTTP</SelectItem>
				</SelectContent>
			</Select>
		</div>

		<div class="grid gap-2">
			<Label for="address">Address</Label>
			<Input name="address" bind:value={$form.address} type="text" inputmode="url" />
		</div>
		{#if $form.type == 'ftp'}
			<FtpConfig {form} />
		{:else if $form.type == 'http'}
			<HttpConfig {form} />
		{:else}
			<p>NOT IMPLEMENTED YET!!</p>
		{/if}
	</div>

	<Separator orientation="vertical" />

	<div class="flex w-[350px] flex-col gap-4">
		<p class="text-xl font-bold">File information</p>
		<div class="grid gap-2">
			<Label for="filepath">Filepath</Label>
			<Input name="filepath" bind:value={$form.fileInfo.path} type="text" />
		</div>
		<div class="grid gap-2">
			<Label for="filetype">Filetype</Label>
			<Select type="single" name="filetype" bind:value={$form.fileInfo.type}>
				<SelectTrigger class="uppercase">{$form.fileInfo.type}</SelectTrigger>
				<SelectContent>
					<SelectItem value="csv">CSV</SelectItem>
					<SelectItem value="json">JSON</SelectItem>
				</SelectContent>
			</Select>
		</div>
	</div>

	<Separator orientation="vertical" />

	<div class="ml-auto flex w-fit items-end justify-end">
		<Button type="submit">Submit</Button>
	</div>
</form>

<SuperDebug data={form} />
