/* View Notes Page (Dashboard Page) */

import {useEffect, useState} from "react";
import Cookies from "universal-cookie";
import axios from "axios";
import "./view-notes.css";

export const Notes = () => {

	const [notes, setNotes] = useState([]);

	useEffect(() => {
		/* Fetch all the Notes for the Current User */
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

	function DeleteNote(note_id) {
		/* Delete Note Page Redirection */
		// create new url of the delete note page
		var url = new URL("/delete-note", "http://localhost:3000");

		// add url parameter
		url.searchParams.append("note_id", note_id);

		// redirect to that url
		window.location.href = url;
	}

	function EditNote(note_id) {
		/* Edit Note Page Redirection */
		// create new url of the edit note page
		var url = new URL("/edit-note", "http://localhost:3000");

		// add url parameter
		url.searchParams.append("note_id", note_id);

		// redirect to that url
		window.location.href = url;
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
							<button className="btn btn-primary"
								onClick={ () => EditNote(note.ID) }>Edit</button>
						</div>
						<div className="col col-md-auto">	
							<button className="btn btn-danger" 
							onClick={ () => DeleteNote(note.ID) }>Delete</button>
						</div>
					</div>
				</div>
				)
			})}
		</div>
	);
}
