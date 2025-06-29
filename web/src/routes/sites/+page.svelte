<script lang="ts">
	import CreateSiteForm from '$lib/components/create-site-form.svelte';
	import type { PageProps } from './$types';

	const { data }: PageProps = $props();
</script>

<main class="mx-auto flex h-svh w-[90svw] flex-col items-center py-6">
	<CreateSiteForm data={data.form} />
</main>
