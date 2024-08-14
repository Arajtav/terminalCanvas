package terminalCanvas

import (
    "image"
)

type Material interface {
    GetColor(F32Vec2, F32Vec3) Color    // UV, Normal
}

// Material that will always return one color
type MaterialFlat struct { C Color }
func (M MaterialFlat) GetColor(coord F32Vec2, normal F32Vec3) Color { return M.C }

// Material that will return UV as RGB
type MaterialUV struct {}
func (M MaterialUV) GetColor(coord F32Vec2, normal F32Vec3) Color {
    return Color{uint8(coord.X*255), uint8(coord.Y*255), 0, 255}
}

// pi -> 255
// Material that will return color based on normals
type MaterialNormal struct {}
func (M MaterialNormal) GetColor(coord F32Vec2, normal F32Vec3) Color {
    normal.X = max(min(normal.X, 1), 0)
    normal.Y = max(min(normal.Y, 1), 0)
    normal.Z = max(min(normal.Z, 1), 0)
    return Color{R: uint8(normal.X*255.0), G: uint8(normal.Y*255.0), B: uint8(normal.Z*255.0), A: 255}
}

type MaterialImage struct {Img image.Image}
func (M MaterialImage) GetColor(coord F32Vec2, normal F32Vec3) Color {
    r, g, b, a := M.Img.At( int(float32(M.Img.Bounds().Dx()-1)*coord.X),
                            int(float32(M.Img.Bounds().Dy()-1)*coord.Y)).RGBA()
    return Color{uint8(r/256), uint8(g/256), uint8(b/256), uint8(a/256)}
}

type MaterialPseudoRandom struct {w uint32; u uint32}
func (M *MaterialPseudoRandom) GetColor(coord F32Vec2, normal F32Vec3) Color {
    // patterns are always there anyway, idk
    M.w = 15 + M.w*77
    M.u += 13 + M.u*M.w*M.u
    tmp := uint8(M.w+M.u+M.w)
    return Color{R: tmp, G: tmp, B: tmp, A: 255}
}

type MaterialPseudoRandomRGB struct {w uint32; u uint32}
func (M *MaterialPseudoRandomRGB) GetColor(coord F32Vec2, normal F32Vec3) Color {
    M.w = 15 + M.w*77
    M.u += 13 + M.u*M.w*M.u
    r := uint8(M.w+M.u+M.w)
    M.w = 13 + M.w*77
    M.u += 11 + M.u*M.w*M.u
    g := uint8(M.w+M.u+M.w)
    M.w = 11 + M.w*77
    M.u += 7 + M.u*M.w*M.u
    return Color{R: r, G: g, B: uint8(M.w+M.u+M.w), A: 255}
}
