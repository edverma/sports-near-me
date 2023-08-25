<script lang="ts">
	import { Label, Input, Button } from 'flowbite-svelte';
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
				<Label for="first_name">First Name
					<Input name="first_name" class="mt-2" required/>
				</Label>
			</div>
			<br/>
			<div>
				<Label for="last_name">Last Name
					<Input name="last_name" class="mt-2" required/>
				</Label>
			</div>
			<br/>
			<div>
				<Label for="email">Email
					<Input name="email" class="mt-2" required/>
				</Label>
			</div>
			<br/>
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