<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import jsQR from 'jsqr';

	let isScanning = $state(false);
	let cameraStream: MediaStream | undefined;

	let {
		options
	}: {
		options: {
			onPermissionError: () => void;
			onResulted: (data: string) => void;
			onUpdateStatus: (update: string) => void;
		};
	} = $props();

	let video: HTMLVideoElement;
	let canvas: HTMLCanvasElement;

	onMount(() => {
		requestCamera();
	});

	function requestCamera() {
		navigator.mediaDevices
			.getUserMedia({
				audio: false,
				video: {
					facingMode: 'environment'
				}
			})
			.then(async (userStream) => {
				video.srcObject = userStream;
				video.setAttribute('playsinline', 'true');
				await video.play();
				isScanning = true;
				cameraStream = userStream;
			})
			.catch(() => {
				options.onPermissionError();
			});
	}

	function startScan() {
		if (!isScanning) {
			return;
		}
		const context = canvas.getContext('2d', { willReadFrequently: true });
		const { width, height } = canvas;

		context?.drawImage(video, 0, 0, width, height);

		const imageData = context?.getImageData(0, 0, width, height);
		if (imageData) {
			const qrCode = jsQR(imageData.data, width, height);

			if (qrCode) {
				options.onResulted(qrCode.data);
				isScanning = false;
			} else {
				options.onUpdateStatus('Scanning for qr code');
				setTimeout(startScan, 500);
			}
		}
	}

	function stopCamera() {
		isScanning = false;
		video.srcObject = null;
		if (cameraStream) {
			cameraStream.getTracks().forEach((t) => {
				t.stop();
			});
			cameraStream = undefined;
			console.log('closed camera');
		}
	}

	onDestroy(() => {
		stopCamera();
	});

	function onCanPlay() {
		canvas.width = video.videoWidth;
		canvas.height = video.videoHeight;
		isScanning = true;
		startScan();
	}
</script>

<div class="relative h-80 w-80 overflow-clip rounded-sm border-8 border-solid border-red-200">
	<video
		class="absolute inset-0 h-full w-full object-cover"
		oncanplay={onCanPlay}
		bind:this={video}
	>
		<track kind="captions" src="" />
	</video>

	<canvas class="absolute inset-0 hidden h-full w-full object-cover" bind:this={canvas}></canvas>
</div>
