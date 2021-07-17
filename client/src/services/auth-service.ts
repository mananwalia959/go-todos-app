import { TokenResponse } from './../models/auth/TokenResponse';
import  axios  from 'axios';

interface AuthService{
    getToken(code :string): Promise<TokenResponse>
    getOauthUrl(): Promise<any>
}

class AuthServiceImpl implements AuthService {
    getToken(code:string): Promise<TokenResponse>  {
        return axios.post<TokenResponse>(`/api/auth/token/google`,{code}).then(resp => resp.data);
    }

    getOauthUrl(): Promise<any> {
        return axios.get<OauthResponse>('/api/auth/loginurl/google', {headers :
             {'Accept': 'application/json',
             'Content-Type': 'application/json'}
            }).then(res => window.location.href = (res.headers.location));
    }
}

type OauthResponse = {
    redirectUrl : string
}

const instance:AuthService = new AuthServiceImpl();
export { instance as authService}