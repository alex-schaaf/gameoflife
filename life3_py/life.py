import typer
import numpy as np
import scipy.ndimage
from tqdm import tqdm
from datetime import datetime
import os
from PIL import Image


class Life:
    def __init__(self, nx: int, ny: int, seed: int, density: float = 0.5):
        self.nx = nx
        self.ny = ny
        self.seed = seed

        self.density = density
        self.world = np.random.binomial(1, self.density, size=(nx, ny))

        self.conv_edge_mode = "wrap"

    def evolve(self):
        # this awesome implementation comes from
        # http://greenteapress.com/complexity/html/thinkcomplexity008.html
        kernel = np.array([[1, 1, 1],
                           [1,10, 1],
                           [1, 1, 1]])

        neighbors = scipy.ndimage.filters.convolve(
            self.world, kernel, mode=self.conv_edge_mode
        )

        boolean = (neighbors == 3) | (neighbors == 12) | (neighbors == 13)
        self.world = np.int8(boolean)

    def zoom(self, zoom: int):
        return np.kron(self.world, np.ones((zoom, zoom))).astype(np.uint8)

    def write_pillow(self, fileapth: str, zoom: int = 5):
        arr = self.zoom(zoom)
        img = Image.fromarray(arr * 255)
        img.save(fileapth)


app = typer.Typer()


@app.command()
def simulate(
        nx: int,
        ny: int,
        seed: int,
        iterations: int,
        density: float = 0.5,
        zoom: int = 5,
        animate: bool = False,
        format: str = "jpg",
        debug: bool = False,
):
    if debug:
        for k, v in locals().items():
            msg = typer.style(k, fg="white") + " " + typer.style(str(v), fg="red")
            typer.echo(msg)

    folder = str(datetime.now()).replace(":", "-").split(".")[0].replace(" ", "-")
    os.mkdir(folder)
    life = Life(nx, ny, seed, density=density)

    typer.secho("Evolving life..", fg="green")

    n_digits = len(str(iterations))
    for i in tqdm(range(iterations)):
        life.evolve()
        life.write_pillow(f"{folder}/{str(i).zfill(n_digits)}.{format}", zoom=zoom)

    if animate:
        typer.echo("Animating image...")
        os.system(f'magick convert ./{folder}/*.{format} ./{folder}/animation.gif')
        typer.echo("All done.")


if __name__ == "__main__":
    app()
