package terminalCanvas

import (
    "image"
)

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
    return Color{uint8(coord.X*255), uint8(coord.Y*255), 0, 255};
}

// Material to check if UV range is correct
// Green - correct (UV X and Y between 0.0 and 1.0)
// Red   - wrong
type MaterialUVtest struct {};
func (M MaterialUVtest) GetColor(coord F32Vec2) Color {
    if coord.X > 1.0 || coord.X < 0.0 || coord.Y > 1.0 || coord.Y < 0.0 {
        return Color{255, 0, 0, 255};
    }
    return Color{0, 255, 0, 255};
}

type MaterialImage struct {Img image.Image};
func (M MaterialImage) GetColor(coord F32Vec2) Color {
    r, g, b, a := M.Img.At( int(float32(M.Img.Bounds().Dx()-1)*coord.X),
                            int(float32(M.Img.Bounds().Dy()-1)*coord.Y)).RGBA();
    return Color{uint8(r/256), uint8(g/256), uint8(b/256), uint8(a/256)};
}

type MaterialPseudoRandom struct {w uint32};
func (M *MaterialPseudoRandom) GetColor(coord F32Vec2) Color {
    M.w = M.w*98765 + 12345678;
    tmp := uint8(M.w);
    return Color{R: tmp, G: tmp, B: tmp, A: 255};
}
