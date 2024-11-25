<script lang="ts">	
	import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();
	import Overlay from '$lib/components/Overlay.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import { cross } from '$lib/assets/action-button-icons';
	import { checkCircle } from '$lib/assets/action-button-icons';
	import caution from '$lib/assets/caution.png';

	let activity = data.activity;

	let expectation = $state('');
	let selectedRole = $state('');
	let statusCode: number = $state(-1);
	let message: string = $state('');

	let showOverlay = $state(false);
	const closeOverlay = () => {
		showOverlay = false;
	};
	const closeOverlayAndExit = () => {
		closeOverlay();
		window.location.href = '/home';
	};

	const handleRegisterSubmit = async (event: Event) => {
		event.preventDefault();

		const formData = {
			role: selectedRole,
			expectation
		};

		try {
			const response = await fetch(`/api/activities/${activity.id}/registration`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(formData)
			});

			statusCode = response.status;
			let resJson = await response.json();
			if (response.ok) {
				message = resJson.message;
				console.log('Form submitted successfully');
			} else {
				message = resJson.error;
				console.error('Form submission failed');
			}
			showOverlay = true;
		} catch (error) {
			console.error('Error submitting form:', error);
		}
	};
</script>

<h1 class="my-10 text-2xl font-semibold">{activity.title}</h1>

<form onsubmit={handleRegisterSubmit} class="mx-auto flex w-72 flex-col">
	<label for="role" class="mb-2 font-bold">Select Role:</label>
	<select
		id="role"
		bind:value={selectedRole}
		required
		class="mb-4 rounded border border-gray-300 p-2 text-lg"
	>
		<option value="" disabled selected>Select a role</option>
		{#each activity.activityRoles as role}
			<option value={role}>{role}</option>
		{/each}
	</select>
	<label for="expectation" class="mb-2 mt-14 block font-bold"
		>What do you expect from this activity?</label
	>
	<textarea
		id="expectation"
		name="expectation"
		bind:value={expectation}
		required
		class="mb-14 h-60 rounded border border-gray-300 p-2 text-lg"
	></textarea>
	<button
		type="submit"
		class="mt-4 w-40 cursor-pointer rounded bg-blue-500 p-2 text-lg text-white hover:bg-blue-700"
		>Submit</button
	>
</form>

<Overlay {showOverlay}>
	<div class="flex h-full flex-col">
		<div class="flex flex-none justify-end">
			<ActionButton
				imgsrc={cross}
				action={statusCode == 200 ? closeOverlayAndExit : closeOverlay}
				width="20px"
				alt="Close"
			/>
		</div>
		<div class="flex flex-col items-center justify-center text-center">
			<img
				src={statusCode == 200 ? checkCircle.click : caution}
				width="64px"
				alt={statusCode == 200 ? 'successful' : 'caution'}
			/>
			<span class="mt-4 text-base">{message}</span>
		</div>
	</div></Overlay
>
