package terminalCanvas

import (
    "math"
)

func interpolateColor(a Color, b Color, t float64) Color {
    var c Color;
    c.R = uint8(float64(a.R) + ((float64(b.R)-float64(a.R))*t));
    c.G = uint8(float64(a.G) + ((float64(b.G)-float64(a.G))*t));
    c.B = uint8(float64(a.B) + ((float64(b.B)-float64(a.B))*t));
    return c;
}

func distI16Vec2(a I16Vec2, b I16Vec2) float64 {
    return math.Sqrt(math.Pow(float64(a.X-b.X), 2.0) + math.Pow(float64(a.Y-b.Y), 2.0));
}

func (c *Canvas) DrawLineI(a U16Vec2C, b U16Vec2C) {
    na := I16Vec2C{int16(a.X), int16(a.Y), a.C};
    nb := I16Vec2C{int16(b.X), int16(b.Y), b.C};
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
    cp := I16Vec2{na.X, na.Y};

    for {
        // TODO: ALIGN POINTS TO CANVAS EDGE FIRST INSTEAD DOING THAT CHECK
        if cp.X >= int16(c.sizeX) || cp.Y >= int16(c.sizeY) || cp.X < 0 || cp.Y < 0 { break; }
        c.data[cp.Y][cp.X] = interpolateColor(a.C, b.C, distI16Vec2(I16Vec2{na.X, na.Y}, cp)/distI16Vec2(I16Vec2{na.X, na.Y}, I16Vec2{nb.X, nb.Y}));
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
