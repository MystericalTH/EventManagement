<script lang="ts">
	import Overlay from '$lib/components/Overlay.svelte';
	import { cross } from '$lib/assets/action-button-icons';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import { onMount } from 'svelte';
	let showOverlay = $state(false);
	let overlayMessage = $state('');
	let fname = $state('');
	let lname = $state('');
	let email = $state('');
	let phone = $state('');
	let githuburl = $state('');
	let interest = $state('');
	let reason = $state('');
	let getDetail = async () => {
		const response: Response = await fetch('/api/login/callback');
		let loginInfo = await response.json();

		console.log(response.status);
		if (loginInfo.role == null) {
			window.location.href = '/api/login?role=default&redirect_uri=/signup';
		}
		email = loginInfo.user.email;
	};

	onMount(getDetail);

	const closeOverlay = () => {
		showOverlay = false;
		window.location.href = '/api/logout';
	};
	const handleMemberSubmit = async (event: Event) => {
		event.preventDefault();

		const formData = {
			fname,
			lname,
			email,
			phone,
			githuburl,
			interest,
			reason
		};

		try {
			const response = await fetch('/api/members', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(formData)
			});

			if (response.ok) {
				console.log('Form submitted successfully');
				overlayMessage = 'Request sent!';
			} else {
				console.error('Form submission failed');
				overlayMessage = 'Request failed. Please try again later.';
			}
			showOverlay = true;
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
				<label for="fname" class="mb-2 font-bold">First Name</label>
			</div>
			<div>
				<input
					type="text"
					id="name"
					bind:value={fname}
					required
					class="w-72 rounded border border-gray-300 p-2 text-base"
				/>
			</div>
			<div class="content-center text-right">
				<label for="lname" class="mb-2 font-bold">Last Name</label>
			</div>
			<div>
				<input
					type="text"
					id="name"
					bind:value={lname}
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
					value={email}
					class="w-72 rounded border border-gray-300 bg-gray-200 p-2 text-base text-gray-400"
					disabled
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

<Overlay {showOverlay} width="[200px]" height="[150px]">
	<div class="flex justify-end">
		<ActionButton imgsrc={cross} action={closeOverlay} width="20px" alt="Close" />
	</div>
	<div class="mt-4 h-16 w-56 flex-1 text-center text-xl">
		{overlayMessage}
	</div>
</Overlay>
