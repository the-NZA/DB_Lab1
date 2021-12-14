<template>
	<teleport to="body">
		<transition name="slide-fade">
			<div v-if="showModal" class="modal-view">
				<div class="modal-view__wrapper">
					<div class="modal-view__header">
						<button @click="emit('closePressed')">Close</button>
					</div>

					<div class="modal-view__body">
						<slot></slot>
					</div>

					<div class="modal-view__footer">
						<button @click="emit('savePressed')">Save</button>
					</div>
				</div>
			</div>
		</transition>
	</teleport>
</template>

<script lang="ts" setup>
defineProps({
	showModal: {
		type: Boolean,
		required: true,
	},
})

const emit = defineEmits<{
	(e: "savePressed"): void
	(e: "closePressed"): void
}>()

</script>

<style>
.modal-view {
	position: absolute;
	top: 0;
	bottom: 0;
	left: 0;
	right: 0;
	background-color: rgba(var(--black), 0.15);
	box-sizing: border-box;

	z-index: 999;
}

.modal-view__wrapper {
	min-width: calc(var(--site-width) / 1.5);
	min-height: calc(var(--site-width) / 2);
	height: auto;
	position: absolute;
	top: 50%;
	left: 50%;
	transform: translate(-50%, -50%);

	border-radius: var(--offset-half);
	box-sizing: border-box;
	box-shadow: 10px 10px 20px rgba(var(--black), 0.2),
		-20px -20px 20px rgba(var(--leanen), 0.3);

	background-color: rgb(var(--white));
	padding: var(--offset);

	display: grid;
	grid-template-rows: min-content 1fr min-content;
	gap: var(--offset);
}

.modal-view__body {
}

.modal-view__header {
	text-align: right;
}
.modal-view__footer {
}

.slide-fade-enter-active,
.slide-fade-leave-active {
	transition: all 0.15s ease-out;
}

.slide-fade-enter-from,
.slide-fade-leave-to {
	opacity: 0;
}
</style>