<script lang="ts">
	import { writable } from 'svelte/store';
	import type { SidebarItem } from '$lib/types';
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
	class="flex h-[100] w-48 flex-none flex-col justify-between bg-gray-700 p-4 font-custom text-base tracking-wider"
>
	<nav>
		<ul class="list-none">
			{#each items as item}
				<li class="mb-2">
					{#if item.subitems}
						<button
							class="space-between flex w-full flex-row items-center justify-between text-white hover:text-indigo-300 hover:drop-shadow-lg"
							on:click={() => toggleItem(item.text)}
						>
							<span>{item.text}</span>
							<span class="arrow ml-2 text-sm">{$openItems.has(item.text) ? '▲' : '▼'}</span>
						</button>
					{:else}
						<a
							class="text-white no-underline hover:text-indigo-300 hover:drop-shadow-lg"
							href={item.href}>{item.text}</a
						>
					{/if}{#if item.subitems && $openItems.has(item.text)}
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
				</li>
			{/each}
		</ul>
	</nav>
	<div>
		{#if hasLogout}
			<button class="my-2 items-center text-white hover:text-rose-300 hover:drop-shadow-lg">
				<a href="/api/logout">Logout</a>
			</button>
		{/if}
	</div>
</div>
