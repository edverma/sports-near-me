<script>
	import { Label, Input, Textarea, Button, NumberInput, P } from 'flowbite-svelte';
	import { writable } from 'svelte/store';
	import {browser} from "$app/environment";

	const maxArticles = 8;
	let numArticles;
	let gridCols;
	let subheadingsStore = writable([]);
	let contentStore = writable([]);
	const date = new Date()
	const currentDateString = date.getFullYear() + '-' + (date.getMonth()+1).toLocaleString(undefined, {minimumIntegerDigits: 2})
		+ '-' + date.getDate().toLocaleString(undefined, {minimumIntegerDigits: 2});

	if (browser) {
		const storedNumArticles = localStorage.getItem('numArticles');
		const storedGridCols = localStorage.getItem('gridCols');
		numArticles = writable(parseInt(storedNumArticles) || 4);
		gridCols = writable(parseInt(storedGridCols) || 2);
		numArticles.subscribe((value) => localStorage.setItem('numArticles', value));
		gridCols.subscribe((value) => localStorage.setItem('gridCols', value));

		let storedSubheadings = [];
		let storedContent = [];
		for (let i = 0; i < maxArticles; i++) {
			storedSubheadings[i] = localStorage.getItem(`subheadings-${i}`);
			storedContent[i] = localStorage.getItem(`content-${i}`);
			subheadingsStore.update(arr => {
				arr.push(storedSubheadings[i] || '');
				return arr;
			});
			contentStore.update(arr => {
				arr.push(storedContent[i] || '');
				return arr;
			});
		}

		subheadingsStore.subscribe(arr => {
			for (let i = 0; i < arr.length; i++) {
				localStorage.setItem(`subheadings-${i}`, arr[i])
			}
		});
		contentStore.subscribe(arr => {
			for (let i = 0; i < arr.length; i++) {
				localStorage.setItem(`content-${i}`, arr[i])
			}
		});
	}
</script>

<form method="POST" class="p-12">
	<div class="grid grid-cols-3 gap-12">
		<div>
			<Label>How many articles?
				<NumberInput name="numArticles" min="1" max="8" class="mt-2" bind:value={$numArticles} />
			</Label>
		</div>
		<div>
			<Label>How many columns?
				<NumberInput name="numColumns" min="1" max="3" class="mt-2" bind:value={$gridCols} />
			</Label>
		</div>
		<div>
			<Label class="mb-2">When should we send it?</Label>
				<Input type="date" name="sendDate" value={currentDateString} required/>
		</div>
	</div>
	<div>
		<br/>
		<hr/>
		<br/>
	</div>
	<div>
		<div class="grid grid-cols-{$gridCols} gap-12">
			{#each Array($numArticles) as _, i (i)}
				<div class="{ i === $numArticles - 1 && i % $gridCols === 0 ? `col-span-${$gridCols}` : ''}">
					<Label for="content-{i}" class="mb-2">Article {i+1}</Label>
					<Input name="subheading-{i}" id="subheading-{i}" bind:value={$subheadingsStore[i]} type="text" placeholder="Heading {i+1}" required />
					<br/>
					<Textarea name="content-{i}" id="content-{i}" bind:value={$contentStore[i]} placeholder="Content for Article {i+1}" rows="12"/>
				</div>
			{/each}
		</div>
		<br/>
		<div class="flex flex-col justify-center items-center">
			<Button type="submit">Submit</Button>
		</div>
	</div>
</form>
<br/>