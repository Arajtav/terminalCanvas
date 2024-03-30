package terminalCanvas

// RGB one byte per channel
type Color struct {
    R uint8;
    G uint8;
    B uint8;
}

type U16Vec2 struct {
    X uint16;
    Y uint16;
}

type I16Vec2 struct {
    X int16;
    Y int16;
}

type I16Vec2C struct {
    X int16;
    Y int16;
    C Color;
}

type U16Vec2C struct {
    X uint16;
    Y uint16;
    C Color;
}

type F32Vec2 struct {
    X float32;
    Y float32;
}

// Fragment data
type U16Frag struct {
    Pos U16Vec2;
    UV  F32Vec2;
}

type I16Frag struct {
    Pos I16Vec2;
    UV  F32Vec2;
}
