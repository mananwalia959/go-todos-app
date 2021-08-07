
import { Todo } from "../models/Todo";
import { TodoCreateRequest } from "../models/TodoCreateRequest";
import { TodoEditRequest } from "../models/TodoEditRequest";
import ProtectedAxios from "./axios/ProtectedAxios";

interface TodosService{
    getAllTodos(): Promise<Todo[]>
    saveNewTodo(todoCreateRequest : TodoCreateRequest): Promise<Todo>
    editTodo(todoId:string,todoEditRequest : TodoEditRequest):Promise<Todo>
    deleteTodo(todoId : string) : Promise<void>
}

class TodosServiceImpl implements TodosService{
    getAllTodos(): Promise<Todo[]>{
        return ProtectedAxios.get<Todo[]>(`/api/todos`).then(resp => resp.data)
    }
    saveNewTodo(todoCreateRequest : TodoCreateRequest): Promise<Todo> {
        return ProtectedAxios.post<Todo>(`/api/todos`,todoCreateRequest).then(resp => resp.data);
    }
    editTodo(todoId :string ,todoEditRequest: TodoEditRequest): Promise<Todo> {
        return ProtectedAxios.put<Todo>(`/api/todos/${todoId}`,todoEditRequest).then(resp => resp.data);
    }

    deleteTodo(todoId : string) : Promise<void> {
        return ProtectedAxios.delete<void>(`/api/todos/${todoId}`).then();
    }
}

const instance:TodosService = new TodosServiceImpl();

export {instance as todoService}

 