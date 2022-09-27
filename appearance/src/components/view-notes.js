import "./home.css";

export const Notes = () => {
	return (
		<div>
			<h2> Notes </h2>
			<form method="post">
				<button type="button" class="btn btn-primary" onclick="window.location.href='/add-new-note';">
					Add New Note
				</button>
			</form>
			<p>
				<div class="container">
					<h2 class="note-title"> Title </h2>
					<p> Description </p>
					<div class="row">
						<div class="col col-md-auto">
							<form method="post">
								<div class="form-group">
									<input type="text" id="edit" name="edit" value="0" hidden/>
								</div>
								<button type="submit" class="btn btn-primary">Edit</button>
							</form>
						</div>
						<div class="col col-md-auto">
							<form method="post">
								<div class="form-group">
									<input type="text" id="delete" name="delete" value="0" hidden/>
								</div>
								<button type="submit" class="btn btn-danger">Delete</button>
							</form>
						</div>
					</div>
				</div>
			</p>
		</div>
	);
}
