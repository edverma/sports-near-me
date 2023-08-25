<script lang="ts">
	import SvelteMarkdown from "svelte-markdown";
	import type {DbIssue} from "../models";
	import {formatDateReadable} from "../helper.js";
	import PrintableIssue from "./PrintableIssue.svelte";

	export let dbIssue: DbIssue;

	const gridCols = dbIssue.columns > 0 && dbIssue.columns < 4 ? dbIssue.columns : 1;
</script>

<div class="hidden print:inline">
	<PrintableIssue date={formatDateReadable(dbIssue.send_date)} dbIssue={dbIssue} gridCols={gridCols}/>
</div>
<div class="print:hidden">
	<div class="row-xs">
		<p class="text-fuchsia-500">
			{formatDateReadable(dbIssue.send_date)}
		</p>
	</div>
	<div class="grid grid-cols-1 sm:grid-cols-{gridCols}">
		{#each dbIssue.subheadings as _, i}
				<div class="
				border-2 pl-4 pr-4 pb-4 border-fuchsia-400
				sm:border-8 sm:pl-8 sm:pb-8">
					<div class="text-center text-xl print:text-sm">
						<SvelteMarkdown source={dbIssue.subheadings[i]}/>
					</div>
					<br/>
					<div>
						<SvelteMarkdown source={dbIssue.content[i]}/>
					</div>
				</div>
		{/each}
	</div>
	<div><br/></div>
</div>