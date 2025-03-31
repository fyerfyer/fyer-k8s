<template>
    <div class="pod-list">
        <div class="table-header">
            <h2>Kubernetes Pods</h2>
            <div class="table-actions">
                <!-- Namespace Selector -->
                <el-select 
                    v-model="selectedNamespace" 
                    placeholder="Select Namespace" 
                    size="small"
                    @change="fetchPods"
                >
                    <el-option label="All Namespaces" value="" />
                    <el-option 
                        v-for="ns in namespaces" 
                        :key="ns" 
                        :label="ns" 
                        :value="ns" 
                    />
                </el-select>
                
                <el-button type="primary" @click="fetchPods" :loading="loading" size="small">
                    <el-icon>
                        <refresh />
                    </el-icon>
                    Refresh
                </el-button>
            </div>
        </div>

        <el-table v-loading="loading" :data="pods" style="width: 100%" border>
            <el-table-column prop="name" label="Name" min-width="180" />
            <el-table-column prop="namespace" label="Namespace" min-width="120" />
            <el-table-column prop="status" label="Status" min-width="100">
                <template #default="scope">
                    <el-tag :type="getPodStatusType(scope.row.status)">
                        {{ scope.row.status }}
                    </el-tag>
                </template>
            </el-table-column>
            <el-table-column prop="podIP" label="Pod IP" min-width="120" />
            <el-table-column prop="nodeName" label="Node" min-width="150" />
            <el-table-column label="Created" min-width="180">
                <template #default="scope">
                    {{ formatDate(scope.row.creationTime) }}
                </template>
            </el-table-column>
            <el-table-column label="Actions" width="120" fixed="right">
                <template #default="scope">
                    <el-button 
                        type="primary" 
                        size="small" 
                        @click="viewPodDetails(scope.row)"
                        :disabled="loading"
                    >
                        View
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Refresh } from '@element-plus/icons-vue'
import kubernetesApi from '@/api/kubernetes'
import { useRouter } from 'vue-router'

export default {
    name: 'PodList',
    components: {
        Refresh
    },
    setup() {
        const pods = ref([])
        const loading = ref(false)
        const selectedNamespace = ref('')
        const namespaces = ref(['default', 'kube-system', 'kube-public'])
        const router = useRouter()

        // Fetch the list of Pods
        const fetchPods = async () => {
            loading.value = true
            try {
                const response = await kubernetesApi.listPods(selectedNamespace.value)
                if (response && response.pods) {
                    pods.value = response.pods
                    
                    // Extract unique namespace list
                    const uniqueNamespaces = [...new Set(response.pods.map(pod => pod.namespace))]
                    if (uniqueNamespaces.length > 0) {
                        namespaces.value = uniqueNamespaces
                    }
                    
                    ElMessage({
                        type: 'success',
                        message: `Loaded ${response.pods.length} pods`,
                        duration: 2000
                    })
                } else {
                    ElMessage.warning('No pods data received')
                }
            } catch (error) {
                console.error('Failed to fetch pods:', error)
                ElMessage.error(`Failed to fetch pods: ${error.message || 'Unknown error'}`)
            } finally {
                loading.value = false
            }
        }

        // View Pod details
        const viewPodDetails = (pod) => {
            router.push(`/pods/${pod.namespace}/${pod.name}`)
        }

        // Return different tag types based on Pod status
        const getPodStatusType = (status) => {
            const types = {
                'Running': 'success',
                'Pending': 'warning',
                'Failed': 'danger',
                'Succeeded': 'info',
                'Unknown': 'info'
            }
            return types[status] || 'info'
        }

        // Format date
        const formatDate = (dateString) => {
            if (!dateString) return '-'
            const date = new Date(dateString)
            return date.toLocaleString()
        }

        // Load data when the component is mounted
        onMounted(() => {
            fetchPods()
        })

        return {
            pods,
            loading,
            selectedNamespace,
            namespaces,
            fetchPods,
            viewPodDetails,
            getPodStatusType,
            formatDate
        }
    }
}
</script>

<style scoped>
.pod-list {
    padding: 20px;
    background-color: white;
    border-radius: 4px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.table-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.table-actions {
    display: flex;
    gap: 10px;
    align-items: center;
}
</style>