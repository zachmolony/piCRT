<script lang="ts">
	import { onMount } from 'svelte';

	let view: 'categories' | 'videos' = 'categories';
	let categories: any[] = [];
	let videos: string[] = [];
	let selectedCategory: string | null = null;
	let loadingCategories = false;
	let loadingVideos = false;
	let nowPlaying = '';

	onMount(() => {
		loadCategories();
	});

	async function loadCategories() {
		loadingCategories = true;
		try {
			const res = await fetch('/categories');
			let data = await res.json();
			if (!Array.isArray(data)) data = [];
			categories = data;
		} catch (e) {
			categories = [];
		}
		loadingCategories = false;
	}

	async function openCategory(category: string) {
		selectedCategory = category;
		loadingVideos = true;
		let data = [];
		try {
			const res = await fetch(`/videos/${encodeURIComponent(category)}`);
			data = await res.json();
			if (!Array.isArray(data)) data = [];
			videos = data;
		} catch (e) {
			videos = [];
		}
		view = 'videos';
		loadingVideos = false;
	}

	async function playVideo(category: string, video: string) {
		await fetch(`/play/${category}/${encodeURIComponent(video)}`, { method: 'POST' });
		nowPlaying = `${category}/${video}`;
	}

	async function shuffleCategory(category: string) {
		await fetch(`/play/${category}`, { method: 'POST' });
		nowPlaying = `${category} (shuffled)`;
	}

	async function shuffleAllCategories() {
		for (const category of categories) {
			await shuffleCategory(category.name || category);
		}
	}

	function backToCategories() {
		view = 'categories';
		selectedCategory = null;
		videos = [];
	}
</script>

<div class="mx-6 mb-5 text-left font-mono text-sm">
	<h1 class="mb-1 ml-3">[ 📼 Now Playing ]</h1>
	<div class="mx-auto flex flex-col">
		<div class="relative">
			<pre>+-----------------------------------+</pre>
			<pre>| {nowPlaying || 'Nothing playing'}                   |</pre>
			<pre>|                                   |</pre>
			<pre>+-----------------------------------+</pre>
		</div>
	</div>
</div>

<div class="mx-6 max-w-xs text-left font-mono text-sm">
	{#if view === 'categories'}
		<h1 class="mb-1 ml-3">[ 🎥 Categories ]</h1>
		<button class="mb-2 cursor-pointer px-1 py-1 text-xs" on:click={shuffleAllCategories}
			>{'>'} SHUFFLE ALL {'<'}</button
		>
		{#if loadingCategories}
			<p>Loading categories...</p>
		{:else}
			<div class="flex flex-col">
				{#each categories as category, i}
					<div class="relative mb-2">
						<pre>+------------------------------------+</pre>
						<pre>|                                    |</pre>
						<pre>|                                    |</pre>
						<pre>+------------------------------------+</pre>
						<div class="absolute left-0 top-0 flex h-full w-full items-center gap-4 px-4">
							<img class="aspect-square h-8" src={`/thumbs/${category}.png`} alt="" />
							<p class="flex-1 cursor-pointer" on:click={() => openCategory(category)}>
								{category}
							</p>
							<button
								class="cursor-pointer px-1 py-1 text-xs"
								on:click={() => shuffleCategory(category)}>{'>'} SHUFFLE {'<'}</button
							>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	{:else if view === 'videos'}
		<h1 class="mb-1 ml-3">[ 📂 {selectedCategory} ]</h1>
		<button class="mb-2 cursor-pointer px-1 py-1 text-xs" on:click={backToCategories}
			>{'>'} BACK {'<'}</button
		>
		{#if loadingVideos}
			<p>Loading videos...</p>
		{:else if videos.length === 0}
			<p>No videos found in this category.</p>
		{:else}
			<div class="flex flex-col">
				{#each videos as video, i}
					<div class="relative mb-2">
						<pre>+---------------------------------+</pre>
						<pre>|                                 |</pre>
						<pre>|                                 |</pre>
						<pre>+---------------------------------+</pre>
						<div class="absolute left-0 top-0 flex h-full w-full items-center gap-4 px-4">
							<p
								class="flex-1 cursor-pointer"
								on:click={() => selectedCategory && playVideo(selectedCategory, video)}
							>
								{video}
							</p>
						</div>
					</div>
				{/each}
			</div>
		{/if}
	{/if}
</div>
