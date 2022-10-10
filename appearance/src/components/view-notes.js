import {useEffect, useState} from "react";
import Cookies from "universal-cookie";
import axios from "axios";
import "./view-notes.css";

export const Notes = () => {

	const [notes, setNotes] = useState([]);

	useEffect(() => {
		const cookies = new Cookies();
		const user_id = parseInt(cookies.get("user_id"));
		axios.post(`http://localhost:8080/api/view-notes`, {
			user_id: user_id,
		}).then((response) => {
			setNotes(response.data.notes);
            console.log("Response data " + response.data.notes);
		}).catch(e => {
            console.log(e);
        })
	}, []);

	console.log("Notes data: " + notes);

	const AddNewNote = () => {
		window.location.href='/add-new-note';
	}

	return (
		<div>
			<h2> Notes </h2>
			<form method="post">
				<button type="button" className="btn btn-primary" onClick={ AddNewNote }>
					Add New Note
				</button>
			</form>
			{notes.map((note, i) => {
				return (
				<div className="container note-entry" key={i}>
					<h2 className="note-title"> { note.Title } </h2>
					<p> { note.Description } </p>
					<div className="row">
						<div className="col col-md-auto">
							<form method="post">
								<div className="form-group">
									<input type="text" id="edit" 
										name="edit" 
										defaultValue={ note.ID }
										hidden
									/>
								</div>
								<button type="submit" className="btn btn-primary">Edit</button>
							</form>
						</div>
						<div className="col col-md-auto">
							<form method="post">
								<div className="form-group">
									<input type="text" id="delete" 
									name="delete" 
									defaultValue={ note.ID }
									hidden
								/>
								</div>
								<button type="submit" className="btn btn-danger">Delete</button>
							</form>
						</div>
					</div>
				</div>
				)
			})}
		</div>
	);
}
