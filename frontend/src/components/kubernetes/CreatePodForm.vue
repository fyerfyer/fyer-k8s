<template>
  <div class="create-pod-form">
    <el-tabs v-model="activeTab">
      <el-tab-pane label="Form" name="form">
        <el-form :model="podForm" label-width="120px" label-position="left">
          <el-form-item label="Namespace" required>
            <el-select 
              v-model="podForm.namespace" 
              placeholder="Select Namespace"
              style="width: 100%"
            >
              <el-option 
                v-for="ns in namespaces" 
                :key="ns" 
                :label="ns" 
                :value="ns" 
              />
            </el-select>
          </el-form-item>
          
          <el-form-item label="Pod Name" required>
            <el-input v-model="podForm.name" placeholder="Enter pod name" />
          </el-form-item>
          
          <el-divider>Container</el-divider>
          
          <el-form-item label="Container Name" required>
            <el-input v-model="podForm.containerName" placeholder="Enter container name" />
          </el-form-item>
          
          <el-form-item label="Image" required>
            <el-input v-model="podForm.image" placeholder="Enter container image (e.g. nginx:latest)" />
          </el-form-item>
          
          <el-form-item label="Command">
            <el-input 
              v-model="podForm.command" 
              placeholder="Optional: Command to run (e.g. /bin/sh -c 'echo hello')" 
            />
          </el-form-item>
          
          <el-collapse>
            <el-collapse-item title="Environment Variables" name="env">
              <div v-for="(env, index) in podForm.envVars" :key="index" class="env-row">
                <el-input v-model="env.name" placeholder="Name" class="env-input" />
                <el-input v-model="env.value" placeholder="Value" class="env-input" />
                <el-button 
                  type="danger" 
                  circle 
                  @click="removeEnvVar(index)" 
                  :disabled="podForm.envVars.length <= 1"
                  size="small"
                >
                  <el-icon><delete /></el-icon>
                </el-button>
              </div>
              <el-button type="primary" size="small" @click="addEnvVar" plain class="add-button">
                <el-icon><plus /></el-icon> Add Environment Variable
              </el-button>
            </el-collapse-item>
          </el-collapse>
        </el-form>
      </el-tab-pane>
      
      <el-tab-pane label="YAML" name="yaml">
        <el-input
          v-model="yamlContent"
          type="textarea"
          :rows="20"
          placeholder="Paste your Pod YAML here"
        />
      </el-tab-pane>
    </el-tabs>
    
    <div class="form-actions">
      <el-button @click="$emit('cancel')">Cancel</el-button>
      <el-button type="primary" @click="submitForm" :loading="loading">Create Pod</el-button>
    </div>
  </div>
</template>

<script>
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Plus, Delete } from '@element-plus/icons-vue'
import kubernetesApi from '@/api/kubernetes'
import yaml from 'js-yaml'

export default {
  name: 'CreatePodForm',
  components: {
    Plus,
    Delete
  },
  props: {
    namespaces: {
      type: Array,
      default: () => ['default']
    },
    initialNamespace: {
      type: String,
      default: 'default'
    }
  },
  emits: ['created', 'cancel'],
  setup(props, { emit }) {
    const activeTab = ref('form')
    const loading = ref(false)
    
    // Form data
    const podForm = ref({
      namespace: props.initialNamespace,
      name: '',
      containerName: '',
      image: '',
      command: '',
      envVars: [{ name: '', value: '' }]
    })
    
    // YAML representation
    const yamlContent = ref('')
    
    // Generate YAML from form
    const generateYaml = () => {
      // Basic pod structure
      const podSpec = {
        apiVersion: 'v1',
        kind: 'Pod',
        metadata: {
          name: podForm.value.name,
          namespace: podForm.value.namespace
        },
        spec: {
          containers: [
            {
              name: podForm.value.containerName || podForm.value.name,
              image: podForm.value.image
            }
          ]
        }
      }
      
      // Add command if specified
      if (podForm.value.command) {
        // Split the command string into an array
        podSpec.spec.containers[0].command = podForm.value.command.split(' ')
      }
      
      // Add environment variables if any are defined
      const validEnvVars = podForm.value.envVars.filter(env => env.name && env.value)
      if (validEnvVars.length > 0) {
        podSpec.spec.containers[0].env = validEnvVars.map(env => ({
          name: env.name,
          value: env.value
        }))
      }
      
      return yaml.dump(podSpec)
    }
    
    // Add a new environment variable field
    const addEnvVar = () => {
      podForm.value.envVars.push({ name: '', value: '' })
    }
    
    // Remove an environment variable field
    const removeEnvVar = (index) => {
      if (podForm.value.envVars.length > 1) {
        podForm.value.envVars.splice(index, 1)
      }
    }
    
    // Update YAML when form changes
    watch(podForm, () => {
      if (activeTab.value === 'yaml') {
        yamlContent.value = generateYaml()
      }
    }, { deep: true })
    
    // Update form when switching to YAML tab
    watch(activeTab, (newValue) => {
      if (newValue === 'yaml') {
        yamlContent.value = generateYaml()
      }
    })
    
    // Submit the form
    const submitForm = async () => {
      // Validation
      if (activeTab.value === 'form') {
        if (!podForm.value.namespace) {
          ElMessage.warning('Please select a namespace')
          return
        }
        if (!podForm.value.name) {
          ElMessage.warning('Please enter a pod name')
          return
        }
        if (!podForm.value.image) {
          ElMessage.warning('Please enter a container image')
          return
        }
      } else if (!yamlContent.value) {
        ElMessage.warning('YAML content cannot be empty')
        return
      }
      
      loading.value = true
      try {
        if (activeTab.value === 'form') {
          // Create pod using the form data
          const containerName = podForm.value.containerName || podForm.value.name
          
          // Basic pod spec
          const podSpec = {
            name: podForm.value.name,
            containers: [{
              name: containerName,
              image: podForm.value.image
            }]
          }
          
          // Add command if specified
          if (podForm.value.command) {
            podSpec.containers[0].command = podForm.value.command.split(' ')
          }
          
          // Add environment variables
          const validEnvVars = podForm.value.envVars.filter(env => env.name && env.value)
          if (validEnvVars.length > 0) {
            podSpec.containers[0].env = validEnvVars.map(env => ({
              name: env.name,
              value: env.value
            }))
          }
          
          await kubernetesApi.createPod(podForm.value.namespace, podSpec)
        } else {
          // Create pod using the YAML content
          try {
            const podSpec = yaml.load(yamlContent.value)
            const namespace = podSpec.metadata?.namespace || 'default'
            
            // Create pod from the parsed YAML
            await kubernetesApi.createPod(namespace, podSpec)
          } catch (yamlError) {
            ElMessage.error(`Invalid YAML: ${yamlError.message}`)
            loading.value = false
            return
          }
        }
        
        ElMessage.success('Pod created successfully')
        emit('created')
      } catch (error) {
        console.error('Failed to create pod:', error)
        ElMessage.error(`Failed to create pod: ${error.message || 'Unknown error'}`)
      } finally {
        loading.value = false
      }
    }
    
    return {
      activeTab,
      loading,
      podForm,
      yamlContent,
      addEnvVar,
      removeEnvVar,
      submitForm
    }
  }
}
</script>

<style scoped>
.create-pod-form {
  padding: 20px 0;
}

.env-row {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
  gap: 10px;
}

.env-input {
  flex: 1;
}

.add-button {
  margin-top: 10px;
}

.form-actions {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>