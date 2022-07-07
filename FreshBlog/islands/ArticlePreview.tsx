/** @jsx h */
import { h } from "preact";
import { useState } from "preact/hooks";
import { IS_BROWSER } from "$fresh/runtime.ts";

interface PreviewProps {
  titleImgId: string;
  title: string;
  text: string;
  tags: string[];
  date: Date;
  link: string;
}

const ArticlePreview = (props: PreviewProps) => {
  return (
    <div className="card mb-2">
      <a href={props.link}>
        <img
          className="card-img-top img-fluid rounded-b-none"
          src="https://images.unsplash.com/photo-1579353977828-2a4eab540b9a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1074&q=80"
          alt={props.title}
        />
        <div className="card-body">
          <h3 className="card-title">{props.title}</h3>
          <p className="card-text">{props.text}</p>
          <p className="card-text">
            <small className="text-muted">{props.date}</small>
          </p>
        </div>
      </a>
    </div>
  );
};

export default ArticlePreview;
