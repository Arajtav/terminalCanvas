package terminalCanvas

import (
    "fmt"
)

type Canvas struct {
    sizeX   int16;
    sizeY   int16;
    data    [][]Color;
}

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

func (c *Canvas) Fill(col Color) {
    for i := int16(0); i<c.sizeY; i++ {
        for j := int16(0); j<c.sizeX; j++ {
            c.data[i][j] = col;
        }
    }
}

func (c *Canvas) Print() {
    for i := int16(0); i<c.sizeY-1; i+=2 {
        for j := int16(0); j<c.sizeX; j++ {
            fmt.Printf("\033[48;2;%d;%d;%dm\033[38;2;%d;%d;%dm▄", c.data[i][j].R, c.data[i][j].G, c.data[i][j].B, c.data[i+1][j].R, c.data[i+1][j].G, c.data[i+1][j].B);
        }
        fmt.Println("\033[0m");
    }
}

func (c *Canvas) SetPixel(pos U16Vec2, col Color) {
    if pos.X >= uint16(c.sizeX) || pos.Y >= uint16(c.sizeY) { return; }
    c.data[pos.Y][pos.X] = col;
}

func cvPosCenter(p I16Vec2, sx int16, sy int16) U16Vec2 { return U16Vec2{uint16(p.X+(sx/2)), uint16(p.Y+(sy/2))}; }

func (c *Canvas) SetPixelC(pos I16Vec2, col Color) {
    c.SetPixel(cvPosCenter(pos, c.sizeX, c.sizeY), col);
}

// Bresenham's line algorithm
func (c *Canvas) DrawLine(a U16Vec2, b U16Vec2, col Color) {
    na := I16Vec2{int16(a.X), int16(a.Y)};
    nb := I16Vec2{int16(b.X), int16(b.Y)};
    d := I16Vec2{nb.X - na.X, nb.Y - na.Y};
    g := I16Vec2{1, 1};

    if d.X < 0 {
        d.X = -d.X;
        g.X = -1;
    }
    if d.Y < 0 {
        d.Y = -d.Y;
        g.Y = -1;
    }

    err := d.X - d.Y;
    cp := na;

    for {
        // TODO: ALIGN POINTS TO CANVAS EDGE FIRST INSTEAD DOING THAT CHECK
        if cp.X >= int16(c.sizeX) || cp.Y >= int16(c.sizeY) || cp.X < 0 || cp.Y < 0 { break; }
        c.data[cp.Y][cp.X] = col;
        if cp.X == nb.X && cp.Y == nb.Y { break; }

        e2 := 2 * err;
        if e2 > -d.Y {
            err -= d.Y;
            cp.X += g.X;
        }
        if e2 < d.X {
            err += d.X;
            cp.Y += g.Y;
        }
    }
}

func (c *Canvas) DrawLineC(a I16Vec2, b I16Vec2, col Color) {
    c.DrawLine(cvPosCenter(a, c.sizeX, c.sizeY), cvPosCenter(b, c.sizeX, c.sizeY), col);
}
