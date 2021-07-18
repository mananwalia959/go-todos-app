
interface StorageService {
    setToken(token:string):void
    getToken():string|undefined
}

class LocalStorageServiceImpl implements StorageService{
    setToken(token: string): void {
        localStorage.setItem("JWT_TOKEN", token)
    }
    getToken(): string|undefined {
        return localStorage.getItem("JWT_TOKEN") || undefined
    }
    
}

const storageService:StorageService = new LocalStorageServiceImpl()

export default storageService;

