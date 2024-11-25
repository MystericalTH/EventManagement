<script lang="ts">
	import type { ChatChannelInfo } from '$lib/types';
	let {
		href,
		data,
		text,
		selected,
		onclick
	}: { data: ChatChannelInfo; text: string; selected: number; onclick: () => void; href: string } =
		$props();
</script>

<a {href}>
	<button
		{onclick}
		class={`h-16 w-full min-w-24 border-b-2 border-gray-200 p-1 text-left hover:bg-indigo-200 ${selected == data.id ? 'bg-indigo-100' : 'bg-white'}`}
	>
		<div class="... h-6 truncate text-base text-gray-800">{data.fname + ' ' + data.lname}</div>
		<span class="... block truncate text-sm text-gray-500">{text}</span></button
	></a
>
