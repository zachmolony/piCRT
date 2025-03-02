<script lang="ts">
	import { playCategory, stopPlayback } from '$lib';

	let currentlyPlaying = $state({
		name: 'Ghost in the Shell',
		thumb: '/thumbs/anime.png'
	});

	let categories = $state([
		{ name: 'anime', thumb: '/thumbs/anime.png', count: 43 },
		{ name: 'jdm', thumb: '/thumbs/jdm.png', count: 2 },
		{ name: 'longplays', thumb: '/thumbs/longplays.png', count: 52 },
		{ name: 'misc', thumb: '/thumbs/misc.png', count: 32 },
		{ name: 'test', thumb: '/thumbs/misc.png', count: 12 }
	]);
</script>

<div class="mx-auto mb-5 w-full text-left font-mono text-sm">
	<h1 class="mb-1 ml-3 font-normal">[ 📼 Now Playing ]</h1>

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
	<h1 class="mb-1 ml-3 font-normal">[ 🎥 Categories ]</h1>

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
						<p class="text-xs">[ {category.count} videos ]</p>
					</div></button
				>
			</div>
		{/each}
	</div>
</div>
