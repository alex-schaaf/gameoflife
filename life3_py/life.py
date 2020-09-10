import typer
import numpy as np
import scipy.ndimage
from tqdm import tqdm
from datetime import datetime
import os
from PIL import Image
from matplotlib.pyplot import get_cmap
from abc import ABC, abstractmethod


class GameOfLife(ABC):
    def __init__(self, nx: int, ny: int, seed: int):
        self.nx = nx
        self.ny = ny
        self.seed = seed
        self.world = None

    @abstractmethod
    def seed_world(self):
        pass

    @abstractmethod
    def evolve(self):
        pass


class Life1(GameOfLife):
    def seed_world(self, density: float = 0.5):
        self.world = np.random.binomial(1, density, size=(self.nx, self.ny))

    def evolve(self, conv_edge_mode: str = "wrap"):
        # this awesome implementation comes from
        # http://greenteapress.com/complexity/html/thinkcomplexity008.html
        kernel = np.array([[1, 1, 1],
                           [1, 10, 1],
                           [1, 1, 1]])

        neighbors = scipy.ndimage.filters.convolve(
            self.world, kernel, mode=conv_edge_mode
        )

        boolean = (neighbors == 3) | (neighbors == 12) | (neighbors == 13)
        self.world = np.int8(boolean)


# class Life:
#     def __init__(self, nx: int, ny: int, seed: int, density: float = 0.5):
#         self.nx = nx
#         self.ny = ny
#         self.seed = seed
#
#         self.density = density
#         self.world = np.random.binomial(1, self.density, size=(nx, ny))
#         self.worldc = np.zeros_like(self.world)
#
#         self.conv_edge_mode = "reflect"
#         self.cmap = get_cmap("viridis_r")
#
#     def evolve(self):
#         # this awesome implementation comes from
#         # http://greenteapress.com/complexity/html/thinkcomplexity008.html
#         kernel = np.array([[1, 1, 1],
#                            [1,10, 1],
#                            [1, 1, 1]])
#
#         neighbors = scipy.ndimage.filters.convolve(
#             self.world, kernel, mode=self.conv_edge_mode
#         )
#
#         boolean = (neighbors == 3) | (neighbors == 12) | (neighbors == 13)
#         self.world = np.int8(boolean)
#
#     def evolve2(self):
#         kernel = np.array([[1, 1, 1],
#                            [1,10, 1],
#                            [1, 1, 1]])
#
#         neighbors = scipy.ndimage.filters.convolve(
#             self.world, kernel, mode=self.conv_edge_mode
#         )
#
#         boolean = (neighbors == 3) | (neighbors == 12) | (neighbors == 13)
#         self.world = np.int8(boolean)
#
#         self.worldc[boolean] = 255
#         self.worldc[~boolean] -= 31
#         self.worldc[self.worldc <= 0] = 0


def enlarge(arr: np.array, zoom: int) -> np.ndarray:
    """Enlarge the world array by given zoom level."""
    return np.kron(arr, np.ones([zoom for _ in range(arr.ndim)])).astype(np.uint8)


def write_pillow(arr: np.ndarray, filepath: str):
    """Writes current world array to given filepath using Pillow."""
    img = Image.fromarray(arr * 255)
    img.save(filepath)


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
        filetype: str = "jpg",
        debug: bool = False,
):
    if debug:
        for k, v in locals().items():
            msg = typer.style(k, fg="white") + " " + typer.style(str(v), fg="red")
            typer.echo(msg)

    folder = str(datetime.now()).replace(":", "-").split(".")[0].replace(" ", "-")
    os.mkdir(folder)

    life = Life1(nx, ny, seed)
    life.seed_world(density=density)

    typer.secho("Evolving life..", fg="green")
    n_digits = len(str(iterations))
    for i in tqdm(range(iterations)):
        life.evolve()
        arr = enlarge(life.world, zoom=zoom)
        write_pillow(arr, f"{folder}/{str(i).zfill(n_digits)}.{filetype}")

    if animate:
        typer.echo("Animating image...")
        os.system(f'magick convert ./{folder}/*.{filetype} ./{folder}/animation.gif')
        typer.echo("All done.")


if __name__ == "__main__":
    app()
