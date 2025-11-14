import { useState, type FormEvent } from "react";
import DefaultHeader from "../components/DefaultHeader";
import { useNavigate } from "react-router";

export default function Signup() {
  const navigate = useNavigate();

  const [username, setUsername] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const [loading, setLoading] = useState<boolean>(false);

  async function register(event: FormEvent) {
    event.preventDefault();
    setLoading(true);
    const res = await fetch("/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username,
        email,
        password,
      })
    });

    if (res.ok) {
      navigate("/login");
    } else {
      alert("Something went wrong...");
      setLoading(false);
    }
  }

  return (
    <>
      <DefaultHeader />
      <main>
        <form
          className="signup-signin-form"
          onSubmit={register}>
          <label htmlFor="username">Username</label>
          <input id="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            type="text"
            autoComplete="username" required />

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
            type="password" required />

          <button
            className="col-span-full"
            disabled={loading}
            type="submit">Sign Up</button>
        </form>
      </main>
    </>
  )
}
