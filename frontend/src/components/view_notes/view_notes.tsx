/* View Notes (Dashboard) */

import { useEffect, useState, ChangeEventHandler, FormEventHandler } from "react";
import { getToken } from "components/token/token";
import axios from "axios";
import "./view_notes.css";

//import EditNoteForm from "components/note_actions/edit_note";

export default function Notes() {
	/* View Notes Page (Dashboard Page) */
	const [notes, setNotes] = useState([]);
	const [open_edit, setOpenEdit] = useState(false);

	const [title, setTitle] = useState('');
	const [description, setDescription] = useState('');

	const handleTitleChange: ChangeEventHandler = e => {
		const target = e.target as HTMLInputElement;
    	setTitle(target.value);
  	};
	const handleDescriptionChange: ChangeEventHandler = e => {
		const target = e.target as HTMLInputElement;
    	setDescription(target.value);
  	};

	const Close = () => {
		setOpenEdit(false);
	}

	const handleEditSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
		e.preventDefault();
		let note_id_string: string = localStorage.getItem("note_id")!;
		let note_id = parseInt(note_id_string);

		console.log(typeof note_id);
		axios.post(`http://localhost:8080/api/edit-note`, {
			note_id: note_id,
			title: title,
			description: description,
		}).then(() => {
			// redirect to the dashboard
			window.location.reload();
		}).catch(e => {
            console.log(e);
        })
	};

	useEffect(() => {
		/* Fetch all the Notes for the Current User */
		const token = getToken();
		let user_id = undefined;
		if (token !== undefined) {
			axios.post(`http://localhost:8080/api/get-decoded-token`, {
				token: token,
			}).then((response) => {
				if (response.data.user_id !== undefined) {
					user_id = response.data.user_id;
					axios.post(`http://localhost:8080/api/view-notes`, {
						user_id: user_id,
					}).then((response) => {
						setNotes(response.data.notes);
					}).catch(e => {
						console.log(e);
					})
				}

			}).catch(e => {
				console.log(e);
			})
		}
	}, []);

	const AddNewNote = () => {
		window.location.href = '/add-new-note';
	}

	function DeleteNote(note_id: string) {
		/* Delete Note Page Redirection */
		// create new url of the delete note page
		var url = new URL("/delete-note", "http://localhost:3000");

		// add url parameter
		url.searchParams.append("note_id", note_id);

		// redirect to that url
		window.location.href = String(url);
	}
	function EditNote(note_id: string) {
		/* Edit Note Page Redirection */
		// create new url of the edit note page
		//var url = new URL("/edit-note", "http://localhost:3000");

		// add url parameter
		//url.searchParams.append("note_id", note_id);

		/* Fetch Note */
		axios.post(`http://localhost:8080/api/fetch-note`, {
			note_id: note_id,
		}).then((response) => {
			if (note_id !== undefined) {
				setTitle(response.data.title);
				setDescription(response.data.description);
			}
		}).catch(e => {
            console.log(e);
        })

		localStorage.setItem("note_id", note_id);

		// redirect to that url
		//window.location.href = String(url);

		setOpenEdit(true);
	}

	return (
		<div>
			<h2> Notes </h2>
			<form method="post">
				<button type="button" className="btn btn-primary" onClick={AddNewNote}>
					Add New Note
				</button>
			</form>
			{notes.map((note, i) => {
				return (
					<div className="container note-entry" key={i}>
						<h2 className="note-title"> {note["Title"]} </h2>
						<p> {note["Description"]} </p>
						<div className="row">
							<div className="col col-md-auto">
								<button className="btn btn-primary"
									onClick={() => EditNote(note["ID"])}>Edit</button>
							</div>
							<div className="col col-md-auto">
								<button className="btn btn-danger"
									onClick={() => DeleteNote(note["ID"])}>Delete</button>
							</div>
						</div>
					</div>
				)
			})}
			{open_edit === true &&
				<div className="box">
					<h2> Edit Note </h2>
					<form method="post" onSubmit={handleEditSubmit}>
						<div className="form-group">
							<label htmlFor="exampleFormControlInput1">Title</label>
							<input className="form-control" id="exampleFormControlInput1"
								name="title" placeholder="Title"
								value={title}
								onChange={handleTitleChange}
								required />
						</div>
						<div className="form-group">
							<label htmlFor="exampleFormControlTextarea1">Description</label>
							<textarea className="form-control" id="exampleFormControlTextarea1"
								name="description" rows={3} placeholder="Description"
								value={description}
								onChange={handleDescriptionChange}
								required> </textarea>
						</div>

						<button type="submit" className="btn btn-primary">Submit</button>
						<button type="button" className="btn btn-secondary"
							onClick={Close}>
							Close
						</button>
					</form>
				</div>
			}
		</div>
	);
}
