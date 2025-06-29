<script lang="ts">
	import { Label } from '$lib/components/ui/label';
	import { Textarea } from '$lib/components/ui/textarea';
	import type { SuperFormData } from 'sveltekit-superforms/client';

	interface Props<
		T extends {
			httpConfig: {
				headers: Record<string, string>;
			};
		} = {
			httpConfig: {
				headers: Record<string, string>;
			};
		}
	> {
		form: SuperFormData<T>;
	}

	let text = $state('');

	const { form }: Props = $props();

	const oninput = () => {
		try {
			$form.httpConfig.headers = JSON.parse(text);
		} catch (e) {
			console.error(e);
		}
	};
</script>

<div class="grid gap-2">
	<Label for="headers">Headers</Label>
	<Textarea name="headers" bind:value={text} {oninput}></Textarea>
</div>
