package terminalCanvas

import (
    "fmt"
    "math"
)

// Canvas object, stores pixels and stuff
type Canvas struct {
    sizeX   int16;
    sizeY   int16;
    data    [][]Pixel;
    Ccolor  Color;
}

// Creates new canvas with given size
func NewCanvas(sx int16, sy int16) *Canvas {
    var c Canvas;

    c.sizeX = sx;
    c.sizeY = sy;

    c.data = make([][]Pixel, sy);
    for i := int16(0); i<sy; i++ {
        c.data[i] = make([]Pixel, sx);
    }
    return &c;
}

// Sets every pixel in canvas to given color
func (c *Canvas) Fill(col Color) {
    for i := int16(0); i<c.sizeY; i++ {
        for j := int16(0); j<c.sizeX; j++ {
            c.data[i][j].C = col;
            c.data[i][j].Z = math.MaxFloat32;
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
            fmt.Printf("\033[48;2;%d;%d;%dm\033[38;2;%d;%d;%dm▄", c.data[i][j].C.R, c.data[i][j].C.G, c.data[i][j].C.B, c.data[i+1][j].C.R, c.data[i+1][j].C.G, c.data[i+1][j].C.B);
        }
        fmt.Println("\033[0m");
    }
}
// Prints Z buffer, larger Z is brighter
func (c *Canvas) PrintZ() {
    var min float32 = math.MaxFloat32;
    var max float32 = -math.MaxFloat32;
    for i := int16(0); i<c.sizeY-1; i+=2 {
        for j := int16(0); j<c.sizeX; j++ {
            if c.data[i][j].Z < min && c.data[i][j].Z != -math.MaxFloat32 { min = c.data[i][j].Z };
            if c.data[i][j].Z > max && c.data[i][j].Z !=  math.MaxFloat32 { max = c.data[i][j].Z };
        }
    }
    var d float32 = (max-min)/255.0;
    for i := int16(0); i<c.sizeY-1; i+=2 {
        for j := int16(0); j<c.sizeX; j++ {
            var pt uint8 = 255 - uint8((c.data[i][j].Z-min)/d);
            var pb uint8 = 255 - uint8((c.data[i+1][j].Z-min)/d);
            fmt.Printf("\033[48;2;%d;%d;%dm\033[38;2;%d;%d;%dm▄", pt, pt, pt, pb, pb, pb);
        }
        fmt.Println("\033[0m");
    }

}

// Sets pixel at given position to given color
func (c *Canvas) SetPixel(pos U16Vec2, col Color) {
    if pos.X >= uint16(c.sizeX) || pos.Y >= uint16(c.sizeY) { return; }
    c.data[pos.Y][pos.X].C = col;
    c.data[pos.Y][pos.X].Z = math.MaxFloat32;
}

// Convert coordinates so that {0, 0} is top left corner instead center of canvas
func cvPosCenter(p I16Vec2, sx int16, sy int16) U16Vec2 { return U16Vec2{uint16(p.X+(sx/2)), uint16(p.Y+(sy/2))}; }

// Sets pixel but {0, 0} is center of canvas
func (c *Canvas) SetPixelC(pos I16Vec2, col Color) {
    c.SetPixel(cvPosCenter(pos, c.sizeX, c.sizeY), col);
}
