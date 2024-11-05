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
      <textarea id="interest" bind:value={interest} required></textarea>
  </div>
  <div>
      <label for="reason">Why do you want to join the club:</label>
      <textarea id="reason" bind:value={reason} required></textarea>
  </div>
  <button type="submit">Submit</button>
</form>

<style>
  @import '../../styles.css';
</style>
