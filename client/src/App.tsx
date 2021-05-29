import { useEffect, useState } from "react";
import { Todo } from "./models/todos";
import { todoService } from "./services/todos-service";

function App() {
  const [todos ,setTodos] = useState([] as Todo[])

  useEffect(() => {
    todoService.getAllTodos().then(t => setTodos(t)).catch(err => console.log(err))
  },[])

  return (
   <>

      <div> Header here </div>
      
      {todos.map(todo => <p id = {todo.id}> {todo.name} </p>)}
      
    </>
  );
}

export default App;
