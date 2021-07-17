import { UserPrincipal } from './../models/auth/UserPrincipal';
import { TokenResponse } from './../models/auth/TokenResponse';
import  axios  from 'axios';

interface AuthService{
    getToken(code :string): Promise<TokenResponse>
    getOauthUrl(): Promise<any>
    tokenToUserPrincipal(token:string): UserPrincipal
}

class AuthServiceImpl implements AuthService {
    getToken(code:string): Promise<TokenResponse>  {
        return axios.post<TokenResponse>(`/api/auth/token/google`,{code}).then(resp => resp.data);
    }

    getOauthUrl(): Promise<any> {
        return axios.get<OauthResponse>('/api/auth/loginurl/google')
            .then(res => window.location.href = (res.headers.location));
    }

    tokenToUserPrincipal(token:string): UserPrincipal {
        return {} as UserPrincipal
    }
}

type OauthResponse = {
    redirectUrl : string
}

const instance:AuthService = new AuthServiceImpl();
export { instance as authService}