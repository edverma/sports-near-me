<script>
    import Admin from "../../../components/Admin.svelte";
    import {browser} from "$app/environment";

    /** @type {import('../../../.svelte-kit/types/src/routes').PageData} */
    export let data;
    const user = data.user;
    const authz = data.success;
</script>

{#if authz}
    <Admin username={user.username}/>
{:else if authz !== null && browser}
    {goto('/login')}
{/if}
