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
  }
};

export default kubernetesApi;