<script lang="ts">
	import type { ImageSource } from '$lib/types';
	import type { HTMLButtonAttributes } from 'svelte/elements';
	let {
		imgsrc,
		action,
		alt,
		width,
		...props
	}: {
		imgsrc: ImageSource;
		action: () => void;
		alt: string;
		width: string;
	} & HTMLButtonAttributes = $props();
	let src = $state(imgsrc.idle);
	let srcEnter = () => {
		src = imgsrc.enter;
	};
	let srcLeave = () => {
		src = imgsrc.idle;
	};
	let srcClick = () => {
		src = imgsrc.click;
	};
</script>

<button
	onmousedown={srcClick}
	onclick={action}
	onmouseenter={srcEnter}
	onmouseleave={srcLeave}
	{...props}
>
	<img {src} {alt} {width} />
</button>
