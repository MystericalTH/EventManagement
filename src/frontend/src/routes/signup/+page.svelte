<script lang="ts">
	let fName = '';
	let lName = '';
	let email = '';
	let phone = '';
	let githubUrl = '';
	let interest = '';
	let reason = '';

	const handleMemberSubmit = async (event: Event) => {
		event.preventDefault();

		const formData = {
			fName,
			lName,
			email,
			phone,
			githubUrl,
			interest,
			reason
		};

		try {
			const response = await fetch('/api/member/request', {
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

<h1 class="my-5 text-center text-4xl font-bold">Sign Up</h1>
<form on:submit={handleMemberSubmit} class="mx-auto flex w-72 flex-col">
	<div class="mb-4">
        <label for="fName" class="mb-2 font-bold">First Name:</label>
        <input
            type="text"
            id="fName"
            bind:value={fName}
            required
            class="rounded border border-gray-300 p-2 text-lg"
        />
    </div>
    <div class="mb-4">
        <label for="lName" class="mb-2 font-bold">Last Name:</label>
        <input
            type="text"
            id="lName"
            bind:value={lName}
            required
            class="rounded border border-gray-300 p-2 text-lg"
        />
    </div>
	<div class="mb-4">
		<label for="email" class="mb-2 font-bold">Email:</label>
		<input
			type="email"
			id="email"
			bind:value={email}
			required
			class="rounded border border-gray-300 p-2 text-lg"
		/>
	</div>
	<div class="mb-4">
		<label for="phone" class="mb-2 font-bold">Phone:</label>
		<input
			type="tel"
			id="phone"
			bind:value={phone}
			required
			class="rounded border border-gray-300 p-2 text-lg"
		/>
	</div>
	<div class="mb-4">
		<label for="githubUrl" class="mb-2 font-bold">GitHub:</label>
		<input
			type="url"
			id="githubUrl"
			bind:value={githubUrl}
			required
			class="rounded border border-gray-300 p-2 text-lg"
		/>
	</div>
	<div class="mb-4">
		<label for="interest" class="mb-2 font-bold">What is your interest:</label>
		<textarea
			id="interest"
			bind:value={interest}
			required
			class="rounded border border-gray-300 p-2 text-lg"
		></textarea>
	</div>
	<div class="mb-4">
		<label for="reason" class="mb-2 font-bold">Why do you want to join the club:</label>
		<textarea
			id="reason"
			bind:value={reason}
			required
			class="rounded border border-gray-300 p-2 text-lg"
		></textarea>
	</div>
	<button
		type="submit"
		class="mt-4 cursor-pointer rounded bg-blue-500 p-2 text-lg text-white hover:bg-blue-700"
		>Submit</button
	>
</form>
