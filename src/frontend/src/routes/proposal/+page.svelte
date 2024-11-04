<script lang='ts'>
    let title = '';
    let proposer = '';
    let startDate = '';
    let endDate = '';
    let maxNumber = 0;
    let format = '';
    let description = '';
    let proposeDateTime = '';

    const handleProposalSubmit = async (event: Event) => {
      event.preventDefault();
  
      const formData = {
      title,
      proposer,
      startDate,
      endDate,
      maxNumber,
      format,
      description,
      proposeDateTime
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
        <input type="text" id="format" bind:value={format} required />
    </div>
    <div>
        <label for="startDate">Start Date:</label>
        <input type="text" id="startDate" bind:value={startDate} required />
    </div>
    <div>
        <label for="endDate">End Date:</label>
        <input type="text" id="endDate" bind:value={endDate} required />
    </div>

    {#if format === 'project'}


    {:else if format === 'workshop'}



    <button type="submit">Submit</button>
</form>

<style>
    h1 {
      text-align: center;
      font-weight: bold;
      font-size: xx-large;
      margin: 20px 0;
    }
    form {
      display: flex;
      flex-direction: column;
      width: 300px;
      margin: 0 auto;
    }
  
    div {
      margin-bottom: 16px;
    }
  
    label {
      margin-bottom: 8px;
      font-weight: bold;
    }
  
    input {
      padding: 8px;
      font-size: 16px;
      border: 1px solid #ccc;
      border-radius: 4px;
    }

    textarea {
      padding: 8px;
      font-size: 16px;
      border: 1px solid #ccc;
      border-radius: 4px;
    }
  
    button {
      padding: 10px;
      font-size: 16px;
      background-color: #007BFF;
      color: white;
      border: none;
      border-radius: 4px;
      cursor: pointer;
    }
  
    button:hover {
      background-color: #0056b3;
    }
  </style>