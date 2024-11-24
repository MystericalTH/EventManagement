<script lang="ts">
	import type { Activity } from '$lib/types';
	// export let data: { activity: Activity; activityRoles: string[] };
	import type { PageData } from './$types';

	let { data }: { data: { activity: Activity } } = $props();

	let expectation = $state('');
	let selectedRole = $state('');

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

			if (response.ok) {
				console.log('Form submitted successfully');
			} else {
				console.error('Form submission failed');
			}
		} catch (error) {
			console.error('Error submitting form:', error);
		}
	};
</script>

<h1 class="my-10 text-2xl font-semibold">{data.activity.title}</h1>

<form onsubmit={handleRegisterSubmit} class="mx-auto flex w-72 flex-col">
	<label for="role" class="mb-2 font-bold">Select Role:</label>
	<select
		id="role"
		bind:value={selectedRole}
		required
		class="mb-4 rounded border border-gray-300 p-2 text-lg"
	>
		<option value="" disabled selected>Select a role</option>
		{#each data.activity.activityRoles as role}
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
