/* View Notes (Dashboard) */

import { useEffect, useState, ChangeEventHandler, FormEventHandler, MouseEventHandler } from "react";
import { getToken } from "components/token/token";
import axios from "axios";
import "./view_notes.css";
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';

export default function Notes() {
	/* View Notes Page (Dashboard Page) */
	const [notes, setNotes] = useState([]);
	const [note_id, setNoteId] = useState<number | undefined>(undefined);

	const [title_edit, setTitleEdit] = useState('');
	const [description_edit, setDescriptionEdit] = useState('');

	const [title_add, setTitleAdd] = useState('');
	const [description_add, setDescriptionAdd] = useState('');

	const [show_add_note, setShowAddNote] = useState(false);
	const [show_edit_note, setShowEditNote] = useState(false);
	const [show_delete_note_confirm, setShowDeleteNoteConfirm] = useState(false);

	const handleCloseAddNote = () => {
		setShowAddNote(false);
		setTitleAdd("");
		setDescriptionAdd("");
	};

	const handleCloseEditNote = () => {
		setShowEditNote(false);
		setTitleEdit("");
		setDescriptionEdit("");
	};

	const handleCloseDeleteNoteConfirm = () => {
		setShowDeleteNoteConfirm(false);
	};

	const handleShowAddNote = () => setShowAddNote(true);

	function handleShowDeleteNoteConfirm(note_id: string) {
		/* Open Delete Confirm Popup Window */
		setNoteId(parseInt(note_id));
		setShowDeleteNoteConfirm(true);
	}

	function handleShowEditNote(note_id: string) {
		/* Open Edit Note Form Popup Window */
		// Fetch Note
		axios.get(`http://localhost:8080/api/fetch-note?id=` + note_id)
		.then((response) => {
			if (note_id !== undefined) {
				setTitleEdit(response.data.title);
				setDescriptionEdit(response.data.description);
			}
		}).catch(e => {
			console.log(e);
		})

		// set note id
		setNoteId(parseInt(note_id));
		setShowEditNote(true)
	}

	/*	
	Handle title and description changes
	for the edit and add note forms 
	*/
	const handleTitleEditChange: ChangeEventHandler = e => {
		const target = e.target as HTMLInputElement;
		setTitleEdit(target.value);
	};
	const handleDescriptionEditChange: ChangeEventHandler = e => {
		const target = e.target as HTMLInputElement;
		setDescriptionEdit(target.value);
	};

	const handleTitleAddChange: ChangeEventHandler = e => {
		const target = e.target as HTMLInputElement;
		setTitleAdd(target.value);
	};
	const handleDescriptionAddChange: ChangeEventHandler = e => {
		const target = e.target as HTMLInputElement;
		setDescriptionAdd(target.value);
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

	/* Add Note */
	const handleAddSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
		/* Add New Note Submission */
		e.preventDefault();
		const token = getToken();
		let user_id = undefined;
		if (token !== undefined) {
			axios.post(`http://localhost:8080/api/get-decoded-token`, {
				token: token,
			}).then((response) => {
				if (response.data.user_id !== undefined) {
					user_id = response.data.user_id;
					axios.post(`http://localhost:8080/api/add-new-note`, {
						title: title_add,
						description: description_add,
						user_id: user_id,
					}).then(() => {
						// redirect to the dashboard
						window.location.href = '/dashboard';
					}).catch(e => {
						console.log(e);
					})
				}
			}).catch(e => {
				console.log(e);
			})
		}
	};

	/* Edit Note */
	const handleEditSubmit: MouseEventHandler<HTMLFormElement> = async (e) => {
		e.preventDefault();
		axios.patch(`http://localhost:8080/api/edit-note`, {
			note_id: note_id,
			title: title_edit,
			description: description_edit,
		}).then(() => {
			// redirect to the dashboard
			window.location.reload();
		}).catch(e => {
			console.log(e);
		})
	};

	/* Delete Note */
	const handleDeleteNote: FormEventHandler<HTMLButtonElement> = async (e) => {
		e.preventDefault();
		axios.delete(`http://localhost:8080/api/delete-note`, {
			data: {note_id: note_id},
		}).then(() => {
			// redirect to the dashboard
			window.location.reload();
		}).catch(e => {
			console.log(e);
		})
	};

	return (
		<div>
			<h2> Notes </h2>
			<Button variant="primary" onClick={handleShowAddNote}>
				Add
			</Button>
			{notes.map((note, i) => {
				return (
					<div className="container note-entry" key={i}>
						<h2 className="note-title"> {note["Title"]} </h2>
						<p> {note["Description"]} </p>
						<div className="row">
							<div className="col col-md-auto">
								<Button variant="primary" onClick={() => handleShowEditNote(note["ID"])}>
									Edit
								</Button>
							</div>
							<div className="col col-md-auto">
								<Button variant="danger" onClick={() => handleShowDeleteNoteConfirm(note["ID"])}>
									Delete
								</Button>
							</div>
						</div>
					</div>
				)
			})}

			{/* Add Note */}
			<Modal className="modal" show={show_add_note} onHide={handleCloseAddNote}>
				<Modal.Header closeButton>
					<Modal.Title>Add Note</Modal.Title>
				</Modal.Header>
				<Modal.Body>
					<Form id="myform" method="post" onSubmit={handleAddSubmit}>
						<Form.Group className="mb-3">
							<Form.Label>Title</Form.Label>
							<Form.Control name="title" placeholder="Title" 
								value={title_add}
								onChange={handleTitleAddChange}
								required/>
						</Form.Group>
						<Form.Group className="mb-3">
							<Form.Label>Description</Form.Label>
							<Form.Control name="description" rows={3}
								as="textarea" 
								placeholder="Description"
								value={description_add}
								onChange={handleDescriptionAddChange}
								required/>
						</Form.Group>
					</Form>
				</Modal.Body>
				<Modal.Footer>
					<Button variant="primary" type="submit" form="myform">
						Submit
					</Button>
					<Button variant="secondary" onClick={handleCloseAddNote}>
						Close
					</Button>
				</Modal.Footer>
			</Modal>

			{/* Edit Note */}
			<Modal show={show_edit_note} onHide={handleCloseEditNote}>
				<Modal.Header closeButton>
					<Modal.Title>Edit Note</Modal.Title>
				</Modal.Header>
				<Modal.Body>
					<Form id="myform" method="post" onSubmit={handleEditSubmit}>
						<Form.Group className="mb-3">
							<Form.Label>Title</Form.Label>
							<Form.Control name="title" placeholder="Title" 
								value={title_edit}
								onChange={handleTitleEditChange}
								required/>
						</Form.Group>
						<Form.Group className="mb-3">
							<Form.Label>Description</Form.Label>
							<Form.Control name="description" rows={3}
								as="textarea" 
								placeholder="Description"
								value={description_edit}
								onChange={handleDescriptionEditChange}
								required/>
						</Form.Group>
					</Form>
				</Modal.Body>
				<Modal.Footer>
					<Button variant="primary" type="submit" form="myform">
						Submit
					</Button>
					<Button variant="secondary" onClick={handleCloseEditNote}>
						Close
					</Button>
				</Modal.Footer>
			</Modal>

			{/* Delete Note */}
			<Modal show={show_delete_note_confirm} onHide={handleCloseDeleteNoteConfirm}>
				<Modal.Header closeButton>
					<Modal.Title>Delete Note</Modal.Title>
				</Modal.Header>
				<Modal.Body>
					<p> You sure you want to delete the note? </p>
				</Modal.Body>
				<Modal.Footer>
					<Button variant="danger" onClick={handleDeleteNote}>
						Delete
					</Button>
					<Button variant="secondary" onClick={handleCloseDeleteNoteConfirm}>
						Close
					</Button>
				</Modal.Footer>
			</Modal>
		</div>
	);
}
