<script lang="ts">
	import Login from "../../components/Login.svelte";
	import {beforeUpdate} from "svelte";
	import {goto} from "$app/navigation";
	import {browser} from "$app/environment";

	/** @type {import('../../../.svelte-kit/types/src/routes').ActionData} */
	export let form;

	beforeUpdate(() => {
		if (form?.success && browser) {
			goto('/');
		}
	})
</script>

<Login/>
