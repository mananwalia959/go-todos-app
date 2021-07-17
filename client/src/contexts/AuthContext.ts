import { UserPrincipal } from './../models/auth/UserPrincipal';
import React from 'react'

type AuthContextType = {
    userPrincipal?: UserPrincipal ,
    token?:string
    setToken :(token:string) => void
}

const AuthContext = React.createContext({} as AuthContextType)

export default AuthContext
