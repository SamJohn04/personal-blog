import './App.css'
import { BrowserRouter, Route, Routes } from "react-router"
import Home from "./pages/Home"
import Blog from "./pages/Blog"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/blog/:id" element={<Blog />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
