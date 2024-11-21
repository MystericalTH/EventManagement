<script lang="ts">
	let title = $state('');
	let startDate = $state('');
	let endDate = $state('');
	let maxNumber = $state(1);
	let format = $state('');
	let description = $state('');
	let advisor = $state('');
	let startTime = $state('');
	let endTime = $state('');
	let activityRole: string[] = $state([]);
	let newActivityRole = $state('');

	const addActivityRole = () => {
		if (newActivityRole.trim() !== '' && !activityRole.includes(newActivityRole.trim())) {
			activityRole = [...activityRole, newActivityRole.trim()];
			newActivityRole = '';
		}
	};
	const deleteActivityRole = (roleToDelete: string) => {
		activityRole = activityRole.filter((role) => role !== roleToDelete);
	};

	const handleProposalSubmit = async (event: Event) => {
		event.preventDefault();

		const formData = {
			title: title,
			startdate: startDate,
			enddate: endDate,
			maxnumber: maxNumber,
			format: format,
			description: description,
			advisor: advisor,
			starttime: startTime,
			endtime: endTime,
			activityrole: activityRole
		};

		try {
			const response = await fetch('/api/proposal/submit', {
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
	$effect(() => {
		if (maxNumber < 0) maxNumber = 0;
	});
</script>

<div class="min-w-[900px] border-2 border-gray-100 p-5 shadow-md">
	<h1 class="my-5 text-center text-4xl font-bold">Activity Proposal</h1>
	<form onsubmit={handleProposalSubmit} class="py-4">
		<!-- Left column -->
		<div class="flex flex-row text-sm">
			<div class="flex-initial">
				<div class="grid grid-cols-[30%_70%] gap-4 pr-2">
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
							<option value="project">Project</option>
							<option value="workshop">Workshop</option>
						</select>
					</div>
					{#if format === 'project'}
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
					{:else if format === 'workshop'}
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
						<label for="maxNumber" class="font-bold">Number of participant</label>
					</div>
					<div>
						<input
							type="number"
							id="maxNumber"
							bind:value={maxNumber}
							onchange={() => {
								if (maxNumber < 1) maxNumber = 1;
							}}
							required
							class="w-28 rounded border border-gray-300 p-2 text-base"
						/>
					</div>
				</div>
			</div>
			<!-- Right column -->
			<div class="flex-1">
				<div class="grid flex-1 grid-cols-[120px_auto] gap-4 pr-2">
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
					<div class="mt-2 flex flex-wrap content-center" style="column-gap: 10px; row-gap: 10px;">
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
