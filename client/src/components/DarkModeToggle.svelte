<script lang="ts">
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { writable } from 'svelte/store';

	let dark = writable(browser ? localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches) : false);
	onMount(() => {
		if (browser) {
			dark.update(() => localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches));
		}
	});

	dark.subscribe(value => {
		if (browser) {
			localStorage.theme = value ? 'dark' : 'light';
			if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
				document.documentElement.classList.add('dark')
			} else {
				document.documentElement.classList.remove('dark')
			}
		}
	})
</script>

<label class="relative inline-flex items-center cursor-pointer print:hidden">
	<input type="checkbox" class="sr-only peer" bind:checked={$dark}>
	<div class="w-11 h-6 bg-gray-200 rounded-full peer dark:peer-focus:ring-blue-800 dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-fuchsia-600"></div>
	<span class="
	hidden ml-3 text-sm font-medium text-gray-900 dark:text-gray-300
	sm:inline
">
		Dark Mode
	</span>
</label>