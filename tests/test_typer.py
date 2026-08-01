import typer
app = typer.Typer()
@app.command()
def start() -> None:
    print("hello")
if __name__ == '__main__':
    app()
