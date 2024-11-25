<script>
	import { page } from '$app/stores';
	import { error } from '@sveltejs/kit';
</script>

<div class="center">
	<div>
		<h1 class="text-6xl font-bold">{$page.status}</h1>
		<p class="mt-4 text-3xl">{$page.error?.message}</p>
	</div>
</div>

<style>
	.center {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
		text-align: center;
	}
</style>
