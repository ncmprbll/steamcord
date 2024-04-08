<script lang="ts">
    export let data;
	let loginVisible = false;

	import Nav from '$lib/components/Nav.svelte';
	import Error from '$lib/components/Error.svelte';
	import Login from '$lib/components/Login.svelte';
</script>

<svelte:head>
    <title>Steamcord</title> 
</svelte:head>

<Nav data={data} bind:loginVisible/>

<main>
	<Error />
	<slot />
</main>

<footer>
	<div class="footer-container">
		© 2024 Steamcord. All rights reserved.
	</div>
</footer>

<Login locale={data.locale} bind:visible={loginVisible}/>