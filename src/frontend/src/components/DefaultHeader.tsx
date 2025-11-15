export default function DefaultHeader() {
  const email = getLoggedInEmail();
  return (
    <header>
      <span id="header-text">Personal Blog</span>
      <span id="email">{email ? email : <a href="/signup">Sign Up</a>}</span>
    </header>
  )
}

function getLoggedInEmail() {
  return localStorage.getItem("email");
}
