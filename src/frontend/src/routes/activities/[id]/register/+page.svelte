<script lang="ts">
    import type { ActivityData } from "$lib/types/activity";
    export let data: { activity: ActivityData};

    let { activity } = data;
    let expectation = '';

    const handleRegisterSubmit = async (event: Event) => {
      event.preventDefault();
  
      const formData = {
      id: activity.id,
      expectation
    };

    try {
      const response = await fetch('/api/register/submit', {
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

<h1 class="text-2xl font-semibold my-10">{data.activity.title}</h1>

<form on:submit={handleRegisterSubmit} class="flex flex-col w-72 mx-auto">
    <label for="expectation" class="mb-2 mt-14 font-bold block">What do you expect from this activity?:</label>
    <textarea id="expectation" name="expectation" bind:value={expectation} required class="p-2 text-lg border border-gray-300 rounded h-60 mb-14"></textarea>
    <button type="submit" class="p-2 text-lg bg-blue-500 text-white rounded cursor-pointer hover:bg-blue-700 mt-4">Submit</button>
</form>