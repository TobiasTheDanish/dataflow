<script lang="ts">
	import SuperDebug, { superForm } from 'sveltekit-superforms';
	import type { PageProps } from './$types';
	import { Label } from '$lib/components/ui/label';
	import { Select, SelectContent, SelectItem, SelectTrigger } from '$lib/components/ui/select';
	import { Input } from '$lib/components/ui/input';
	import { Button } from '$lib/components/ui/button';
	import { Separator } from '$lib/components/ui/separator';
	import CreateSiteDialog from '$lib/components/create-site-dialog.svelte';

	const { data }: PageProps = $props();
	const { form, enhance } = superForm(data.form, {
		dataType: 'json'
	});

	const sites = $derived(data.sites);
</script>

<form method="POST" class=" m-6 flex h-[90vh] gap-6 rounded-lg border p-6 shadow-sm" use:enhance>
	<div class="flex w-[350px] flex-col gap-4">
		<p class="text-xl font-bold">Site information</p>

		<div class="flex gap-2">
			<!-- INSERT SITE SELECT HERE -->
			<Select
				type="single"
				value={$form.siteId.toString()}
				onValueChange={(v) => {
					if (sites.some((site) => site.id == parseInt(v))) {
						$form.siteId = parseInt(v);
					}
				}}
			>
				<SelectTrigger>
					{sites.find((site) => site.id == $form.siteId)?.name ?? 'Select site'}
				</SelectTrigger>
				<SelectContent>
					{#each sites as site (site.id)}
						<SelectItem value={site.id.toString()}>{site.name}</SelectItem>
					{/each}
				</SelectContent>
			</Select>
			<CreateSiteDialog />
		</div>
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
