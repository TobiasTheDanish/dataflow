<script lang="ts">
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import type { SuperFormData } from 'sveltekit-superforms/client';

	interface Props<
		T extends {
			ftpConfig: {
				port: number;
				authenticationRequired: boolean;
				username?: string | undefined;
				password?: string | undefined;
			};
		} = {
			ftpConfig: {
				port: number;
				authenticationRequired: boolean;
				username?: string | undefined;
				password?: string | undefined;
			};
		}
	> {
		form: SuperFormData<T>;
	}

	const { form }: Props = $props();
</script>

<div class="grid gap-2">
	<Label for="port">Port</Label>
	<Input name="port" bind:value={$form.ftpConfig.port} type="number" step="1" />
</div>

<div class="flex gap-2">
	<input
		type="checkbox"
		bind:checked={$form.ftpConfig.authenticationRequired}
		name="authenticationRequired"
	/>
	<label for="authenticationRequired">Require Auth</label>
</div>
<div class="grid gap-2">
	<Label for="username">Username</Label>
	<Input
		name="username"
		disabled={!$form.ftpConfig.authenticationRequired}
		bind:value={$form.ftpConfig.username}
		type="text"
	/>
</div>
<div class="grid gap-2">
	<Label for="password">Password</Label>
	<Input
		autocomplete="current-password"
		name="password"
		disabled={!$form.ftpConfig.authenticationRequired}
		bind:value={$form.ftpConfig.password}
		type="password"
	/>
</div>
