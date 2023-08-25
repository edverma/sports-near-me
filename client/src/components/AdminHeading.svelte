<script>
	import { Heading, Button } from 'flowbite-svelte';
	import DarkModeToggle from "./DarkModeToggle.svelte";
</script>

<div class="grid grid-cols-3">
	<div class="justify-self-start ml-2 mt-4
sm:ml-12">
		<DarkModeToggle/>
	</div>
	<div class="justify-self-center mt-4">
		<Heading class="text-center">Friendmail Admin</Heading>
	</div>
	<div class="justify-self-end mr-2 mt-4
sm:mr-12">
		<form method="POST" action="/admin?/logout">
			<Button color="light" size="xs" type="submit">Logout</Button>
		</form>
	</div>
</div>
<br/><br/>