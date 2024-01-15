<script lang="ts">
	import { onMount } from 'svelte';

	const socket = new WebSocket('ws://localhost:8080');

	function sendMessage(inputText: any) {
		console.log(inputText);
		socket.send(inputText);
	}

	onMount(() => {
		socket.onmessage = (res) => {
			console.log(res);
		};
	});
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<form>
	<input type="text" />
	<button on:submit={sendMessage}>Send</button>
</form>
