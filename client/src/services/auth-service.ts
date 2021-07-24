import { UserPrincipal } from './../models/auth/UserPrincipal';
import { TokenResponse } from './../models/auth/TokenResponse';
import  axios  from 'axios';
import jwtDecode from 'jwt-decode';

interface AuthService{
    getToken(code :string): Promise<TokenResponse>
    getOauthUrl(): Promise<any>
    tokenToUserPrincipal(token:string): UserPrincipal
    isTokenValidTillNextDay(token:string):Promise<any>
}

class AuthServiceImpl implements AuthService {
    getToken(code:string): Promise<TokenResponse>  {
        return axios.post<TokenResponse>(`/api/auth/token/google`,{code}).then(resp => resp.data);
    }

    getOauthUrl(): Promise<any> {
        return axios.get<OauthResponse>('/api/auth/loginurl/google')
            .then(res => window.location.href = (res.headers.location));
    }
    //i can check expired field locally but in case if the secret is rotated , 
    //its better to validate this at the backend too 
    isTokenValidTillNextDay(token:string):Promise<any>{
        return axios.post<any>('/api/auth/validatetoken', {jwtToken: token})
    }

    tokenToUserPrincipal(token:string): UserPrincipal {
        return token ? jwtDecode(token) :{} as UserPrincipal
    }
}

type OauthResponse = {
    redirectUrl : string
}

const instance:AuthService = new AuthServiceImpl();
export { instance as authService}