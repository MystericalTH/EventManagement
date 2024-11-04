<script lang="ts">
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<h1>{data.post.title}</h1>
<div>{@html data.post.content}</div>