load("//tools:defs.bzl", "go_library")

package(licenses = ["notice"])

go_library(
    name = "bigquery",
    testonly = 1,
    srcs = ["bigquery.go"],
    nogo = False,  # FIXME(b/184974218): Analysis failing for cloud libraries.
    visibility = [
        "//:sandbox",
    ],
    deps = [
        "@com_google_cloud_go_bigquery//:go_default_library",
        "@org_golang_google_api//option:go_default_library",
        "@org_golang_x_oauth2//:go_default_library",
    ],
)
