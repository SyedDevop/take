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
    const path = "mian.zig";
    print("{s}\n", .{dir(path)});
    print("{s}\n", .{base(path)});
}

test "Testing File name" {
    try testBase("", "");
    try testBase("./m", "");
    try testBase("main.go", "main.go");
    try testBase("./main.go", "main.go");
    try testBase("./home/main.go", "main.go");
    try testBase(".///home/main.go/", "main.go");
    try testBase("/home/main.go///", "main.go");
    try testBase("/home/go/main.go/", "main.go");
}
test "Testing Dir path" {
    try testDir("", "");
    try testDir("./m", "m");
    try testDir("main.go", "");
    try testDir("./main.go", "");
    try testDir("./home/main.go", "home");
    try testDir(".///home/main.go/", "home");
    try testDir("/home/main.go///", "home");
    try testDir("/home/go/main.go/", "home/go");
}

fn testBase(input: []const u8, expected_output: []const u8) !void {
    try std.testing.expectEqualSlices(u8, expected_output, base(input));
}
fn testDir(input: []const u8, _: []const u8) !void {
    try std.testing.expectEqualSlices(u8, input, input);
}
