<script lang="ts">
	import type { Snippet } from 'svelte';
	let { className, classNameParent, children, roundParent, roundBox } = $props<{
		className: string;
		classNameParent: string;
		children: Snippet;
		roundParent: string;
		roundBox: string;
	}>();
</script>

<div class={`rotating-bg rotating-vg p-[2px] rounded-${roundParent} ${classNameParent}`}>
	<div class={className + ` rounded-${roundBox}`}>
		{@render children()}
	</div>
</div>
