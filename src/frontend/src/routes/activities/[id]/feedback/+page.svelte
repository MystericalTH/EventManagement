<script lang="ts">
	import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();

	let feedback = $state('');

	const handleFeedbackSubmit = async (event: Event) => {
		event.preventDefault();

		const formData = {
			feedback
		};

		try {
			const response = await fetch('/api/feedback/${activity.id}/submit', {
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

<form onsubmit={handleFeedbackSubmit} class="mx-auto flex w-72 flex-col">
	<label for="feedback" class="mb-2 mt-14 block font-bold">Feedback</label>
	<textarea
		id="feedback"
		name="feedback"
		bind:value={feedback}
		required
		class="mb-14 h-60 rounded border border-gray-300 p-2 text-lg"
	></textarea>
	<button
		type="submit"
		class="mt-4 cursor-pointer rounded bg-blue-500 p-2 text-lg text-white hover:bg-blue-700"
		>Submit</button
	>
</form>
