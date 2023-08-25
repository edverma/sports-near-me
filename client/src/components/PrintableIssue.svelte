<script lang="ts">
	import SvelteMarkdown from "svelte-markdown";
	import type {DbIssue} from "../models";

	export let date: string;
	export let gridCols: number;
	export let dbIssue: DbIssue;

	let textSize = ''
	const contentLength = dbIssue.content.reduce((len, content) => len + (content.length / gridCols), 0);
	const lengthTextSizeMap = new Map();
	lengthTextSizeMap.set(0, '[font-size:14px]');
	lengthTextSizeMap.set(2000, '[font-size:12px]');
	lengthTextSizeMap.set(2500, '[font-size:11px]');
	lengthTextSizeMap.set(3000, '[font-size:10px]');
	lengthTextSizeMap.set(3500, '[font-size:9px]');
	lengthTextSizeMap.set(4000, '[font-size:8px]');
	let highest = 0;
	for (let [key, value] of lengthTextSizeMap) {
		if (key >= highest && contentLength > key) {
			textSize = value;
			highest = key;
		}
	}
</script>

<div class="
	grid grid-cols-{gridCols}
	">
	{#each dbIssue.subheadings as _, i}
		<div class="border-2 pl-8 pr-4 pb-8 border-fuchsia-400">
			<div class="text-center text-xl print:text-sm">
				<SvelteMarkdown source={dbIssue.subheadings[i]}/>
			</div>
			<br/>
			<div class={textSize}>
				<SvelteMarkdown source={dbIssue.content[i]}/>
			</div>
		</div>
	{/each}
</div>
<div class="grid grid-cols-3">
	<div>
	</div>
	<div>
		<p class="text-fuchsia-500 text-[10px] text-center">
			{date}
		</p>
	</div>
	<div>
		<p class="text-fuchsia-500 text-[10px] text-end">
			friendmail.co/user/{dbIssue.username}
		</p>
	</div>
</div>