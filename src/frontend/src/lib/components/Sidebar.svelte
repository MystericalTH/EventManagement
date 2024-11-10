<script lang="ts">
	import { writable } from 'svelte/store';
	export let items: {
		text: string;
		href: string | null;
		subitems: { text: string; href: string | null; subitems: null }[] | null;
	}[];
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

<div class="bg-gray-700">
	<div
		class="sidebar-container font-custom m-4 flex h-[calc(100vh-2rem)] w-36 flex-col justify-between text-base tracking-wider"
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
				{#if $openItems.has(item.text) && item.subitems}
					<ul class="subitems pl-2 text-sm">
						{#each item.subitems as subitem}
							<li class="my-2">
								<a
									class="text-white text-white no-underline hover:text-indigo-300 hover:drop-shadow-lg"
									href={subitem.href}>{subitem.text}</a
								>
							</li>
						{/each}
					</ul>
				{/if}
			{/each}
		</nav>
		<button class="my-2 flex items-center text-white hover:text-rose-300 hover:drop-shadow-lg">
			Logout
		</button>
	</div>
</div>
