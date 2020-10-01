"""
https://stackoverflow.com/questions/59773795/generating-rows-of-a-rule-30-cellular-automaton
"""
import matplotlib.pyplot as plt
import typer


app = typer.Typer()


def rule30(t):
    return t[0] ^ (t[1] or t[2])


def get_initial_state(width: int) -> list:
    initial_state = [0] * width
    if width % 2:
        initial_state[width // 2] = 1
    else:
        initial_state.insert(width // 2, 1)
    return initial_state


def get_triples(line):
    return zip(line, line[1:], line[2:])


def apply(rule, line: list) -> list:
    return [rule(triple) for triple in get_triples([0] + line + [0])]  # padding


@app.command()
def simulate(
        width: int = 200,
):
    state = [get_initial_state(width)]
    while not state[-1][0]:
        state.append(apply(rule30, state[-1]))

    plt.imshow(state)
    plt.show()


if __name__ == '__main__':
    app()
