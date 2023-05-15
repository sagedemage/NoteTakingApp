/* Render JSX */

import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./app";
import { BrowserRouter } from "react-router-dom";

const root_element = document.getElementById("root")!;
const root = ReactDOM.createRoot(root_element);
root.render(
  <BrowserRouter>
    <App />
  </BrowserRouter>
);

