import { TokenResponse } from './../models/auth/TokenResponse';
import  axios  from 'axios';

interface AuthService{
    getToken(code :string): Promise<TokenResponse>
}

class AuthServiceImpl implements AuthService {
    getToken(code:string): Promise<TokenResponse>  {
        return axios.post<TokenResponse>(`/api/auth/token/google`,{code}).then(resp => resp.data);
    }
}

const instance:AuthService = new AuthServiceImpl();
export { instance as authService}