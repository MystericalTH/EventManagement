<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	let channels: ChatChannelInfo[] = $state([]);
	import ChannelBox from '$lib/components/chat/ChannelBox.svelte';
	import type { ChatChannelInfo } from '$lib/types';
	let selected: number = $state(-1);
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
				console.log(channels);
			})
			.catch((err) => console.error(err));
		return;
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
</div>
