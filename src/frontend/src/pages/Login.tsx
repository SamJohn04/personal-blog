import { useState, type FormEvent } from "react";
import DefaultHeader from "../components/DefaultHeader";
import { Link, useNavigate } from "react-router";

export default function Login() {
  const navigate = useNavigate();

  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const [loading, setLoading] = useState<boolean>(false);

  async function login(event: FormEvent) {
    event.preventDefault();
    setLoading(true);
    const res = await fetch("/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email,
        password,
      }),
    });
    if (!res.ok) {
      alert("Something went wrong...");
      setLoading(false);
      return;
    }
    const body = await res.json();
    localStorage.setItem("authToken", body.token);
    localStorage.setItem("authLevel", body.authLevel);
    localStorage.setItem("email", email);
    navigate("/");
  }

  return (
    <>
      <DefaultHeader />
      <main>
        <form
          className="signup-signin-form"
          onSubmit={login}>
          <label htmlFor="email">Email</label>
          <input id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            type="email"
            autoComplete="email" required />

          <label htmlFor="password">Password</label>
          <input id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            type="password"
            required />

          <button className="col-span-full" disabled={loading} type="submit">Log In</button>
          <Link to="/signup" id="signup-link-login-page" className="col-span-full">Sign Up Instead</Link>
        </form>
      </main>
    </>
  );
}
