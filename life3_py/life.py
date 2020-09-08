import typer
import numpy as np
import scipy.ndimage
from tqdm import tqdm
from datetime import datetime
import os


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

    def write_ppm(self, filepath: str, zoom: int = 5):
        with open(filepath, "w") as file:
            file.write("P1\n")
            file.write("#\n")
            file.write(f'{self.nx * zoom} {self.ny * zoom}\n')
            # file.write('255\n')  # max value

            img = np.kron(self.world, np.ones((zoom, zoom))).astype(np.uint8)
            for row in img:
                file.write(str(row)[1:-1] + "\n")


app = typer.Typer()


@app.command()
def simulate(
        nx: int,
        ny: int,
        seed: int,
        iterations: int,
        density: float = 0.5,
        zoom: int = 5,
        name: str = None,
):
    folder = name if name else str(datetime.now())
    os.mkdir(folder)
    life = Life(nx, ny, seed, density=density)

    for i in tqdm(range(iterations)):
        life.evolve()
        life.write_ppm(f"{folder}/{i}.ppm", zoom=zoom)


if __name__ == "__main__":
    app()
