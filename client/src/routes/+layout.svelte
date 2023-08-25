<script>
	import "../app.css";
	import Analytics from '../components/Analytics.svelte'
</script>

<div class="dark:bg-zinc-900 dark:text-zinc-200 min-h-screen">
	<Analytics />
	<slot/>
</div>