const std = @import("std");
const print = std.debug.print;

const Point = struct {
    sIdx: usize,
    eIdx: usize,
};

fn cleanPath(p: []const u8) Point {
    var sIdx: usize = 0;
    var eIdx: usize = p.len;

    while (p[eIdx - 1] == '/') : (eIdx -= 1) {}

    while (p[sIdx] == '/' or p[sIdx] == '.') : (sIdx += 1) {}

    return .{
        .sIdx = sIdx,
        .eIdx = eIdx,
    };
}

fn dir(p: []const u8) []const u8 {
    if (p.len == 0) {
        return "";
    }

    const point = cleanPath(p);

    var tail_is_file = false;
    var i: usize = point.eIdx - 1;
    while (i > point.sIdx) : (i -= 1) {
        if (p[i] == '.') {
            tail_is_file = true;
        }
        if (tail_is_file and p[i] == '/') {
            return p[point.sIdx..i];
        }
    }

    if (tail_is_file) {
        return "";
    }

    return p[point.sIdx..point.eIdx];
}

fn base(p: []const u8) []const u8 {
    if (p.len == 0) {
        return "";
    }

    const point = cleanPath(p);

    var tail_is_file = false;
    var i: usize = point.eIdx - 1;
    while (i > point.sIdx) : (i -= 1) {
        if (p[i] == '.') {
            tail_is_file = true;
        }
        if (p[i] == '/') {
            if (tail_is_file) {
                return p[(i + 1)..point.eIdx];
            }
            return "";
        }
    }

    if (tail_is_file) {
        return p[point.sIdx..point.eIdx];
    }

    return "";
}

pub fn main() !void {
    const path = "./hello/world/mian.zig/";
    print("{s}\n", .{dir(path)});
    print("{s}\n", .{base(path)});
}

test "Dir test" {
    const pathTable = [_][_]u8{
        "",                  "",
        "./m",               "m",
        "main.go",           "",
        "./main.go",         "",
        "./home/main.go",    "home",
        ".///home/main.go/", "home",
        "/home/main.go///",  "home",
        "/home/go/main.go/", "home/go",
    };
    for (pathTable) |path| {
        try std.testing.expect(path[1] == dir(path[0]));
    }
}
test "File test" {
    const pathTable = [_][2][32]u8{
        "",                  "",
        "./m",               "",
        "main.go",           "main.go",
        "./main.go",         "main.go",
        "./home/main.go",    "main.go",
        ".///home/main.go/", "main.go",
        "/home/main.go///",  "main.go",
        "/home/go/main.go/", "main.go",
    };
    for (pathTable) |path| {
        try std.testing.expect(path[1] == base(path[0]));
    }
}
