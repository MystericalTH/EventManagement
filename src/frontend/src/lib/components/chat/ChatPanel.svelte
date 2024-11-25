<script lang="ts">
	import type { ChatMessage } from '$lib/types';
	import { onMount, tick } from 'svelte';
	let newMessageBuffer = $state('');
	import { parseAndFormatDate } from '$lib/utils/time';
	let element: HTMLDivElement;
	let {
		sendMessage,
		messages,
		channelid,
		sender,
		recipient
	}: {
		sendMessage: (channelid: number, newMessageBuffer: string) => void;
		messages: ChatMessage[];
		channelid: number;
		sender: string;
		recipient: string;
	} = $props();
	let loading = $state(true);
	$inspect(messages);

	const scrollToBottom = async (node: HTMLDivElement) => {
		node.scroll({ top: node.scrollHeight, behavior: 'smooth' });
	};
	const jumpToBottom = async (node: HTMLDivElement) => {
		node.scroll({ top: node.scrollHeight, behavior: 'auto' });
	};

	onMount(() => {
		setTimeout(async () => {
			loading = false;
			await tick();
			jumpToBottom(element);
		}, 500); // Adjust the delay as needed
	});
	async function sendMessageInternal() {
		sendMessage(channelid, newMessageBuffer);
		if (newMessageBuffer.trim() !== '') {
			let newMessage: ChatMessage = {
				sender: sender.toLowerCase(),
				message: newMessageBuffer,
				timesent: new Date()
			};
			messages = [...messages, newMessage];
			newMessageBuffer = '';
		}
		await tick();
		scrollToBottom(element);
	}
</script>

<div class="min-w-0 flex-1">
	<div class=" flex h-[96vh] flex-1 flex-col">
		<div class="flex-1 overflow-y-auto border-b border-gray-300 p-2" bind:this={element}>
			{#if loading}
				<div class="flex h-full items-center justify-center">
					<p class="text-center text-2xl text-gray-400">Loading...</p>
				</div>
			{:else}
				{#each messages as message}
					<div
						class={`message mb-2 rounded border-2 bg-gray-100 ${message.sender == sender.toLowerCase() ? ' border-blue-200' : 'border-indigo-400'} p-2`}
					>
						<div class="flex justify-between pb-2">
							<span
								class={`${message.sender == sender.toLowerCase() ? 'text-blue-400' : 'text-indigo-400'}`}
							>
								{message.sender == sender.toLowerCase() ? 'You' : recipient}
							</span>
							<span class="text-sm text-gray-400">
								{parseAndFormatDate(new Date(message.timesent))}
							</span>
						</div>
						<span class="text-wrap break-words">{message.message}</span>
					</div>
				{/each}
			{/if}
		</div>
		<div class="input-panel flex border-t border-gray-300 p-2">
			<input
				type="text"
				bind:value={newMessageBuffer}
				onkeydown={(e) => e.key === 'Enter' && sendMessageInternal()}
				placeholder="Type your message..."
				class="mr-2 flex-1 rounded border border-gray-300 p-2"
			/>
			<button
				onclick={sendMessageInternal}
				class="cursor-pointer rounded bg-blue-500 p-2 px-4 text-white hover:bg-blue-700"
				>Send</button
			>
		</div>
	</div>
</div>
