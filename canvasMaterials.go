package terminalCanvas

import (
    "sort"
)

func interpolateUV(a F32Vec2, b F32Vec2, t float32) F32Vec2 {
    return F32Vec2{a.X + ((b.X - a.X)*t), a.Y + ((b.Y - a.Y)*t)}
}

func (c *Canvas) DrawLineF(a U16Frag, b U16Frag, f Material, normal F32Vec3) {
    na := I16Vec2{int16(a.Pos.X), int16(a.Pos.Y)}
    nb := I16Vec2{int16(b.Pos.X), int16(b.Pos.Y)}
    d := I16Vec2{nb.X - na.X, nb.Y - na.Y}
    g := I16Vec2{1, 1}

    if d.X < 0 {
        d.X = -d.X
        g.X = -1
    }
    if d.Y < 0 {
        d.Y = -d.Y
        g.Y = -1
    }

    err := d.X - d.Y
    cp := na

    tl := distI16Vec2(I16Vec2{na.X, na.Y}, I16Vec2{nb.X, nb.Y})
    if tl < 1 {
        c.SetPixel(U16Vec2{a.Pos.X, a.Pos.Y}, f.GetColor(interpolateUV(a.UV, b.UV, 0.5), normal))
        return
    }

    for {
        // TODO: ALIGN POINTS TO CANVAS EDGE FIRST INSTEAD DOING THAT CHECK
        if cp.X >= int16(c.sizeX) || cp.Y >= int16(c.sizeY) || cp.X < 0 || cp.Y < 0 { break }
        t := float32(distI16Vec2(I16Vec2{na.X, na.Y}, cp)/tl)
        if c.data[cp.Y][cp.X].Z >= (a.Z+((b.Z-a.Z)*t)) {
            c.setPixelUnsafe(U16Vec2{uint16(cp.X), uint16(cp.Y)}, f.GetColor(interpolateUV(a.UV, b.UV, t), normal), (a.Z+((b.Z-a.Z)*t)))
        }
        if cp.X == nb.X && cp.Y == nb.Y { break }

        e2 := 2 * err
        if e2 > -d.Y {
            err -= d.Y
            cp.X += g.X
        }
        if e2 < d.X {
            err += d.X
            cp.Y += g.Y
        }
    }
}

func (c *Canvas) DrawLineFC(a I16Frag, b I16Frag, f Material, normal F32Vec3) {
    c.DrawLineF(U16Frag{cvPosCenter(a.Pos, c.sizeX, c.sizeY), a.UV, a.Z}, U16Frag{cvPosCenter(b.Pos, c.sizeX, c.sizeY), b.UV, b.Z}, f, normal)
}

func getLineF(a U16Frag, b U16Frag) []U16Frag {
    d := I16Vec2{int16(b.Pos.X) - int16(a.Pos.X), int16(b.Pos.Y) - int16(a.Pos.Y)}
    g := I16Vec2{1, 1}

    if d.X < 0 {
        d.X = -d.X
        g.X = -1
    }
    if d.Y < 0 {
        d.Y = -d.Y
        g.Y = -1
    }

    err := d.X - d.Y
    cp := I16Vec2{int16(a.Pos.X), int16(a.Pos.Y)}

    var points []U16Frag

    tl := distI16Vec2(I16Vec2{int16(a.Pos.X), int16(a.Pos.Y)}, I16Vec2{int16(b.Pos.X), int16(b.Pos.Y)})
    if tl < 1 {
        points = append(points, U16Frag{a.Pos, interpolateUV(a.UV, b.UV, 0.5), a.Z+((b.Z-a.Z)*0.5)})
        return points
    }

    for {
        t := float32(distI16Vec2(I16Vec2{int16(a.Pos.X), int16(a.Pos.Y)}, cp)/tl)
        points = append(points, U16Frag{U16Vec2{uint16(cp.X), uint16(cp.Y)}, interpolateUV(a.UV, b.UV, t), a.Z + ((b.Z-a.Z)*t)})
        if cp.X == int16(b.Pos.X) && cp.Y == int16(b.Pos.Y) { break }

        e2 := 2 * err
        if e2 > -d.Y {
            err -= d.Y
            cp.X += g.X
        }
        if e2 < d.X {
            err += d.X
            cp.Y += g.Y
        }
    }
    return points
}

func (c *Canvas) DrawTriangleF(p0 U16Frag, p1 U16Frag, p2 U16Frag, f Material, normal F32Vec3) {
    var points []U16Frag
    points = append(points, getLineF(p0, p1)...)
    points = append(points, getLineF(p1, p2)...)
    points = append(points, getLineF(p2, p0)...)

    sort.Slice(points, func(i int, j int) bool {
        return points[i].Pos.Y < points[j].Pos.Y || (points[i].Pos.Y == points[j].Pos.Y && points[i].Pos.X <= points[j].Pos.X)
    })

    i := 0
    for i<len(points) {
        startp := points[i]
        j := i
        for points[j].Pos.Y == points[i].Pos.Y {    // finding last position with same y
            j++
            if j >= len(points) { break }
        }
        endp := points[j-1]
        c.DrawLineF(startp, endp, f, normal)
        i = j
    }
}

func (c *Canvas) DrawTriangleFC(p0 I16Frag, p1 I16Frag, p2 I16Frag, f Material, normal F32Vec3) {
    c.DrawTriangleF(U16Frag{cvPosCenter(p0.Pos, c.sizeX, c.sizeY), p0.UV, p0.Z},
                    U16Frag{cvPosCenter(p1.Pos, c.sizeX, c.sizeY), p1.UV, p1.Z},
                    U16Frag{cvPosCenter(p2.Pos, c.sizeX, c.sizeY), p2.UV, p2.Z},
                    f, normal)
}
