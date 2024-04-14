<script lang="ts">
	import { flip } from 'svelte/animate';
	import { cart } from '$lib/stores/users';

    export let data;
	let loginVisible = false;

	import Nav from '$lib/components/Nav.svelte';
	import Error from '$lib/components/Error.svelte';
	import Login from '$lib/components/Login.svelte';
</script>

<svelte:head>
    <title>Steamcord</title> 
</svelte:head>

<Nav bind:loginVisible/>

<main>
	<Error error={data.error}/>
	<slot />
</main>

<footer>
	<div class="footer-container">
		© 2024 Steamcord. All rights reserved.
	</div>
</footer>

<Login bind:visible={loginVisible}/>