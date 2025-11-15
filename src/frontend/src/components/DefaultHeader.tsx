import { Link } from "react-router";

export default function DefaultHeader() {
  const email = getLoggedInEmail();
  return (
    <header>
      <span id="header-text"><Link to="/">Personal Blog</Link></span>
      <span id="email">{email ? email : <Link to="/signup">Sign Up</Link>}</span>
    </header>
  )
}

function getLoggedInEmail() {
  return localStorage.getItem("email");
}
