import React from "react";
import { BrowserRouter as Router, Route, Routes} from "react-router-dom";
import Login from "./pages/login";
import Register from "./pages/register";
import TodoList from "./pages/todo";

const App = () => {
	return (
		<Router>
			<Routes>
				<Route exact path="/" Component={Login} />
				<Route path="/register" Component={Register} />
				<Route path="/todo" Component={TodoList} />
			</Routes>
		</Router>
	)
}

export default App;
