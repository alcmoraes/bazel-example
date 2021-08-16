load("@io_bazel_rules_go//go:def.bzl", "go_binary")
go_binary(
    name = "main_app",
    srcs = ["main.go"],
    importpath = "go-webview-wasm/packages/main",
    deps = [
        "//packages/shared/handlers/router:router"
    ],
)