import { useState, type FormEvent } from "react";
import DefaultHeader from "../components/DefaultHeader";

export default function Login() {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  async function login(event: FormEvent) {
    event.preventDefault();
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
    if (res.ok) {
      alert("Success!");
    } else {
      alert("Something went wrong...");
    }
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

          <button className="col-span-full" type="submit">Log In</button>
        </form>
      </main>
    </>
  );
}
