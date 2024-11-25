<script lang="ts">
	import Overlay from '$lib/components/Overlay.svelte';
	import ActionButton from '$lib/components/ActionButton.svelte';
	import { cross } from '$lib/assets/action-button-icons';
	let title = $state('');
	let startDate = $state('');
	let endDate = $state('');
	let maxParticipant = $state(1);
	let format = $state('');
	let description = $state('');
	let advisor = $state('');
	let startTime = $state('');
	let endTime = $state('');
	let activityRole: string[] = $state([]);
	let newActivityRole = $state('');
	import { checkCircle } from '$lib/assets/action-button-icons';
	import caution from '$lib/assets/caution.png';

	let showOverlay = $state(false);
	const closeOverlay = () => {
		showOverlay = false;
	};
	const closeOverlayAndExit = () => {
		closeOverlay();
		window.location.href = '/home';
	};
	const addActivityRole = () => {
		if (newActivityRole.trim() !== '' && !activityRole.includes(newActivityRole.trim())) {
			activityRole = [...activityRole, newActivityRole.trim()];
			newActivityRole = '';
		}
	};
	const deleteActivityRole = (roleToDelete: string) => {
		activityRole = activityRole.filter((role) => role !== roleToDelete);
	};
	let statusCode: number = $state(-1);
	let message: string = $state('');
	const handleProposalSubmit = async (event: Event) => {
		event.preventDefault();

		const formData = {
			title,
			startDate,
			endDate,
			maxParticipant,
			format,
			description,
			advisor,
			startTime,
			endTime,
			activityRole
		};

		try {
			const response = await fetch('/api/proposal/submit', {
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
	$effect(() => {
		if (maxParticipant < 0) maxParticipant = 0;
	});
</script>

<div class="flex justify-center overflow-y-scroll">
	<div class="h-fit border-2 border-gray-100 p-5 shadow-md lg:h-auto">
		<h1 class="my-5 text-center text-4xl font-bold">Activity Proposal</h1>
		<form onsubmit={handleProposalSubmit} class="py-4">
			<!-- Left column -->
			<div class="flex flex-col text-sm lg:flex-row">
				<div class="flex-initial">
					<div class="grid grid-cols-[120px_auto] gap-4 pr-2 lg:grid-cols-[30%_70%]">
						<div class="content-center text-right">
							<label for="title" class="font-bold">Title</label>
						</div>
						<div>
							<input
								type="text"
								id="title"
								bind:value={title}
								required
								class="w-56 rounded border border-gray-300 p-2 text-base"
							/>
						</div>
						<div class=" content-center text-right">
							<label for="format" class="font-bold">Format</label>
						</div>
						<div>
							<select
								id="format"
								bind:value={format}
								required
								class="w-56 rounded border border-gray-300 p-2 text-base"
							>
								<option value="" disabled selected>Select format</option>
								<option value="Project">Project</option>
								<option value="Workshop">Workshop</option>
							</select>
						</div>
						{#if format === 'Project'}
							<div class=" content-center text-right">
								<label for="startDate" class="font-bold">Start Date</label>
							</div>
							<div>
								<input
									type="date"
									id="startDate"
									bind:value={startDate}
									required
									class="w-56 rounded border border-gray-300 p-2 text-base"
								/>
							</div>

							<div class=" content-center text-right">
								<label for="endDate" class="font-bold">End Date</label>
							</div>
							<div>
								<input
									type="date"
									id="endDate"
									bind:value={endDate}
									required
									class="w-56 rounded border border-gray-300 p-2 text-base"
								/>
							</div>
							<div class=" content-center text-right">
								<label for="advisor" class="font-bold">Advisor</label>
							</div>
							<div>
								<input
									type="text"
									id="advisor"
									bind:value={advisor}
									required
									class="w-56 rounded border border-gray-300 p-2 text-base"
								/>
							</div>
						{:else if format === 'Workshop'}
							<div class=" content-center text-right">
								<label for="startDate" class="font-bold">Start Date</label>
							</div>
							<div>
								<input
									type="date"
									id="startDate"
									bind:value={startDate}
									required
									class="w-56 rounded border border-gray-300 p-2 text-base"
								/>
							</div>
							<div class=" content-center text-right">
								<label for="startTime" class="font-bold">Start Time</label>
							</div>
							<div>
								<input
									type="time"
									id="startTime"
									bind:value={startTime}
									required
									class="w-56 rounded border border-gray-300 p-2 text-base"
								/>
							</div>
							<div class=" content-center text-right">
								<label for="endDate" class="font-bold">End Date</label>
							</div>
							<div>
								<input
									type="date"
									id="endDate"
									bind:value={endDate}
									required
									class="w-56 rounded border border-gray-300 p-2 text-base"
								/>
							</div>
							<div class=" content-center text-right">
								<label for="endTime" class="font-bold">End Time</label>
							</div>
							<div>
								<input
									type="time"
									id="endTime"
									bind:value={endTime}
									required
									class="w-56 rounded border border-gray-300 p-2 text-base"
								/>
							</div>
						{/if}
						<div class=" content-center text-right">
							<label for="maxParticipant" class="font-bold">Number of participant</label>
						</div>
						<div>
							<input
								type="number"
								id="maxParticipant"
								bind:value={maxParticipant}
								onchange={() => {
									if (maxParticipant < 1) maxParticipant = 1;
								}}
								required
								class="w-28 rounded border border-gray-300 p-2 text-base"
							/>
						</div>
					</div>
				</div>
				<!-- Right column -->
				<div class="flex-1">
					<div class="mt-4 grid flex-1 grid-cols-[120px_auto] gap-4 pr-2 lg:mt-0">
						<div class="mt-4 content-start text-right">
							<label for="description" class="block font-bold">Description</label>
						</div>
						<textarea
							id="description"
							bind:value={description}
							required
							class="max-h-96 w-56 min-w-32 max-w-96 resize rounded border border-gray-300 p-2 text-base"
						></textarea>
						<div class=" align-center content-center text-right">
							<label for="activityRole" class="block font-bold">Activity Role</label>
						</div>
						<div class="align-center">
							<div>
								<div class="flex h-12 flex-row space-x-2 pb-2 pt-2">
									<input
										type="text"
										id="newActivityRole"
										bind:value={newActivityRole}
										class="w-56 rounded border border-gray-300 p-2 text-base"
									/>
									<button
										type="button"
										onclick={addActivityRole}
										class="block cursor-pointer rounded bg-blue-500 px-2 py-1 text-base text-white hover:bg-blue-700"
										>Add</button
									>
								</div>
							</div>
						</div>

						<div></div>
						<div class="flex flex-wrap content-center" style="column-gap: 10px; row-gap: 10px;">
							{#each activityRole as role}
								<button
									type="button"
									onclick={() => deleteActivityRole(role)}
									aria-label="remove"
									class="rounded-lg border-2 border-indigo-500 p-1 text-indigo-500 hover:bg-indigo-500 hover:text-white"
									>{role}
								</button>
							{/each}
						</div>
					</div>
				</div>
			</div>
			<!-- Submit button centered -->
			<div class="mt-4 w-full text-center">
				<button
					type="submit"
					class="w-40 cursor-pointer rounded bg-blue-500 p-2 text-base text-white hover:bg-blue-700"
					>Submit</button
				>
			</div>
		</form>
	</div>
</div>

<Overlay {showOverlay} height='[200px]' width='[300px]'>
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
			<span class="mt-4 text-xl">{message}</span>
		</div>
	</div></Overlay
>
