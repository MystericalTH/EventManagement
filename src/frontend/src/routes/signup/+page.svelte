<script lang="ts">
    let name = '';
    let email = '';
    let phone = '';
    let githubUrl = '';
    let interest = '';
    let reason = '';
  
    const handleMemberSubmit = async (event: Event) => {
      event.preventDefault();
  
      const formData = {
      name,
      email,
      phone,
      githubUrl,
      interest,
      reason
    };

    try {
      const response = await fetch('/api/member/submit', {
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
  
  <h1>Sign Up</h1>
  <form on:submit={handleMemberSubmit}>
    <div>
      <label for="name">Name:</label>
      <input type="text" id="name" bind:value={name} required />
    </div>
    <div>
      <label for="email">Email:</label>
      <input type="email" id="email" bind:value={email} required />
    </div>
    <div>
      <label for="phone">Phone:</label>
      <input type="tel" id="phone" bind:value={phone} required />
    </div>
    <div>
      <label for="githubUrl">GitHub:</label>
      <input type="url" id="githubUrl" bind:value={githubUrl} required />
    </div>
    <div>
        <label for="interest">What is your interest:</label>
        <textarea type="url" id="interest" bind:value={interest} required />
    </div>
    <div>
        <label for="reason">Why do you want to join the club:</label>
        <textarea type="url" id="reason" bind:value={reason} required />
    </div>
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
