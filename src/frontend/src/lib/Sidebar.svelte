<script lang="ts">
	import { writable } from 'svelte/store';
	type SidebarItem = {
		text: string;
		href: string | null;
		subitems: SidebarItem[] | null;
	};
	export let items: Array<SidebarItem> = [];

	const openItems = writable(new Set<number>());

	function toggleItem(index: number) {
		openItems.update((set) => {
			if (set.has(index)) {
				set.delete(index);
			} else {
				set.add(index);
			}
			return set;
		});
	}
</script>

<aside>
	<nav>
		<ul>
			{#each items as item, index}
				<li>
					{#if item.subitems}
						<div
							on:click={() => toggleItem(index)}
							class="collapsible"
							role="button"
							aria-expanded={$openItems.has(index)}
							aria-controls={'submenu-' + index}
						>
							{item.text}
							<span class="arrow">{$openItems.has(index) ? '▼' : '▶'}</span>
						</div>
						{#if $openItems.has(index)}
							<ul class="subitems">
								{#each item.subitems as subitem}
									<li><a href={subitem.href}>{subitem.text}</a></li>
								{/each}
							</ul>
						{/if}
					{:else}
						<a href={item.href}>{item.text}</a>
					{/if}
				</li>
			{/each}
		</ul>
	</nav>
</aside>

<style>
	aside {
		width: 200px;
		background-color: #2c3e50;
		color: white;
		padding: 1rem;
	}
	nav ul {
		list-style: none;
		padding: 0;
	}
	nav li {
		margin: 0.5rem 0;
	}
	nav a {
		color: white;
		text-decoration: none;
	}
	.collapsible {
		cursor: pointer;
		display: flex;
		justify-content: space-between;
		align-items: center;
	}
	.arrow {
		margin-left: 10px;
	}
	.subitems {
		padding-left: 1rem;
	}
</style>
