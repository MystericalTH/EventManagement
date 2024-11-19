<script lang="ts">
	let title = '';
	let startdate = '';
	let enddate = '';
	let maxnumber = 0;
	let format = '';
	let description = '';
	let advisor = '';
	let starttime = '';
	let endtime = '';
	let activityRoles: string[] = [];
	let newActivityRole = '';

	const addActivityRole = () => {
		if (newActivityRole.trim() !== '') {
			activityRoles = [...activityRoles, newActivityRole.trim()];
			newActivityRole = '';
		}
	};

	const handleProposalSubmit = async (event: Event) => {
		event.preventDefault();

		const formData: any = {
			title,
			maxnumber,
			format,
			description,
			activityRoles
		};

		if (format === 'project') {
            formData.startDate = new Date(startdate).toISOString();
            formData.endDate = new Date(enddate).toISOString();
            formData.advisor = advisor;
        } else if (format === 'workshop') {
            formData.startDate = startdate;
            formData.startTime = starttime;
            formData.endDate = enddate;
            formData.endTime = endtime;
        }

		try {
			const response = await fetch('/api/activities', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(formData)
			});

			if (response.ok) {
				console.log('Form submitted successfully');
				resetForm();
			} else {
				console.error('Form submission failed');
			}
		} catch (error) {
			console.error('Error submitting form:', error);
		}
	};

	const resetForm = () => {
        title = '';
        startdate = '';
        enddate = '';
        maxnumber = 0;
        format = '';
        description = '';
        advisor = '';
        starttime = '';
        endtime = '';
        activityRoles = [];
        newActivityRole = '';
    };
</script>

<h1 class="my-5 text-center text-4xl font-bold">Activity Proposal</h1>

<form on:submit={handleProposalSubmit} class="mx-auto flex w-8/12 flex-wrap">
	<!-- Left column -->
	<div class="mt-4 w-full pr-2 md:w-1/2">
		<div class="mb-4">
			<label for="title" class="mb-2 font-bold">Title:</label>
			<input
				type="text"
				id="title"
				bind:value={title}
				required
				class="w-56 rounded border border-gray-300 p-2 text-lg"
			/>
		</div>
		<div class="mb-4">
			<label for="format" class="mb-2 font-bold">Format:</label>
			<select
				id="format"
				bind:value={format}
				required
				class="w-56 rounded border border-gray-300 p-2 text-lg"
			>
				<option value="" disabled selected>Select format</option>
				<option value="project">Project</option>
				<option value="workshop">Workshop</option>
			</select>
		</div>
		{#if format === 'project'}
			<div class="mb-4">
				<label for="startDate" class="mb-2 font-bold">Start Date:</label>
				<input
                    type="date"
                    id="startDate"
                    bind:value={startdate}
                    required
                    class="w-56 rounded border border-gray-300 p-2 text-lg"
                />
			</div>
			<div class="mb-4">
				<label for="endDate" class="mb-2 font-bold">End Date:</label>
				<input
                    type="date"
                    id="endDate"
                    bind:value={enddate}
                    required
                    class="w-56 rounded border border-gray-300 p-2 text-lg"
                />
			</div>
			<div class="mb-4">
				<label for="advisor" class="mb-2 font-bold">Advisor:</label>
				<input
					type="text"
					id="advisor"
					bind:value={advisor}
					required
					class="w-56 rounded border border-gray-300 p-2 text-lg"
				/>
			</div>
		{:else if format === 'workshop'}
			<div class="mb-4">
				<label for="startDate" class="mb-2 font-bold">Start Date:</label>
				<input
                    type="date"
                    id="startDate"
                    bind:value={startdate}
                    required
                    class="w-56 rounded border border-gray-300 p-2 text-lg"
                />
			</div>
			<div class="mb-4">
				<label for="startTime" class="mb-2 font-bold">Start Time:</label>
				<input
                    type="time"
                    id="startTime"
                    bind:value={starttime}
                    required
                    class="w-56 rounded border border-gray-300 p-2 text-lg"
                />
			</div>
			<div class="mb-4">
				<label for="endDate" class="mb-2 font-bold">End Date:</label>
				<input
                    type="date"
                    id="endDate"
                    bind:value={enddate}
                    required
                    class="w-56 rounded border border-gray-300 p-2 text-lg"
                />
			</div>
			<div class="mb-4">
				<label for="endTime" class="mb-2 font-bold">End Time:</label>
				<input
                    type="time"
                    id="endTime"
                    bind:value={endtime}
                    required
                    class="w-56 rounded border border-gray-300 p-2 text-lg"
                />
			</div>
		{/if}
		<div class="mb-4">
			<label for="maxNumber" class="mb-2 font-bold">Number of participant:</label>
			<input
				type="number"
				id="maxNumber"
				bind:value={maxnumber}
				required
				class="w-28 rounded border border-gray-300 p-2 text-lg"
			/>
		</div>
	</div>
	<!-- Right column -->
	<div class="w-full pl-2 md:w-1/2">
		<div class="mb-4">
			<label for="description" class="mb-2 block font-bold">Description:</label>
			<textarea
				id="description"
				bind:value={description}
				required
				class="w-56 rounded border border-gray-300 p-2 text-lg"
			></textarea>
		</div>
		<div class="mb-4">
			<label for="activityRole" class="mb-2 block font-bold">Activity Role:</label>
			<input
				type="text"
				id="newActivityRole"
				bind:value={newActivityRole}
				class="w-56 rounded border border-gray-300 p-2 text-lg"
			/>
			<button
				type="button"
				on:click={addActivityRole}
				class="mt-2 block cursor-pointer rounded bg-blue-500 p-2 text-lg text-white hover:bg-blue-700"
				>Add Role</button
			>
			<ul class="mt-2">
				{#each activityRoles as role}
					<li class="ml-5 list-disc">{role}</li>
				{/each}
			</ul>
		</div>
	</div>
	<!-- Submit button centered -->
	<div class="mt-4 w-full text-center">
		<button
			type="submit"
			class="w-72 cursor-pointer rounded bg-blue-500 p-2 text-lg text-white hover:bg-blue-700"
			>Submit</button
		>
	</div>
</form>