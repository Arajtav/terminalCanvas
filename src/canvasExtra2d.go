package terminalCanvas

import (
    "math"
    "sort"
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

    tl := distI16Vec2(I16Vec2{na.X, na.Y}, I16Vec2{nb.X, nb.Y});
    if tl < 1 {
        c.SetPixel(U16Vec2{a.X, a.Y}, interpolateColor(a.C, b.C, 0.5));
        return;
    }

    for {
        // TODO: ALIGN POINTS TO CANVAS EDGE FIRST INSTEAD DOING THAT CHECK
        if cp.X >= int16(c.sizeX) || cp.Y >= int16(c.sizeY) || cp.X < 0 || cp.Y < 0 { break; }
        c.data[cp.Y][cp.X] = interpolateColor(a.C, b.C, distI16Vec2(I16Vec2{na.X, na.Y}, cp)/tl);
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

func (c *Canvas) DrawLineIC(a I16Vec2C, b I16Vec2C) {
    na := cvPosCenter(I16Vec2{a.X, a.Y}, c.sizeX, c.sizeY);
    nb := cvPosCenter(I16Vec2{b.X, b.Y}, c.sizeX, c.sizeY);
    c.DrawLineI(U16Vec2C{na.X, na.Y, a.C}, U16Vec2C{nb.X, nb.Y, b.C});
}

func getLineI(a I16Vec2C, b I16Vec2C) []I16Vec2C {
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
    cp := I16Vec2{a.X, a.Y};

    var points []I16Vec2C;

    tl := distI16Vec2(I16Vec2{a.X, a.Y}, I16Vec2{b.X, b.Y});
    if tl < 1 {
        points = append(points, I16Vec2C{a.X, a.Y, interpolateColor(a.C, b.C, 0.5)});
        return points;
    }

    for {
        points = append(points, I16Vec2C{cp.X, cp.Y, interpolateColor(a.C, b.C, distI16Vec2(I16Vec2{a.X, b.X}, cp)/tl)});
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

// TODO THIS IS BROKEN
func (c *Canvas) DrawTriangleI(p0 U16Vec2C, p1 U16Vec2C, p2 U16Vec2C) {
    var points []I16Vec2C;
    points = append(points, getLineI(I16Vec2C{int16(p0.X), int16(p0.Y), p0.C}, I16Vec2C{int16(p1.X), int16(p1.Y), p1.C})...);
    points = append(points, getLineI(I16Vec2C{int16(p1.X), int16(p1.Y), p1.C}, I16Vec2C{int16(p2.X), int16(p2.Y), p2.C})...);
    points = append(points, getLineI(I16Vec2C{int16(p2.X), int16(p2.Y), p2.C}, I16Vec2C{int16(p0.X), int16(p0.Y), p0.C})...);

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
        c.DrawLineI(U16Vec2C{uint16(startp.X), uint16(startp.Y), startp.C}, U16Vec2C{uint16(endp.X), uint16(endp.Y), endp.C});
        i = j;
    }
}

func (c *Canvas) DrawTriangleIC(p0 I16Vec2C, p1 I16Vec2C, p2 I16Vec2C) {
    p0n := cvPosCenter(I16Vec2{p0.X, p0.Y}, c.sizeX, c.sizeY);
    p1n := cvPosCenter(I16Vec2{p1.X, p1.Y}, c.sizeX, c.sizeY);
    p2n := cvPosCenter(I16Vec2{p2.X, p2.Y}, c.sizeX, c.sizeY);
    c.DrawTriangleI(U16Vec2C{p0n.X, p0n.Y, p0.C}, U16Vec2C{p1n.X, p1n.Y, p1.C}, U16Vec2C{p2n.X, p2n.Y, p2.C});
}
