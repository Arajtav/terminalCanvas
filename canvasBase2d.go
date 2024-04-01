package terminalCanvas

import (
	"math"
	"sort"
)

// Draws line using Bresenham's line algorithm
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
        c.setPixelUnsafe(U16Vec2{uint16(cp.X), uint16(cp.Y)}, col, math.MaxFloat32);
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

// DrawLine but {0, 0} is center of canvas
func (c *Canvas) DrawLineC(a I16Vec2, b I16Vec2, col Color) {
    c.DrawLine(cvPosCenter(a, c.sizeX, c.sizeY), cvPosCenter(b, c.sizeX, c.sizeY), col);
}

// Returns points making line, using Bresenham's line algorithm
func getLine(a I16Vec2, b I16Vec2) []I16Vec2 {
    d := I16Vec2{b.X - a.X, b.Y - a.Y};
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
    cp := a;

    var points []I16Vec2;

    for {
        points = append(points, cp);
        if cp.X == b.X && cp.Y == b.Y { break; }

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
    return points;
}

// Draws filled triangle using scan line algorithm
func (c *Canvas) DrawTriangle(p0 U16Vec2, p1 U16Vec2, p2 U16Vec2, col Color) {
    var points []I16Vec2;
    points = append(points, getLine(I16Vec2{int16(p0.X), int16(p0.Y)}, I16Vec2{int16(p1.X), int16(p1.Y)})...);
    points = append(points, getLine(I16Vec2{int16(p1.X), int16(p1.Y)}, I16Vec2{int16(p2.X), int16(p2.Y)})...);
    points = append(points, getLine(I16Vec2{int16(p2.X), int16(p2.Y)}, I16Vec2{int16(p0.X), int16(p0.Y)})...);

    sort.Slice(points, func(i int, j int) bool {
        return points[i].Y < points[j].Y || (points[i].Y == points[j].Y && points[i].X <= points[j].X);
    });

    i := 0;
    for i<len(points) {
        startp := points[i];
        j := i;
        for points[j].Y == points[i].Y {    // finding last position with same y
            j++;
            if j >= len(points) { break; }
        }
        endp := points[j-1];
        c.DrawLine(U16Vec2{uint16(startp.X), uint16(startp.Y)}, U16Vec2{uint16(endp.X), uint16(endp.Y)}, col);
        i = j;
    }
}

// Same as DrawTriangle but {0, 0} is center of canvas
func (c *Canvas) DrawTriangleC(p0 I16Vec2, p1 I16Vec2, p2 I16Vec2, col Color) {
    c.DrawTriangle(cvPosCenter(p0, c.sizeX, c.sizeY), cvPosCenter(p1, c.sizeX, c.sizeY), cvPosCenter(p2, c.sizeX, c.sizeY), col);
}

