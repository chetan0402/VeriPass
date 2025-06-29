<script lang="ts">
	import type { Admin } from '$lib/gen/veripass/v1/admin_pb';
	let admin = $props<{ admin: Admin }>();
	console.log(admin.name);
</script>

<div></div>
