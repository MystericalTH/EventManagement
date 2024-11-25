<script lang="ts">
	import { onMount } from 'svelte';

	let logs = [];

	onMount(() => {
		logs = [
			{
				timestamp: '2023-10-01T12:00:00Z',
				service: 'Server',
				level: 'INFO',
				message: 'Server started on port 8080'
			},
			{
				timestamp: '2023-10-01T12:01:00Z',
				service: 'Frontend',
				level: 'INFO',
				message: 'Frontend service initialized'
			},
			{
				timestamp: '2023-10-01T12:02:00Z',
				service: 'Nginx',
				level: 'INFO',
				message: 'Nginx proxy started'
			},
			{
				timestamp: '2023-10-01T12:03:00Z',
				service: 'MySQL',
				level: 'INFO',
				message: 'MySQL database connected'
			},
			{
				timestamp: '2023-10-01T12:04:00Z',
				service: 'Server',
				level: 'ERROR',
				message: 'Failed to connect to database'
			},
			{
				timestamp: '2023-10-01T12:05:00Z',
				service: 'Frontend',
				level: 'WARN',
				message: 'Missing translation key for "welcome"'
			},
			{
				timestamp: '2023-10-01T12:06:00Z',
				service: 'Nginx',
				level: 'ERROR',
				message: '502 Bad Gateway'
			},
			{
				timestamp: '2023-10-01T12:07:00Z',
				service: 'MySQL',
				level: 'DEBUG',
				message: 'Query executed: SELECT * FROM users'
			},
			{
				timestamp: '2023-10-01T12:08:00Z',
				service: 'Server',
				level: 'INFO',
				message: 'New user registered'
			},
			{
				timestamp: '2023-10-01T12:09:00Z',
				service: 'Frontend',
				level: 'INFO',
				message: 'User logged in'
			},
			{
				timestamp: '2023-10-01T12:10:00Z',
				service: 'Nginx',
				level: 'INFO',
				message: 'Request forwarded to backend'
			},
			{
				timestamp: '2023-10-01T12:11:00Z',
				service: 'MySQL',
				level: 'INFO',
				message: 'New record inserted into database'
			},
			{
				timestamp: '2023-10-01T12:12:00Z',
				service: 'Server',
				level: 'ERROR',
				message: 'Unhandled exception occurred'
			},
			{
				timestamp: '2023-10-01T12:13:00Z',
				service: 'Frontend',
				level: 'WARN',
				message: 'Deprecated API usage'
			},
			{
				timestamp: '2023-10-01T12:14:00Z',
				service: 'Nginx',
				level: 'ERROR',
				message: 'Timeout while connecting to backend'
			},
			{
				timestamp: '2023-10-01T12:15:00Z',
				service: 'MySQL',
				level: 'DEBUG',
				message: 'Transaction started'
			},
			{
				timestamp: '2023-10-01T12:16:00Z',
				service: 'Server',
				level: 'INFO',
				message: 'Scheduled task executed'
			},
			{
				timestamp: '2023-10-01T12:17:00Z',
				service: 'Frontend',
				level: 'INFO',
				message: 'Page loaded successfully'
			},
			{
				timestamp: '2023-10-01T12:18:00Z',
				service: 'Nginx',
				level: 'INFO',
				message: 'SSL certificate renewed'
			},
			{
				timestamp: '2023-10-01T12:19:00Z',
				service: 'MySQL',
				level: 'INFO',
				message: 'Database backup completed'
			}
		];
	});
</script>

<div class="overflow-y-auto">
	<h1 class="mb-4 text-3xl font-bold">Logging</h1>
	<h2 class="mb-3 text-2xl font-semibold">Logging Table</h2>
	<div class="h-96 overflow-y-auto border border-gray-200">
		<table class="min-w-full border border-gray-200 bg-white">
			<thead class="sticky top-0 bg-gray-300">
				<tr>
					<th class="border-b px-4 py-2">Timestamp</th>
					<th class="border-b px-4 py-2">Service</th>
					<th class="border-b px-4 py-2">Level</th>
					<th class="border-b px-4 py-2">Message</th>
				</tr>
			</thead>
			<tbody>
				{#each logs as log}
					<tr>
						<td class="border-b px-4 py-2">{log.timestamp}</td>
						<td class="border-b px-4 py-2">{log.service}</td>
						<td class="border-b px-4 py-2">{log.level}</td>
						<td class="border-b px-4 py-2">{log.message}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>

	<p class="mb-6">Developers can view the logging information of the endpoint on this page.</p>

	<h2 class="mb-3 text-2xl font-semibold">Docker Services Logging Information</h2>
	<p class="mb-4">
		This page provides detailed logging information for the following Docker services:
	</p>
	<ul class="mb-6 list-inside list-disc">
		<li class="mb-2">
			<em>Server:</em> Logs related to the backend server operations, including request handling, error
			messages, and server status updates. This includes logs from the GIN framework in Golang, capturing
			HTTP request handling, middleware processing, and routing information.
		</li>
		<li class="mb-2">
			<em>Frontend:</em> Logs from the frontend service, capturing user interactions, component rendering,
			and client-side errors.
		</li>
		<li class="mb-2">
			<em>Nginx:</em> Logs from the Nginx reverse proxy, including access logs, error logs, and proxy
			request handling.
		</li>
		<li class="mb-2">
			<em>MySQL:</em> Database logs, including query execution, connection status, and error messages.
		</li>
	</ul>

	<h2 class="mb-3 text-2xl font-semibold">Interpreting the Data</h2>
	<p class="mb-4">
		The logs are structured to help developers quickly identify and troubleshoot issues:
	</p>
	<ul class="mb-6 list-inside list-disc">
		<li class="mb-2">
			<em>Timestamps:</em> Each log entry includes a timestamp to indicate when the event occurred.
		</li>
		<li class="mb-2">
			<em>Log Levels:</em> Logs are categorized by severity levels such as INFO, WARN, ERROR, and DEBUG
			to prioritize attention.
		</li>
		<li class="mb-2">
			<em>Service Identifiers:</em> Each log entry is tagged with the service name (e.g., Server, Frontend,
			Nginx, MySQL) to easily filter logs by service.
		</li>
		<li class="mb-2">
			<em>Message Content:</em> Detailed messages provide context about the event, including error descriptions,
			request details, and status updates.
		</li>
	</ul>
	<p class="mb-6">
		Use the filtering and search functionalities to narrow down logs to specific time frames,
		services, or keywords to efficiently diagnose and resolve issues.
	</p>
</div>
