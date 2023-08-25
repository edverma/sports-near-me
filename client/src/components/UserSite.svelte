<script lang="ts">
	import Issue from './Issue.svelte'
	import type {DbIssue} from "../models";

	export let dbIssues: [DbIssue] = [];
</script>

{#if dbIssues !== null}
	{#each dbIssues as dbIssue}
		<Issue dbIssue={dbIssue}/>
		<br/><br/>
	{/each}
{/if}