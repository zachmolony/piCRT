<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

<div class="mx-auto mt-12 max-w-2xl px-4 font-bold">
	<pre class="mx-auto my-4 text-sm tracking-tighter">
██████╗ ██╗       ██████╗██████╗ ████████╗
██╔══██╗██║      ██╔════╝██╔══██╗╚══██╔══╝
██████╔╝██║█████╗██║     ██████╔╝   ██║   
██╔═══╝ ██║╚════╝██║     ██╔══██╗   ██║   
██║     ██║      ╚██████╗██║  ██║   ██║   
╚═╝     ╚═╝       ╚═════╝╚═╝  ╚═╝   ╚═╝   
  </pre>

	<div class="mb-8">piCRT: Media Terminal</div>

	<div class="mx-auto">
		{@render children()}
	</div>
</div>

<style>
	:global(body) {
		background-color: black;
		color: #00ff00;
		font-family: monospace;
		text-align: center;
	}

	pre {
		line-height: 1.1;
		white-space: pre;
	}
</style>
