<script lang='ts'>
    let title = '';
    let proposer = '';
    let startDate = '';
    let endDate = '';
    let maxNumber = 0;
    let format = '';
    let description = '';
    let advisor = '';
    let startTime = '';
    let endTime = '';
    let activityRole: string[] = [];
    let newActivityRole = '';

    const addActivityRole = () => {
      if (newActivityRole.trim() !== '') {
        activityRole = [...activityRole, newActivityRole.trim()];
        newActivityRole = '';
      }
    };

    const formatDateTime = (date: Date): string => {
      const pad = (num: number) => String(num).padStart(2, '0');
      const year = date.getFullYear();
      const month = pad(date.getMonth() + 1);
      const day = pad(date.getDate());
      const hours = pad(date.getHours());
      const minutes = pad(date.getMinutes());
      const seconds = pad(date.getSeconds());
      return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    };

    const handleProposalSubmit = async (event: Event) => {
      event.preventDefault();

      const proposeDateTime = formatDateTime(new Date());
  
      const formData = {
      title,
      proposer,
      startDate,
      endDate,
      maxNumber,
      format,
      description,
      proposeDateTime,
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

<h1 class="text-center font-bold text-4xl my-5">Activity Proposal</h1>
<form on:submit={handleProposalSubmit} class="flex flex-col w-72 mx-auto">
  <div class="mb-4">
    <label for="title" class="mb-2 font-bold">Title:</label>
    <input type="text" id="title" bind:value={title} required class="p-2 text-lg border border-gray-300 rounded" />
  </div>
  <div class="mb-4">
    <label for="format" class="mb-2 font-bold">Format:</label>
    <select id="format" bind:value={format} required class="p-2 text-lg border border-gray-300 rounded">
      <option value="" disabled selected>Select format</option>
      <option value="project">Project</option>
      <option value="workshop">Workshop</option>
    </select>
  </div>

  {#if format === 'project'}
    <div class="mb-4">
      <label for="startDate" class="mb-2 font-bold">Start Date:</label>
      <input type="text" id="startDate" bind:value={startDate} required class="p-2 text-lg border border-gray-300 rounded" />
    </div>
    <div class="mb-4">
      <label for="endDate" class="mb-2 font-bold">End Date:</label>
      <input type="text" id="endDate" bind:value={endDate} required class="p-2 text-lg border border-gray-300 rounded" />
    </div>
    <div class="mb-4">
      <label for="advisor" class="mb-2 font-bold">Advisor:</label>
      <input type="text" id="advisor" bind:value={advisor} required class="p-2 text-lg border border-gray-300 rounded" />
    </div>
  {:else if format === 'workshop'}
    <div class="mb-4">
      <label for="startDate" class="mb-2 font-bold">Start Date:</label>
      <input type="text" id="startDate" bind:value={startDate} required class="p-2 text-lg border border-gray-300 rounded" />
    </div>
    <div class="mb-4">
      <label for="startTime" class="mb-2 font-bold">Start Time:</label>
      <input type="text" id="startTime" bind:value={startTime} required class="p-2 text-lg border border-gray-300 rounded" />
    </div>
    <div class="mb-4">
      <label for="endDate" class="mb-2 font-bold">End Date:</label>
      <input type="text" id="endDate" bind:value={endDate} required class="p-2 text-lg border border-gray-300 rounded" />
    </div>
    <div class="mb-4">
      <label for="endTime" class="mb-2 font-bold">End Time:</label>
      <input type="text" id="endTime" bind:value={endTime} required class="p-2 text-lg border border-gray-300 rounded" />
    </div>
  {/if}
  
  <div class="mb-4">
    <label for="maxNumber" class="mb-2 font-bold">Number of participant:</label>
    <input type="number" id="maxNumber" bind:value={maxNumber} required class="p-2 text-lg border border-gray-300 rounded" />
  </div>
  <div class="mb-4">
    <label for="description" class="mb-2 font-bold">Description:</label>
    <textarea id="description" bind:value={description} required class="p-2 text-lg border border-gray-300 rounded"></textarea>
  </div>
  <div class="mb-4">
    <label for="activityRole" class="mb-2 font-bold">Activity Role:</label>
    <input type="text" id="newActivityRole" bind:value={newActivityRole} class="p-2 text-lg border border-gray-300 rounded" />
    <button type="button" on:click={addActivityRole} class="p-2 text-lg bg-blue-500 text-white rounded cursor-pointer hover:bg-blue-700 mt-2">Add Role</button>
    <ul class="mt-2">
      {#each activityRole as role}
        <li class="list-disc ml-5">{role}</li>
      {/each}
    </ul>
  </div>

  <button type="submit" class="p-2 text-lg bg-blue-500 text-white rounded cursor-pointer hover:bg-blue-700 mt-4">Submit</button>
</form>