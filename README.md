# sandbox-go-myapp
golangのアプリケーションプロジェクト作成の練習です


## Project Making
```
gitignore code node go gitbook windows macos
yarn init
touch LICENSE
yarn add -D gitbook-cli
vim package.json
  "scripts": {
    "doc:install": "gitbook install ./doc",
    "doc": "gitbook serve ./doc --open",
    "pdf": "gitbook pdf ./doc ./doc/book.pdf"
  },
mkdir -p ./doc/styles
vim book.json
  /*
  body {
    background: no-repeat right -20px bottom 0px / 75% url('../confidential.png');
  }
  */
  
  .page .section.toc h1:before {
    display: none;
  }
  
  .page .section.toc ol a:before {
    display: inline;
    content: counter(toc) ". ";
    padding-right: 5px;
    counter-increment: toc;
  }
  .page .section.toc ol {
    counter-reset: toc;
  }
touch ./doc/README.md
yarn doc:install
```