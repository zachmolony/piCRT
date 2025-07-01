<script lang="ts">
	import { playCategory, stopPlayback } from '$lib';
	import { onMount } from 'svelte';

	let currentlyPlaying = $state({
		name: 'Akira (1988)',
		thumb: '/thumbs/anime.png'
	});

	let categories = $state([]);

	async function getCategories() {
		const res = await fetch('http://localhost:5000/categories');
		const data = await res.json();

		console.log(data);

		// If your backend just returns an array of names, map them to thumbs here:
		categories = data.map((name: string) => ({
			name,
			thumb: `/thumbs/${name}.png`
			// You can add count if you want, or leave it out for now
		}));

		console.log(categories);
	}

	onMount(getCategories);
</script>

<div class="mx-8 mb-5 text-left font-mono text-sm">
	<h1 class="mb-1 ml-3">[ 📼 Now Playing ]</h1>

	<div class="mx-auto flex flex-col">
		<div class="relative">
			<pre>+---------------------------------+</pre>
			<pre>|                                 |</pre>
			<pre>|                                 |</pre>
			<pre>+---------------------------------+</pre>

			<button class="absolute left-0 top-0 h-full px-4" onclick={() => stopPlayback()}>
				<div class="flex items-center gap-4">
					<!-- <img class="aspect-square h-8" src={currentlyPlaying.thumb} alt="" /> -->

					<p>{currentlyPlaying.name}</p>
					<p>[ STOP ]</p>
				</div></button
			>
		</div>
	</div>
</div>

<div class="mx-8 text-left font-mono text-sm">
	<h1 class="mb-1 ml-3">[ 🎥 Categories ]</h1>

	<div class="flex flex-col">
		{#each categories as category, i}
			<div class="relative">
				{#if i === 0}
					<pre>+---------------------------------+</pre>
				{/if}
				<pre>|                                 |</pre>
				<pre>|                                 |</pre>
				<pre>+---------------------------------+</pre>

				<button
					class={'absolute left-0 top-0 h-full w-full px-4' + (i !== 0 ? ' pb-4.5' : '')}
					onclick={() => playCategory(category.name)}
				>
					<div class="flex items-center gap-4">
						<img class="aspect-square h-8" src={category.thumb} alt="" />

						<p>{category.name}</p>
						<!-- <p class="text-xs">[ {category.count} videos ]</p> -->
					</div></button
				>
			</div>
		{/each}
	</div>
</div>
