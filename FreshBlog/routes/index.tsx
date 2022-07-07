/** @jsx h */
import { asset } from "$fresh/src/runtime/utils.ts";
import { h } from "preact";
import Footer from "../islands/Footer.tsx";
import Navbar from "../islands/Navbar.tsx";
import ArticlePreview from "../islands/ArticlePreview.tsx";

export default function Home() {
  return (
    <html>
      <head>
        <title>Fresh framework demo</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link
          href={asset("/bootstrap/css/bootstrap.min.css")}
          rel="stylesheet"
        />
        <link
          href={asset("/index.css")}
          rel="stylesheet"
        />
      </head>
      <body className="d-flex flex-column h-100">
        <Navbar />
        <main className="flex-shrink-0">
          <div className="container">
            <h1 className="mt-5">Sticky footer with fixed navbar</h1>
            <div className="lead">
              <ArticlePreview title={"Test"} />
            </div>
          </div>
        </main>
        <Footer />
        <script src={asset("/bootstrap/js/bootstrap.bundle.min.js")} />
      </body>
    </html>
  );
}
