import apiClient from './index';

/**
 * Kubernetes-related API functions
 */
const kubernetesApi = {
  /**
   * Ping the server to check if it's alive
   * @returns {Promise<Object>} The API response
   */
  pingServer() {
    return apiClient.get('/ping');
  },

  /**
   * Get a list of pods from the Kubernetes cluster
   * @param {string} namespace - Optional namespace to filter pods
   * @returns {Promise<Object>} The API response with pod list
   */
  listPods(namespace = '') {
    const url = namespace ? `/api/v1/pods?namespace=${namespace}` : '/api/v1/pods';
    return apiClient.get(url);
  },

  /**
   * Get details for a specific pod
   * @param {string} namespace - Pod namespace
   * @param {string} name - Pod name
   * @returns {Promise<Object>} The API response with pod details
   */
  getPodDetail(namespace, name) {
    return apiClient.get(`/api/v1/pods/${namespace}/${name}`);
  },

  /**
   * Create a new pod in the Kubernetes cluster
   * @param {string} namespace - Namespace for the new pod
   * @param {Object} podSpec - Pod specification
   * @returns {Promise<Object>} The API response with created pod
   */
  createPod(namespace, podSpec) {
    // 将命名空间包含在 podSpec 中，而不是 URL 路径中
    const payload = {
      ...podSpec,
      namespace: namespace
    };
    return apiClient.post('/api/v1/pods', payload);
  }
};

export default kubernetesApi;