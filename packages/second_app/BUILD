load("@io_bazel_rules_go//go:def.bzl", "go_binary")
go_binary(
    name = "second_app",
    srcs = ["main.go"],
    importpath = "bazel-example/packages/second",
    deps = [
        "//packages/shared/handlers/router:router"
    ],
)