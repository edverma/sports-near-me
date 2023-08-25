<script lang="ts">
	import Signup from "../../components/Signup.svelte";
	import { Button, Modal } from 'flowbite-svelte';
	import { goto } from '$app/navigation';
	import {beforeUpdate} from "svelte";

	/** @type {import('../../../.svelte-kit/types/src/routes').ActionData} */
	export let form;
	let modalOpen = true;

	beforeUpdate(() => {
		if (form?.success && !modalOpen) {
			goto('/login');
		}
	})
</script>

<Signup/>

{#if form?.success}
	<Modal title="Success" bind:open={modalOpen}>
		<div>
			Successfully Created User
		</div>
		<div class="center">
			<Button href="/login">Login</Button>
		</div>
	</Modal>
{/if}