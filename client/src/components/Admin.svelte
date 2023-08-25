<script>
	import { Button } from 'flowbite-svelte';

	export let username;
</script>

<div class="grid grid-cols-3">
	<div></div>
	<div class="justify-self-center grid grid-cols-1">
		<Button class="" href="/admin/issue-creator">Create an Issue</Button>
		<br/>
		<Button class="" href="/admin/addresses">Update Your Recipients</Button>
		<br/>
		<Button class="" href={`/user/${username}`}>Go to Your Website</Button>
	</div>
	<div></div>
</div>