import "./home.css"

import { useEffect } from "react";
import axios from "axios";
import {useState} from "react";

export const Login = () => {

	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');

	const handleUsernameChange = event => {
    	setUsername(event.target.value);
  	};

	const handlePasswordChange = event => {
    	setPassword(event.target.value);
  	};

	const handleSubmit = async (e) => {
		e.preventDefault();
		console.log(username, password);
		try {
			const resp = await axios.post(`http://localhost:8080/api/login`, {
				username: username,
				password: password,
			})
			console.log(resp.data);
			//window.location.href = '/dashboard';
		} 
		catch (error) {
			console.log(error.response)
		}
	};

	return (
		<div>
			<h2> Login </h2>
			<form method="post" onSubmit={handleSubmit}>
				<div className="form-group">
					<label htmlFor="exampleInputUsername1">Email or Username</label>
					<input className="form-control" id="exampleInputUsername1" 
						name="username" placeholder="Enter email or username" 
						value={username} 
						onChange={handleUsernameChange}
					required />
				</div>
				<div className="form-group">
					<label htmlFor="exampleInputPassword1">Password</label>
					<input type="password" className="form-control" id="exampleInputPassword1" 
						name="password" placeholder="Enter password" 
						value={password}
						onChange={handlePasswordChange}
					required />
				</div>
				<button type="submit" className="btn btn-primary">Submit</button>
			</form>
		</div>
	);
}
