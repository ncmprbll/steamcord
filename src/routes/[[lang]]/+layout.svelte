<script lang="ts">
    export let data;
	let visible = false;

	import Nav from '$lib/components/Nav.svelte';
	import Login from '$lib/components/Login.svelte';
</script>

<svelte:head>
    <title>Steamcord</title> 
</svelte:head>

<Nav {data} bind:visible/>

<slot />

<footer>
	<div class="footer-container">
		© 2024 Steamcord. All rights reserved.
	</div>
</footer>

<Login bind:visible/>