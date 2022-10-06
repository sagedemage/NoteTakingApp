import {useEffect} from "react";
import "./register.css"

import PasswordValidator from "./password-validator"

import axios from "axios";
import {useState} from "react";

export const Register = () => {
	useEffect(() => {
		// Start Password Validator
		PasswordValidator();
	}, []);

	/* Registration Submition */

	const [email, setEmail] = useState('');
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const [confirm_pwd, setConfirm] = useState('');

	const [error_status, setErrorStatus] = useState(false);
	const [msg_error, setMsgError] = useState('');

	const handleEmailChange = event => {
    	setEmail(event.target.value);
  	};

	const handleUsernameChange = event => {
    	setUsername(event.target.value);
  	};

	const handlePasswordChange = event => {
    	setPassword(event.target.value);
  	};

	const handleConfirmChange = event => {
    	setConfirm(event.target.value);
  	};

	const handleSubmit = async (e) => {
		e.preventDefault();
		console.log(email, username, password, confirm_pwd);
		axios.post(`http://localhost:8080/api/register`, {
			email: email,
			username: username,
			password: password,
			confirm_pwd: confirm_pwd,
		}).then((response) => {
			if (response.data.registered === true) {
				// create new url of the login page
				var url = new URL("/login", "http://localhost:3000");

				// add url parameter
				url.searchParams.append("msg_success", response.data.msg_success);

				// redirect to the login page
				window.location.href = url;
			}
			else {
				setErrorStatus(true);
				setMsgError(response.data.msg_error);
			}
		}).catch(e => {
            console.log(e)
        })
	};

	return (
		<div>
			{ error_status === true &&
			<div class="alert alert-danger" role="alert">
				{ msg_error } 
			</div>
			}
			<h2> Register </h2>
			<div className="row">
				<div className="col-7">
					<form method="post" onSubmit={handleSubmit}>
						<div className="form-group">
							<label htmlFor="exampleInputEmail1">Email address</label>
							<input type="email" className="form-control" id="email" 
								name="email" placeholder="Enter email"
								aria-describedby="emailHelp" 
								value={email}
								onChange={handleEmailChange}
					 		required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputUsername1">Username</label>
							<input className="form-control" id="username" 
								name="username" placeholder="Enter username"
								value={username}
								onChange={handleUsernameChange}
							required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputPassword1">Password</label>
							<input type="password" className="form-control" id="password" 
								name="password" placeholder="Enter password" 
								value={password}
								onChange={handlePasswordChange}
							required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputPassword1">Confirm</label>
							<input type="password" className="form-control" id="confirm" 
								name="confirm" placeholder="Enter password again"
								value={confirm_pwd}
								onChange={handleConfirmChange}
							required />
						</div>
						<button type="submit" className="btn btn-primary">Submit</button>
					</form>
				</div>
				<div id="message" className="col-5">
					<div className="progress">
						<div id="p-bar" className="progress-bar bg-success" role="progressbar" 
							style={{ width: 0 + '%' }} aria-valuenow="25" aria-valuemin="0" aria-valuemax="100">
						</div>
					</div>
					<p id="has_lowercase"> contains a lowercase letter </p>
					<p id="has_uppercase"> contains an uppercase letter </p>
					<p id="has_number"> contains a number </p>
					<p id="good_password_length"> minimum of 8 characters </p>
				</div>
			</div>
		</div>
	);
}
