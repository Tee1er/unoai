# unoai

An Uno game written in Go using the Wails framework.

**Todo**
- Working game + UI
- Computer vs human(s) gameplay
- Computer vs computer(s) mode

## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `npm run dev`. The frontend dev server will run on http://localhost:34115. Connect to this in your
browser and connect to your application.

## Building

To build a redistributable, production mode package, use `wails build`.