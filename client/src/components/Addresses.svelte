<script>
	import { Label, Input, Button } from 'flowbite-svelte';
	import {derived, writable} from 'svelte/store'
	import {onMount} from "svelte";

	export let addresses;

	const addressesStore = writable([]);
	addressesStore.set(addresses);
	const addressesStoreJSON = derived(addressesStore, $addressesStore => JSON.stringify($addressesStore));
	const removeFromAddressesStore = (i) => addressesStore.update(addresses => {
		addresses.splice(i, 1);
		return addresses;
	});
	const addToAddressesStore = () => $addressesStore = [...$addressesStore, {}]

	const numAddressesLimit = 20;

	onMount(() => {
		if ($addressesStore.length === 0) {
			addToAddressesStore();
		}
	})
</script>

<form method="POST" class="grid grid-cols-3 p-12">
	<Input class="hidden" name="addresses" value={$addressesStoreJSON}></Input>
	{#each $addressesStore as address, i (i)}
		<div></div>
		<div>
			<div class="mb-2">
				<Label for="name-{i}" class="mb-2">Name</Label>
				<Input name="name-{i}" id="name-{i}" type="text" placeholder="Name" bind:value={$addressesStore[i].name} required />
			</div>

			<div class="mb-2">
				<Label for="street-address-{i}" class="mb-2">Street Address</Label>
				<Input name="street-address-{i}" id="street-address-{i}" type="text" placeholder="Street Address"
				       bind:value={$addressesStore[i].street_address} required />
			</div>

			<div class="mb-2">
				<Label for="extra-address-line-{i}" class="mb-2">Line 2</Label>
				<Input name="extra-address-line-{i}" id="extra-address-line-{i}" type="text" placeholder="Line 2"
				       bind:value={$addressesStore[i].extra_address_line}/>
			</div>

			<div class="mb-2">
				<Label for="dependent-locality-{i}" class="mb-2">City</Label>
				<Input name="dependent-locality-{i}" id="dependent-locality-{i}" type="text" placeholder="City"
				       bind:value={$addressesStore[i].dependent_locality} required />
			</div>

			<div class="mb-2">
				<Label for="locality-{i}" class="mb-2">State</Label>
				<Input name="locality-{i}" id="locality-{i}" type="text" placeholder="State"
				       bind:value={$addressesStore[i].locality} required />
			</div>

			<div class="mb-2">
				<Label for="post-code-{i}" class="mb-2">Zip Code</Label>
				<Input name="post-code-{i}" id="post-code-{i}" type="text" placeholder="Zip Code"
				       bind:value={$addressesStore[i].post_code} required />
			</div>

			<br/><hr/><br/>
		</div>
		<div class="ml-16 my-auto">
			<Button pill={true} color="red" on:click={removeFromAddressesStore(i)}>Delete</Button>
		</div>
	{/each}
	<br/>
	<div class="m-auto mb-12">
		{#if $addressesStore.length < numAddressesLimit}
			<Button color="light" pill={true} on:click={addToAddressesStore}>+</Button>
		{/if}
		</div>
	<div></div>
	<div></div>
	<div class="m-auto">
		<Button type="submit">Submit</Button>
	</div>
</form>
