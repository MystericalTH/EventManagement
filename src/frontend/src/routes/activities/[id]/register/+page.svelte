<script lang="ts">
    import { formatDateTime } from "$lib/utils/dateTime";
    import type { ActivityData } from "$lib/types/activity";
    
    export let data: { activity: ActivityData; activityRoles: string[] };

    let { activity, activityRoles } = data;
    let expectation = '';
    let selectedRole = '';

    const handleRegisterSubmit = async (event: Event) => {
      event.preventDefault();

      const registerDateTime = formatDateTime();
  
      const formData = {
      activityID: activity.id,
      role: selectedRole,
      expectation,
      datetime: registerDateTime
    };

    try {
      const response = await fetch('/api/registration/submit', {
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

<h1 class="text-2xl font-semibold my-10 text-center">{data.activity.title}</h1>

<form on:submit={handleRegisterSubmit} class="flex flex-col w-72 mx-auto">
  <label for="role" class="mb-2 font-bold">Select Role:</label>
  <select
    id="role"
    bind:value={selectedRole}
    required
    class="p-2 text-lg border border-gray-300 rounded mb-4"
  >
    <option value="" disabled selected>Select a role</option>
    {#each activityRoles as role}
      <option value={role}>{role}</option>
    {/each}
  </select>

  <label for="expectation" class="mb-2 font-bold block">What do you expect from this activity?</label>
  <textarea
    id="expectation"
    bind:value={expectation}
    required
    class="p-2 text-lg border border-gray-300 rounded h-60 mb-4"
  ></textarea>

  <button
    type="submit"
    class="p-2 text-lg bg-blue-500 text-white rounded cursor-pointer hover:bg-blue-700 mt-4"
  >
    Submit
  </button>
</form>