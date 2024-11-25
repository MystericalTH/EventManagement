<script lang="ts">
	import ChatPanel from '$lib/components/chat/ChatPanel.svelte';
	import { page } from '$app/stores';
	import type { ChatMessage } from '$lib/types';
	import { onMount } from 'svelte';
	let messages: ChatMessage[] = $state([]);
	let id: number = $state(-1);

	const fetchData = async (channelid: number) => {
		id = channelid;
		console.log('fetch data');
		const response = await fetch(`/api/chats/${id}`);
		const result: { chats: ChatMessage[] } = await response.json();
		messages = result.chats;
		console.log(messages);
	};
	// Watch for changes in the URL

	onMount(() => {
		page.subscribe(($page) => {
			fetchData(parseInt($page.params.id));
		});
	});
	let sendMessage = (channelid: number, message: string) => {
		console.log(channelid);
		console.log(message);
		if (message.trim() === '') return;
		fetch(`/api/chats`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ developerid: channelid, message })
		})
			.then((res) => res.json())
			.catch((err) => console.error(err));
	};
</script>

{#key $page.url.pathname}
	<ChatPanel {sendMessage} {messages} channelid={id} sender="Developer" recipient="Admin"
	></ChatPanel>
{/key}
