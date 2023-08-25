<script lang="ts">
	import { Label, Input, Button, A } from 'flowbite-svelte';
</script>

<form method="POST" class="p-12">
	<div class="
grid grid-cols-4 gap-12
md:grid-cols-3">
		<div></div>
		<div class="
		col-span-2
		md:col-span-1">
			<div>
				<Label for="username">Username
					<Input name="username" class="mt-2" required/>
				</Label>
			</div>
			<br/>
			<div>
				<Label for="password">Password
					<Input name="password" type="password" class="mt-2" required/>
				</Label>
			</div>
			<br/>
			<div class="flex flex-col justify-center items-center">
				<Button type="submit">Submit</Button>
			</div>
		</div>
		<div></div>
	</div>
</form>

<div class="grid">
	<div class="m-auto">
		Don't have an account? <A href="/signup">Signup Here</A>
	</div>
</div>