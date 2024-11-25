<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import type { ChatChannelInfo } from '$lib/types';
	import { ChannelBox } from '$lib/components/chat';

	let channels: ChatChannelInfo[] = $state([]);
	let selected: number = $state(-1);
	let selectedName: string = $state('');

	let { children } = $props();
	let loadData = () => {
		fetch('/api/chats')
			.then((res) => {
				return res.json();
			})
			.then((resJson) => {
				channels = resJson.data;
				channels.sort((a: ChatChannelInfo, b: ChatChannelInfo) => {
					if (a.timesent == null) return 1;
					else if (b.timesent == null) return -1;
					console.log('message sent');
					return new Date(b.timesent).getTime() - new Date(a.timesent).getTime();
				});
				console.log($state.snapshot(channels));
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

<div class="flex flex-1 flex-row border-2 border-gray-200">
	<div
		class="flex h-full min-w-48 max-w-64 resize-x flex-col overflow-y-scroll border-r-2 border-r-gray-200 p-2"
	>
		{#each channels as channel}
			{#key channel.id}
				<ChannelBox
					href={`/admin/chat/dev/${channel.id}`}
					text={channel.message}
					data={channel}
					onclick={() => {
						selected = channel.id;
						selectedName = channel.fname + ' ' + channel.lname;
					}}
					{selected}
				/>
			{/key}
		{/each}
	</div>
	{@render children()}
</div>
