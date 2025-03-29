import apiClient from './index';

/**
 * Kubernetes-related API functions
 */
const kubernetesApi = {
  /**
   * Get a list of pods from the Kubernetes cluster
   * @returns {Promise<Object>} The API response
   */
  listPods() {
    return apiClient.get('/pod/list');
  },

  /**
   * Ping the server to check if it's alive
   * @returns {Promise<Object>} The API response
   */
  pingServer() {
    return apiClient.get('/ping');
  }
};

export default kubernetesApi;