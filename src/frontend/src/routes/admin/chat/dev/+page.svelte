<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	let channels: ChatChannelInfo[] = $state([]);
	import ChannelBox from '$lib/components/chat/ChannelBox.svelte';
	import ChatPanel from '$lib/components/chat/ChatPanel.svelte';
	import type { ChatChannelInfo } from '$lib/types';
	let selected: number = $state(-1);
	let newMessage: string = $state('');

	let loadData = () => {
		fetch('/api/chats')
			.then((res) => {
				return res.json();
			})
			.then((resJson) => {
				channels = resJson.data;
				channels.sort((a: ChatChannelInfo, b: ChatChannelInfo) => {
					if (a == null) return -1;
					else if (b == null) return 1;
					return a.timesent.getTime() - b.timesent.getTime();
				});
				console.log($state.snapshot(channels));
			})
			.catch((err) => console.error(err));
		return;
	};

	let sendMessage = (chatid: number, message: string) => {
        if (message.trim() === '') return;

        fetch('/api/chats', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ chatid, message })
        })
            .then((res) => res.json())
            .then((resJson) => {
                if (resJson.message === 'Chat submitted successfully') {
                    loadData();
                    newMessage = '';
                } else {
                    console.error('Failed to send message');
                }
            })
            .catch((err) => console.error(err));
    };

    let sendMessageWithId = (message: string) => {
        sendMessage(selected, message);
    };

	onMount(loadData);
	$effect(() => {
		const interval = setInterval(() => {
			loadData();
		}, 5000);

		// Cleanup interval on component destroy
		onDestroy(() => {
			clearInterval(interval);
		});
	});
</script>

<div class="flex flex-1 flex-row">
	<div class="flex h-full min-w-48 flex-col overflow-y-scroll bg-gray-200 p-2">
		{#each channels as channel}
			{#key channel.id}
				<ChannelBox
					text={channel.message}
					data={channel}
					onclick={() => {
						selected = channel.id;
					}}
					{selected}
				/>
			{/key}
		{/each}
	</div>
	{#if selected !== -1}
		<ChatPanel chatid={selected} sendMessage={() => sendMessage(selected, newMessage)} />
    {/if}
</div>
