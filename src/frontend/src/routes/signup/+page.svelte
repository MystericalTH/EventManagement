<script lang="ts">
	let name = '';
	let email = '';
	let phone = '';
	let githuburl = '';
	let interest = '';
	let reason = '';

	const handleMemberSubmit = async (event: Event) => {
		event.preventDefault();

		const formData = {
			name,
			email,
			phone,
			githuburl,
			interest,
			reason
		};

		try {
			const response = await fetch('/api/member', {
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

<div class="min-w-[700px] border-2 border-gray-100 p-5 shadow-md">
	<h1 class="my-5 text-center text-4xl font-bold">Signup as Our Member</h1>
	<form onsubmit={handleMemberSubmit} class="flex flex-col py-4">
		<div class="grid grid-cols-[40%_60%] gap-4 pr-2">
			<div class="content-center text-right">
				<label for="name" class="mb-2 font-bold">Name</label>
			</div>
			<div>
				<input
					type="text"
					id="name"
					bind:value={name}
					required
					class="w-72 rounded border border-gray-300 p-2 text-base"
				/>
			</div>
			<div class="content-center text-right">
				<label for="email" class="mb-2 font-bold">Email</label>
			</div>
			<div>
				<input
					type="email"
					id="email"
					bind:value={email}
					required
					class="w-72 rounded border border-gray-300 p-2 text-base"
				/>
			</div>
			<div class="content-center text-right">
				<label for="phone" class="mb-2 font-bold">Phone Number</label>
			</div>
			<div>
				<input
					type="tel"
					id="phone"
					bind:value={phone}
					required
					class="rounded border border-gray-300 p-2 text-base"
				/>
			</div>
			<div class="content-center text-right">
				<label for="githuburl" class="mb-2 font-bold">GitHub</label>
			</div>
			<div>
				<input
					type="url"
					id="githuburl"
					bind:value={githuburl}
					required
					class="w-72 rounded border border-gray-300 p-2 text-base"
				/>
			</div>
			<div class="mt-4 content-start text-right">
				<label for="interest" class="mb-2 font-bold">What is your interest</label>
			</div>
			<div>
				<textarea
					id="interest"
					bind:value={interest}
					required
					class="min-w-72 max-w-fit resize rounded border border-gray-300 p-2 text-base"
				></textarea>
			</div>
			<div class="mt-4 content-start text-right">
				<label for="reason" class="mb-2 font-bold">Why do you want to join the club</label>
			</div>
			<div>
				<textarea
					id="reason"
					bind:value={reason}
					required
					class="min-w-72 max-w-fit resize rounded border border-gray-300 p-2 text-base"
				></textarea>
			</div>
		</div>
		<div class="flex justify-center">
			<button
				type="submit"
				class="mt-4 w-40 cursor-pointer rounded bg-blue-500 p-2 text-base text-white hover:bg-blue-700"
				>Submit</button
			>
		</div>
	</form>
</div>