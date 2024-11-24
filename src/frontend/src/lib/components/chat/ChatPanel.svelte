<script>
	let messages = [];
	let newMessage = $state('');
	export let chatid: number;
	export let sendMessage();
	function sendMessageCache() {
		if (newMessage.trim() !== '') {
			messages = [...messages, newMessage];
			newMessage = '';
		}
	}
	function sendMessageCache() {}
</script>

<div class="chat-panel">
	<div class="messages">
		{#each messages as message}
			<div class="message">{message}</div>
		{/each}
	</div>
	<div class="input-panel">
		<input
			type="text"
			bind:value={newMessage}
			on:keydown={(e) => e.key === 'Enter' && sendMessage()}
			placeholder="Type your message..."
		/>
		<button on:click={sendMessage}>Send</button>
	</div>
</div>

<style>
	.chat-panel {
		display: flex;
		flex-direction: column;
		height: 100vh;
		width: 100vw;
	}

	.messages {
		flex: 1;
		overflow-y: auto;
		padding: 10px;
		border-bottom: 1px solid #ccc;
	}

	.message {
		margin-bottom: 10px;
		padding: 10px;
		background-color: #f1f1f1;
		border-radius: 5px;
	}

	.input-panel {
		display: flex;
		padding: 10px;
		border-top: 1px solid #ccc;
	}

	.input-panel input {
		flex: 1;
		padding: 10px;
		border: 1px solid #ccc;
		border-radius: 5px;
		margin-right: 10px;
	}

	.input-panel button {
		padding: 10px 20px;
		border: none;
		background-color: #007bff;
		color: white;
		border-radius: 5px;
		cursor: pointer;
	}

	.input-panel button:hover {
		background-color: #0056b3;
	}
</style>
