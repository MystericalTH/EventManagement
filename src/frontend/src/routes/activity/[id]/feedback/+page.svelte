<script lang="ts">
    export let data: { id: string };

    let id = data.id;
    let feedback = '';

    const handleFeedbackSubmit = async (event: Event) => {
      event.preventDefault();
  
      const formData = {
      id,
      feedback
    };

    try {
      const response = await fetch('/api/feedback/submit', {
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

<h1 class="text-center font-bold text-4xl my-5">Activity's Feedback</h1>

<form on:submit={handleFeedbackSubmit} class="flex flex-col w-72 mx-auto">
    <label for="feedback" class="mb-2 mt-14 font-bold block">Feedback:</label>
    <textarea id="feedback" name="feedback" bind:value={feedback} required class="p-2 text-lg border border-gray-300 rounded h-60 mb-14"></textarea>
    <button type="submit" class="p-2 text-lg bg-blue-500 text-white rounded cursor-pointer hover:bg-blue-700 mt-4">Submit</button>
</form>