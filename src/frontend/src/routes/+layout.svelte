<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/Sidebar.svelte';

	const sidebarItems = [
		{
			text: 'Sign Up',
			href: '/signup',
			subitems: null
		},
		{
			text: 'Log In',
			href: null,
			subitems: [
				{ text: 'Member', href: '/login', subitems: null },
				{ text: 'Admin', href: '/logout', subitems: null },
				{ text: 'Developer', href: '/logout', subitems: null }
			]
		}
	];
</script>

<div class="layout">
	<Sidebar items={sidebarItems} />
	<main>
		<slot></slot>
	</main>
</div>

<style>
	.layout {
		display: flex;
		height: 100vh;
	}
	main {
		flex: 1;
		padding: 1rem;
		background-color: #ecf0f1;
	}
</style>