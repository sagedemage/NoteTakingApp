/* Edit Note Page */

import axios from "axios";
import { useState, useEffect, ChangeEventHandler, FormEventHandler } from "react";

export const EditNoteForm = () => {
	
	let url = new URL(window.location.href);
	const note_id = parseInt(url.searchParams.get("note_id"));

	if (note_id === null) {
		window.location.href = '/';
	}

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
	const GoBack = () => {
		window.location.href='/dashboard';
	}

	useEffect(() => {
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
	}, []);

	const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
		e.preventDefault();
		axios.post(`http://localhost:8080/api/edit-note`, {
			note_id: note_id,
			title: title,
			description: description,
		}).then((response) => {
			// redirect to the dashboard
            window.location.href = '/dashboard';
            console.log(response);
		}).catch(e => {
            console.log(e);
        })
	};

	return (		
		<div>
			<h2> Edit Note </h2>
			<form method="post" onSubmit={handleSubmit}>
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
					onClick={ GoBack }>
					Back
				</button>
			</form>
		</div>
	);
}
