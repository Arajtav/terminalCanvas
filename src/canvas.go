package terminalCanvas

import (
    "fmt"
)

// Canvas object, stores pixels and stuff
type Canvas struct {
    sizeX   int16;
    sizeY   int16;
    data    [][]Color;
    Ccolor  Color;
}

// Creates new canvas with given size
func NewCanvas(sx int16, sy int16) *Canvas {
    var c Canvas;

    c.sizeX = sx;
    c.sizeY = sy;

    c.data = make([][]Color, sy);
    for i := int16(0); i<sy; i++ {
        c.data[i] = make([]Color, sx);
    }
    return &c;
}

// Sets every pixel in canvas to given color
func (c *Canvas) Fill(col Color) {
    for i := int16(0); i<c.sizeY; i++ {
        for j := int16(0); j<c.sizeX; j++ {
            c.data[i][j] = col;
        }
    }
}

// Fills canvas with color stored in Ccolor
func (c *Canvas) Clear() {
    c.Fill(c.Ccolor);
}

// Prints canvas to terminal, using escape sequences for colors and `▄` for smaller pixels
func (c *Canvas) Print() {
    for i := int16(0); i<c.sizeY-1; i+=2 {
        for j := int16(0); j<c.sizeX; j++ {
            fmt.Printf("\033[48;2;%d;%d;%dm\033[38;2;%d;%d;%dm▄", c.data[i][j].R, c.data[i][j].G, c.data[i][j].B, c.data[i+1][j].R, c.data[i+1][j].G, c.data[i+1][j].B);
        }
        fmt.Println("\033[0m");
    }
}
// Sets pixel at given position to given color
func (c *Canvas) SetPixel(pos U16Vec2, col Color) {
    if pos.X >= uint16(c.sizeX) || pos.Y >= uint16(c.sizeY) { return; }
    c.data[pos.Y][pos.X] = col;
}

// Convert coordinates so that {0, 0} is top left corner instead center of canvas
func cvPosCenter(p I16Vec2, sx int16, sy int16) U16Vec2 { return U16Vec2{uint16(p.X+(sx/2)), uint16(p.Y+(sy/2))}; }

// Sets pixel but {0, 0} is center of canvas
func (c *Canvas) SetPixelC(pos I16Vec2, col Color) {
    c.SetPixel(cvPosCenter(pos, c.sizeX, c.sizeY), col);
}
