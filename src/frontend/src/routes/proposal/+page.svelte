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

<h1>Activity Proposal</h1>
<form on:submit={handleProposalSubmit}>
    <div>
        <label for="title">Title:</label>
        <input type="text" id="title" bind:value={title} required />
    </div>
    <div>
        <label for="format">Format:</label>
        <select id="format" bind:value={format} required>
          <option value="" disabled selected>Select format</option>
          <option value="project">Project</option>
          <option value="workshop">Workshop</option>
        </select>
    </div>

    {#if format === 'project'}
      <div>
        <label for="startDate">Start Date:</label>
        <input type="text" id="startDate" bind:value={startDate} required />
      </div>
      <div>
          <label for="endDate">End Date:</label>
          <input type="text" id="endDate" bind:value={endDate} required />
      </div>
      <div>
        <label for="advisor">Advisor:</label>
        <input type="text" id="advisor" bind:value={advisor} required />
      </div>

    {:else if format === 'workshop'}
      <div>
        <label for="startDate">Start Date:</label>
        <input type="text" id="startDate" bind:value={startDate} required />
      </div>
      <div>
        <label for="startDate">Start Time:</label>
        <input type="text" id="startDate" bind:value={startDate} required />
      </div>
      <div>
        <label for="endDate">End Date:</label>
        <input type="text" id="endDate" bind:value={endDate} required />
      </div>
      <div>
        <label for="endTime">End Time:</label>
        <input type="text" id="endTime" bind:value={endTime} required />
      </div>
    {/if}
    
    <div>
      <label for="maxNumber">Number of participant:</label>
      <input type="number" id="maxNumber" bind:value={maxNumber} required />
    </div>
    <div>
      <label for="description">Description:</label>
      <textarea id="description" bind:value={description} required></textarea>
    </div>
    <div>
      <label for="activityRole">Activity Role:</label>
      <input type="text" id="newActivityRole" bind:value={newActivityRole} />
      <button type="button" on:click={addActivityRole}>Add Role</button>
      <ul>
          {#each activityRole as role}
              <li>{role}</li>
          {/each}
      </ul>
    </div>

  <button type="submit">Submit</button>
</form>

<style>
  @import '../../styles.css';
</style>