import "./register.css"

export const Register = () => {
	return (
		<div>
			<h2> Register </h2>
			<div className="row">
				<div className="col-7">
					<form method="post">
						<div className="form-group">
							<label htmlFor="exampleInputEmail1">Email address</label>
							<input type="email" className="form-control" id="email" name="email" aria-describedby="emailHelp" 
					placeholder="Enter email" required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputEmail1">Username</label>
							<input className="form-control" id="username" name="username" placeholder="Enter username" required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputPassword1">Password</label>
							<input type="password" className="form-control" id="password" name="password" 
					placeholder="Enter password" required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleInputPassword1">Confirm</label>
							<input type="password" className="form-control" id="confirm" name="confirm" 
					placeholder="Enter password again" required />
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
