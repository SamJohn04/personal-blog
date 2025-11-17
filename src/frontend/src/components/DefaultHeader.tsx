import { Link, useNavigate } from "react-router";

export default function DefaultHeader() {
  const navigate = useNavigate();
  const email = getLoggedInEmail();
  
  function logout() {
    localStorage.removeItem("authLevel");
    localStorage.removeItem("authToken");
    localStorage.removeItem("email");

    navigate("/");
  }

  return (
    <header>
      <span id="header-text"><Link to="/">Samuel's Mind</Link></span>
      <span id="email">{email ?
        <span>{email} <button className="min-y-pad" onClick={logout}>Log Out</button></span>
          : <Link to="/signup">Sign Up</Link>}</span>
    </header>
  )
}

function getLoggedInEmail() {
  return localStorage.getItem("email");
}
