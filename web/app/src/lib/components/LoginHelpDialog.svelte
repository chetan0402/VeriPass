<script lang="ts">
	import { CloseOutline } from 'flowbite-svelte-icons';
	let { close } = $props<{ close: () => void }>();
</script>

<div
	class="relative m-10 flex h-[max(200px,70vh)] w-[clamp(100px,80vw,600px)] flex-col items-center justify-around rounded-2xl border-3 border-[#5555C2] bg-white p-5"
>
	<div class="flex flex-row items-center justify-center">
		<h1 class="text-secondary text-2xl font-bold">Login Help</h1>
	</div>
	<div class="bg-primary-100 mt-2 space-y-4 overflow-y-auto rounded-md p-4 text-sm text-gray-700">
		<div>
			<p class="font-semibold">Q. I don’t have my Institute mail ID.</p>
			<p>
				A. Please check the email ID provided by you during admission. If you haven't received any
				email about your institute ID, contact the Central Computing Facility (CCF), MANIT.
			</p>
		</div>
		<div>
			<p class="font-semibold">Q. I forgot the password of my email ID.</p>
			<p>
				A. Try resetting your password using your phone number. If that doesn’t work, contact CCF,
				MANIT for assistance.
			</p>
		</div>
		<div>
			<p class="font-semibold">Q. I can’t find my account in the Google sign-in list.</p>
			<p>A. Make sure your student Gmail account is logged in on the browser you are using.</p>
		</div>
	</div>
	<p class="text-secondary-700 mt-2 text-sm">Issue not listed here?</p>
	<button class="bg-secondary w-full rounded-full p-3 text-white">Send Mail to Admin</button>
	<button onclick={close} class="text-secondary absolute top-5 right-5">
		<CloseOutline />
	</button>
</div>
