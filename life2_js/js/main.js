var canvas = document.getElementById('canvas');
let ctx = canvas.getContext("2d");

let width = window.innerWidth
let height = window.innerHeight - 150

canvas.width = width / 2
canvas.height = height

let cellWidth = 3;
let cellHeight = 3;

let nCellsX = width / cellWidth
let nCellsY = height / cellHeight

let settings = {
    width: width,
    height: height,
    cellWidth: cellWidth,
    cellHeight: cellHeight,
    nCellsX: nCellsX,
    nCellsY: Math.floor(nCellsY),
    colorAlive: "#FDBD0D",
    colorDead: "#151515",
}

function arrayClone( arr ) {

    var i, copy;

    if( Array.isArray( arr ) ) {
        copy = arr.slice( 0 );
        for( i = 0; i < copy.length; i++ ) {
            copy[ i ] = arrayClone( copy[ i ] );
        }
        return copy;
    } else if( typeof arr === 'object' ) {
        throw 'Cannot clone array containing an object!';
    } else {
        return arr;
    }

}


console.log(settings)

function seedWorld(settings, density=50) {
    let world = []
    for (let y=0;y<settings.nCellsY;y+=1) {
        let row = []
        for (let x=0;x<settings.nCellsX;x+=1) {
            let rn = Math.random() * 100
            if (rn >= 100 - density) {
                row.push(1)
            } else {
                row.push(0)
            }
        }
        world.push(row)
    }
    return world
}

function renderWorld(world, settings) {
    canvas.width += 0
    ctx.fillStyle = settings.colorAlive;  // mango
    for (var y = 0; y < settings.nCellsY; y += 1) {
        for (var x = 0; x < settings.nCellsX; x += 1) {
            if (world[y][x] === 1) {
                ctx.fillRect(
                    x * settings.cellWidth, 
                    y * settings.cellHeight, 3, 3);
            } 
            
        }
    }
}


function countNeighbors(x, y, world) {
    let n = 0
    try {
        n += world[y][x + 1]
        n += world[y][x - 1]
        n += world[y + 1][x]
        n += world[y - 1][x]
    
        n += world[y + 1][x + 1]
        n += world[y + 1][x - 1]
        n += world[y - 1][x + 1]
        n += world[y - 1][x - 1]
        return n
    } catch (error) {
        return n
    }
}


function mutateWorld(world, settings) {
    let newWorld = []
    for (let y = 0; y < settings.nCellsY; y += 1) {
        let row = []
        for (let x = 0; x < settings.nCellsX; x += 1) {
            let alive = world[y][x];
            let neighbors = countNeighbors(x, y, world, settings);

            if (alive == 1 && neighbors < 2) {
                row.push(0)
            } else if (alive == 1 && neighbors <= 3) {
                row.push(1)
            } else if (alive == 1 && neighbors > 3) {
                row.push(0)
            } else if (alive == 0 && neighbors == 3) {
                row.push(1)
            } else {
                row.push(0)
            }
        }
        newWorld.push(row)
    }
    return newWorld
}



let nIterations = 100

let run = false
let density = 50

let world = seedWorld(settings, density=density);

$(document).ready(function() {
    $("#start").click(function(){
        run = true
    }); 

    $("#stop").click(function(){
        run = false
    }); 

    $("#reset").click(function(){
        world = seedWorld(settings, density=density)
        renderWorld(world, settings)
        run = false
    }); 

    $(document).on("input", "#density", function() {
        density = $(this).val()
        if (run == false) {
            world = seedWorld(settings, density=density)
            renderWorld(world, settings)
        }
        
    })

});

let i = 0;
renderWorld(world, settings)
function gameLoop() {
    setTimeout(function() {
        if (run === true) {
            world = mutateWorld(world, settings)
            renderWorld(world, settings)
            
            i++;
        }
        gameLoop()
    }, 16.33)
}

gameLoop();

