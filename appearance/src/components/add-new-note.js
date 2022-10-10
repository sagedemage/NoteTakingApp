import axios from "axios";
import {useState} from "react";
import Cookies from 'universal-cookie';

export const AddNoteForm = () => {
	const [title, setTitle] = useState('');
	const [description, setDescription] = useState('');

	const handleTitleChange = event => {
    	setTitle(event.target.value);
  	};

	const handleDescriptionChange = event => {
    	setDescription(event.target.value);
  	};

	const handleSubmit = async (e) => {
		e.preventDefault();
		const cookies = new Cookies();
		const user_id = parseInt(cookies.get("user_id"));
		axios.post(`http://localhost:8080/api/add-new-note`, {
			title: title,
			description: description,
			user_id: user_id,
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
			<h2> Add Note </h2>
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
						name="description" rows="3" placeholder="Description" 
						value={description}
						onChange={handleDescriptionChange}
						required> </textarea>
				</div>
				<button type="submit" className="btn btn-primary">Submit</button>
			</form>
		</div>
	);
}