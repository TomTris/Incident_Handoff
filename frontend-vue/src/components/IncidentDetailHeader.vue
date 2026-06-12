<script setup lang="ts">
import type { Incident, UserContext } from '@/types';
import { computed } from 'vue';

const props = defineProps<{
    inc: Incident,
    myIdentity: UserContext,
}>()

const emit = defineEmits<{
    updateSeverity: [payload: { key: string, value: string }]
    updateStatus: [payload: { key: string, value: string }]
    updateOnCall: [payload: { key: string, value: string }]
}>()

function onSelectSeverity(e: Event) {
    emit('updateSeverity', { key: "severity", value: (e.target as HTMLSelectElement).value })
}
function onSelectStatus(e: Event) {
    emit('updateStatus', { key: "status", value: (e.target as HTMLSelectElement).value })
}
function onChangeOnCall(e: KeyboardEvent) {
    emit('updateOnCall', { key: "on_call", value: (e.target as HTMLInputElement).value })
}

const canEdit = computed(() =>
    props.myIdentity.role === "admin" || props.myIdentity.username === props.inc.on_call
)
</script>

<template>
    <h1>{{ inc.title }}</h1>

    <select :value="inc.severity" :disabled="!canEdit" @change="onSelectSeverity">
        <option>SEV1</option>
        <option>SEV2</option>
        <option>SEV3</option>
    </select>

    <select :value="inc.status" :disabled="!canEdit" @change="onSelectStatus">
        <option value="triggered">Triggered</option>
        <option value="acknowledged">Acknowledged</option>
        <option value="investigating">Investigating</option>
        <option value="mitigated">Mitigated</option>
        <option value="resolved">Resolved</option>
    </select>

    <input :value="inc.on_call" :disabled="!canEdit" @keyup.enter="onChangeOnCall" />

    <h3>{{ inc.service }}</h3>
    <p>{{ inc.id }}</p>
    <p>{{ inc.opened_by }}</p>
    <p>{{ inc.created_at }}</p>
</template>