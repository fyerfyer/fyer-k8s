<template>
    <div class="namespace-selector">
        <el-select 
            v-model="selectedNamespace" 
            placeholder="Select Namespace"
            clearable
            size="small"
            @change="handleNamespaceChange"
            :loading="loading"
        >
            <el-option label="All Namespaces" value="" />
            <el-option 
                v-for="ns in namespaces" 
                :key="ns" 
                :label="ns" 
                :value="ns" 
            />
        </el-select>
    </div>
</template>

<script>
import { ref, watch, onMounted } from 'vue';
import kubernetesApi from '@/api/kubernetes';

export default {
    name: 'NamespaceSelector',
    props: {
        // Allow the parent component to pass in the initially selected namespace
        modelValue: {
            type: String,
            default: ''
        }
    },
    emits: ['update:modelValue', 'change'],
    setup(props, { emit }) {
        const selectedNamespace = ref(props.modelValue);
        const namespaces = ref(['default', 'kube-system', 'kube-public']);
        const loading = ref(false);

        // Watch for changes in the value passed from the parent component
        watch(() => props.modelValue, (newValue) => {
            selectedNamespace.value = newValue;
        });

        // Trigger an event when the local selected value changes
        watch(selectedNamespace, (newValue) => {
            emit('update:modelValue', newValue);
        });

        // Handle namespace change
        const handleNamespaceChange = () => {
            emit('change', selectedNamespace.value);
        };

        // Extract the list of namespaces from existing Pods
        const extractNamespacesFromPods = async () => {
            loading.value = true;
            try {
                const response = await kubernetesApi.listPods();
                if (response && response.pods && response.pods.length > 0) {
                    // Extract and deduplicate namespaces
                    const uniqueNamespaces = [...new Set(response.pods.map(pod => pod.namespace))];
                    if (uniqueNamespaces.length > 0) {
                        namespaces.value = uniqueNamespaces;
                    }
                }
            } catch (error) {
                console.error('Failed to load namespaces:', error);
            } finally {
                loading.value = false;
            }
        };

        // Load the namespace list when the component is mounted
        onMounted(() => {
            extractNamespacesFromPods();
        });

        return {
            selectedNamespace,
            namespaces,
            loading,
            handleNamespaceChange
        };
    }
}
</script>

<style scoped>
.namespace-selector {
    display: inline-block;
    min-width: 180px;
}
</style>