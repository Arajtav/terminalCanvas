package terminalCanvas

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
