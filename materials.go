package terminalCanvas

// TODO: RENAME THAT PROBABLY
type Material interface {
    GetColor(F32Vec2) Color;
}

// Material that will always return one color
type MaterialFlat struct { C Color; }
func (m MaterialFlat) GetColor(coord F32Vec2) Color { return m.C; }

// Material that will return UV as RGB
type MaterialUV struct {};
func (M MaterialUV) GetColor(coord F32Vec2) Color {
    return Color{uint8(coord.X*255), uint8(coord.Y*255), 0};
}

type MaterialUVtest struct {};
func (M MaterialUVtest) GetColor(coord F32Vec2) Color {
    if coord.X > 1.0 || coord.X < 0.0 || coord.Y > 1.0 || coord.Y < 0.0 {
        return Color{255, 0, 0};
    }
    return Color{0, 255, 0};
}
