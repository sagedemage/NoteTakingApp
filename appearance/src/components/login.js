import "./home.css"

export const Login = () => {
	return (
		<div>
			<h2> Login </h2>
			<form method="post">
				<div className="form-group">
					<label htmlFor="exampleInputUsername1">Email or Username</label>
					<input className="form-control" id="exampleInputUsername1" 
					name="username" placeholder="Enter email or username" required />
				</div>
				<div className="form-group">
					<label htmlFor="exampleInputPassword1">Password</label>
					<input type="password" className="form-control" id="exampleInputPassword1" name="password" 
					placeholder="Enter password" required />
				</div>
				<button type="submit" className="btn btn-primary">Submit</button>
			</form>
		</div>
	);
}
