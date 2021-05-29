import axios from "axios";
import { Todo } from "../models/todos";

interface TodosService{
    getAllTodos(): Promise<Todo[]>
}

class TodosServiceImpl implements TodosService{
    getAllTodos(): Promise<Todo[]>{
       return axios.get<Todo[]>("/api/todos").then(resp => resp.data)
    }
}

const instance:TodosService = new TodosServiceImpl();

export {instance as todoService}

 