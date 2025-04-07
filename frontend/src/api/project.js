import apiAxios from '../utils/apiAxios'
export const getProjectList = (...args) => apiAxios.get('/project/list', ...args)
export const getProjectDetails = (id) => apiAxios.get(`/project/details/${id}`)