import { useEffect, useState } from "react";
import DefaultHeader from "../components/DefaultHeader";
import { Link } from "react-router";

type BlogTitle = {
  id: number,
  title: string,
  createdAt: Date,
  lastEditedAt: Date
}

export default function Home() {
  const [blogTitles, setBlogTitles] = useState<BlogTitle[]>([]);
  useEffect(() => {
    async function getBlogTitles() {
      const res = await fetch("/api/blogs");
      if (!res.ok) {
        return;
      }
      const data = await res.json();
      setBlogTitles(data.map((blogTitle: any) => ({
        id: blogTitle.id,
        title: blogTitle.title,
        createdAt: new Date(blogTitle.createdAt),
        lastEditedAt: new Date(blogTitle.lastEditedAt),
      })));
    }
    getBlogTitles();
  }, []);

  return (
    <>
      <DefaultHeader />
      <main className="main">
      {
        blogTitles.map(blogTitle => <span key={blogTitle.id} className="index-item pad-y-1">
                       <b>
                        <Link to={`/blog/${blogTitle.id}`}>{blogTitle.title}</Link>
                       </b>
                       <span>{blogTitle.createdAt.toDateString()}</span>
                       </span>)
      }
      </main>
    </>
  );
}
