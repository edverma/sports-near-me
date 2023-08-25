<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import {page} from "$app/stores";
	import DarkModeToggle from "./DarkModeToggle.svelte";

	export let heading: string;

	$: home = $page.url.pathname === `/user/${$page.params.user}`;
	$: issueDateSlug = $page.params.issueDate != null;
	$: linkPath = home ? `/user/${$page.params.user}/issues` : issueDateSlug ?
		`/user/${$page.params.user}/issues` : `/user/${$page.params.user}`;
	$: linkText = home ? 'Individual Issues' : issueDateSlug ? 'Individual Issues' : 'All Issues';

</script>

<div class="
grid items-center pt-2
grid-cols-3 sm:gap-4
">
	<div class="">
		<DarkModeToggle/>
	</div>
	<div class="text-center text-2xl
	sm:text-4xl
	print:col-span-3 ">
			{heading != "" ? heading : $page.params.user}
	</div>
	<div class="justify-self-end print:hidden">
		<Button color="light" size="xs" href="/admin">Admin</Button>
	</div>
</div>
<div class="print:hidden">
	<br/>
</div>
<div class="
text-center text-lg text-fuchsia-800 font-semibold
sm:text-xl
print:hidden">
	<a href={linkPath}>{linkText}</a>
</div>
<div class="print:hidden">
	<br>
</div>
