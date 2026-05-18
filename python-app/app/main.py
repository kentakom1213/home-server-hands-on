from fastapi import FastAPI
from fastapi.responses import HTMLResponse

app = FastAPI()


@app.get("/", response_class=HTMLResponse)
def root():
    return """
<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>FastAPI Server</title>
  </head>
  <body>
    <h1>Hello from FastAPI</h1>
  </body>
</html>
"""
