import axios from "axios";
import StorageService from "../StorageService";
 
const ProtectedAxios = axios.create()
ProtectedAxios.interceptors.request.use(function (config) {
    const token = StorageService.getToken();
    
    config.headers.Authorization =  token ? `Bearer ${token}` : '';
    return config;
  });

export default ProtectedAxios;