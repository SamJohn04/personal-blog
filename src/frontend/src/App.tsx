import './App.css'
import { BrowserRouter, Route, Routes } from "react-router"
import Home from "./pages/Home"
import Blog from "./pages/Blog"
import Signup from './pages/Signup'
import Login from './pages/Login'
import { CreateBlog } from './pages/CreateBlog'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/blog/:id" element={<Blog />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/blog/new" element={<CreateBlog />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
