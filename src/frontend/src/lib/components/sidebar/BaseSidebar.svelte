<script lang="ts">
	import { writable } from 'svelte/store';
	import type { SidebarItem } from '$lib/types/sidebar';
	export let items: SidebarItem[];
	export let hasLogout: boolean = true;
	const openItems = writable(new Set<string>());
	function toggleItem(item: string) {
		openItems.update((set) => {
			if (set.has(item)) {
				set.delete(item);
			} else {
				set.add(item);
			}
			return set;
		});
	}
</script>

<div
	class="flex h-[100] w-40 flex-col justify-between bg-gray-700 p-4 font-custom text-base tracking-wider"
>
	<nav>
		{#each items as item}
			<div
				role="button"
				class="collapsible mb-2 flex cursor-pointer items-center justify-between text-white hover:text-indigo-300 hover:drop-shadow-lg"
				on:click={() => toggleItem(item.text)}
			>
				<span>{item.text}</span>

				{#if item.subitems}
					<span class="arrow ml-2 text-xs">{$openItems.has(item.text) ? '▲' : '▼'}</span>
				{/if}
			</div>
			{#if item.subitems && $openItems.has(item.text)}
				<ul class="subitems pl-2 text-sm">
					{#each item.subitems as subitem}
						<li class="my-2">
							<a
								class="text-white no-underline hover:text-indigo-300 hover:drop-shadow-lg"
								href={subitem.href}>{subitem.text}</a
							>
						</li>
					{/each}
				</ul>
			{/if}
		{/each}
	</nav>
	<div>
		{#if hasLogout}
			<button class="my-2 items-center text-white hover:text-rose-300 hover:drop-shadow-lg">
				<a href="/api/logout">Logout</a>
			</button>
		{/if}
	</div>
</div>
