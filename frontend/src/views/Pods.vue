<template>
    <div class="pods-view">
        <pod-list />
    </div>
</template>

<script>
import PodList from '@/components/kubernetes/PodList.vue'

export default {
    name: 'PodsView',
    components: {
        PodList
    }
}
</script>

<style scoped>
.pods-view {
    width: 100%;
    height: 100%;
}
</style>