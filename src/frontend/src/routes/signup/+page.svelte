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

<h1 class="text-center font-bold text-4xl my-5">Sign Up</h1>
<form on:submit={handleMemberSubmit} class="flex flex-col w-72 mx-auto">
  <div class="mb-4">
    <label for="name" class="mb-2 font-bold">Name:</label>
    <input type="text" id="name" bind:value={name} required class="p-2 text-lg border border-gray-300 rounded" />
  </div>
  <div class="mb-4">
    <label for="email" class="mb-2 font-bold">Email:</label>
    <input type="email" id="email" bind:value={email} required class="p-2 text-lg border border-gray-300 rounded" />
  </div>
  <div class="mb-4">
    <label for="phone" class="mb-2 font-bold">Phone:</label>
    <input type="tel" id="phone" bind:value={phone} required class="p-2 text-lg border border-gray-300 rounded" />
  </div>
  <div class="mb-4">
    <label for="githubUrl" class="mb-2 font-bold">GitHub:</label>
    <input type="url" id="githubUrl" bind:value={githubUrl} required class="p-2 text-lg border border-gray-300 rounded" />
  </div>
  <div class="mb-4">
    <label for="interest" class="mb-2 font-bold">What is your interest:</label>
    <textarea id="interest" bind:value={interest} required class="p-2 text-lg border border-gray-300 rounded"></textarea>
  </div>
  <div class="mb-4">
    <label for="reason" class="mb-2 font-bold">Why do you want to join the club:</label>
    <textarea id="reason" bind:value={reason} required class="p-2 text-lg border border-gray-300 rounded"></textarea>
  </div>
  <button type="submit" class="p-2 text-lg bg-blue-500 text-white rounded cursor-pointer hover:bg-blue-700 mt-4">Submit</button>
</form>