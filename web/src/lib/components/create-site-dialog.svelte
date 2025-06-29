<script lang="ts">
	import { Plus } from '@lucide/svelte';
	import CreateSiteForm from './create-site-form.svelte';
	import * as Dialog from './ui/dialog';
	import { buttonVariants } from './ui/button';

	const data: {
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
	} = {
		type: 'ftp',
		name: '',
		ftpConfig: {
			port: 21,
			authenticationRequired: false,
			username: undefined,
			password: undefined
		},
		address: '',
		httpConfig: {
			headers: {}
		}
	};
</script>

<Dialog.Root>
	<Dialog.Trigger
		type="button"
		class={buttonVariants({ variant: 'outline', size: 'icon', class: 'cursor-pointer' })}
	>
		<Plus class="size-4" />
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header class="text-xl leading-none font-semibold">Create a new site</Dialog.Header>
		<div class="px-8">
			<CreateSiteForm {data} />
		</div>
	</Dialog.Content>
</Dialog.Root>
