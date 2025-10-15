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
      const data = await res.json();
      setBlogTitles(data.map((blogTitle: any) => ({
        id: blogTitle.id,
        title: blogTitle.title,
        createdAt: new Date(blogTitle.createdAt),
        lastEditedAt: new Date(blogTitle.lastEditedAt),
      })));
    }
    getBlogTitles()
  });

  return (
    <>
      <DefaultHeader />
      {
        blogTitles.map(blogTitle => <Link key={blogTitle.id} className="index-item pad-y-1" to={`/blog/${blogTitle.id}`}>
                       <span>{blogTitle.title}</span>
                       <span>{blogTitle.createdAt.toDateString()}</span>
                       </Link>)
      }
    </>
  );
}
