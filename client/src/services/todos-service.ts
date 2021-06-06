
import axios from "axios";
import { Todo } from "../models/Todo";
import { TodoCreateRequest } from "../models/TodoCreateRequest";

interface TodosService{
    getAllTodos(): Promise<Todo[]>
    saveNewTodo(todoCreateRequest : TodoCreateRequest): Promise<Todo>
}

class TodosServiceImpl implements TodosService{
    getAllTodos(): Promise<Todo[]>{
       return axios.get<Todo[]>("/api/todos").then(resp => resp.data)
    }
    saveNewTodo(todoCreateRequest : TodoCreateRequest): Promise<Todo> {
        return axios.post<Todo>("/api/todos",todoCreateRequest).then(resp => resp.data);
    }
}

const instance:TodosService = new TodosServiceImpl();

export {instance as todoService}

 