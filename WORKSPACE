load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Setup Rust support.
http_archive(
    name = "rules_rust",
    integrity = "sha256-Y4v6kjQQfXxh5tU6FQB6YXux/ODFGUq3IlpgBV4Bwj8=",
    urls = ["https://github.com/bazelbuild/rules_rust/releases/download/0.41.0/rules_rust-v0.41.0.tar.gz"],
)

load("@rules_rust//rust:repositories.bzl", "rules_rust_dependencies", "rust_register_toolchains")

rules_rust_dependencies()

rust_register_toolchains(
    edition = "2021",
    extra_exec_rustc_flags = [
        "-A",
        "async_fn_in_trait",
    ],
    rust_analyzer_version = "nightly/2024-03-27",
    rustfmt_version = "nightly/2024-03-27",
    versions = [
        "nightly/2024-03-27",
    ],
)

load("@rules_rust//crate_universe:repositories.bzl", "crate_universe_dependencies")

crate_universe_dependencies()

load("@rules_rust//crate_universe:defs.bzl", "crate", "crates_repository")

crates_repository(
    name = "crate_index",
    annotations = {
        "protoc-gen-prost": [crate.annotation(
            gen_binaries = ["protoc-gen-prost"],
        )],
        "protoc-gen-tonic": [crate.annotation(
            gen_binaries = ["protoc-gen-tonic"],
        )],
    },
    cargo_lockfile = "//:Cargo.lock",
    lockfile = "//:Cargo.Bazel.lock",
    packages = {
        "amqprs": crate.spec(
            features = [
                "traces",
                "tls",
                "urispec",
            ],
            version = "1.5.4",
        ),
        "anyhow": crate.spec(
            features = ["backtrace"],
            version = "1.0.75",
        ),
        "argon2": crate.spec(
            version = "0.5.3",
        ),
        "async-trait": crate.spec(
            version = "0.1.77",
        ),
        "aws-config": crate.spec(
            features = ["behavior-version-latest"],
            version = "1.1.7",
        ),
        "aws-sdk-s3": crate.spec(
            features = ["behavior-version-latest"],
            version = "1.21.0",
        ),
        "aws-smithy-types": crate.spec(
            version = "1.1.8",
        ),
        "axum": crate.spec(
            features = [
                "macros",
                "multipart",
            ],
            version = "0.7.5",
        ),
        "axum-core": crate.spec(
            version = "0.4.3",
        ),
        "axum-extra": crate.spec(
            features = ["cookie"],
            version = "0.9.3",
        ),
        "axum-login": crate.spec(
            version = "0.15.1",
        ),
        "base64": crate.spec(
            version = "0.22.0",
        ),
        "bson": crate.spec(
            features = ["serde_with"],
            version = "2.7.0",
        ),
        "bytes": crate.spec(
            features = ["serde"],
            version = "1.6.0",
        ),
        "cargo_metadata": crate.spec(
            version = "0.17.0",
        ),
        "clap": crate.spec(
            features = ["derive"],
            version = "4.3.23",
        ),
        "colored": crate.spec(
            version = "2.0.4",
        ),
        "cgroups-rs": crate.spec(
            version = "0.3.4",
        ),
        "diesel": crate.spec(
            features = [
                "postgres",
                "mysql",
                "sqlite",
            ],
            version = "2.1.5",
        ),
        "futures": crate.spec(
            version = "0.3.28",
        ),
        "handlebars": crate.spec(
            version = "5.1.2",
        ),
        "http": crate.spec(
            version = "1.1.0",
        ),
        "image": crate.spec(
            features = ["rgb"],
            version = "0.24.7",
        ),
        "imageproc": crate.spec(
            version = "0.23.0",
        ),
        "infer": crate.spec(
            version = "0.15.0",
        ),
        "itertools": crate.spec(
            version = "0.11.0",
        ),
        "libc": crate.spec(
            version = "0.2.153",
        ),
        "log": crate.spec(
            version = "0.4.21",
        ),
        "maud": crate.spec(
            features = [
                "axum",
                "axum-core",
            ],
            version = "0.25.0",
        ),
        "mongodb": crate.spec(
            version = "2.8.2",
        ),
        "num-derive": crate.spec(
            version = "0.4.0",
        ),
        "num-traits": crate.spec(
            version = "0.2.16",
        ),
        "oauth2": crate.spec(
            version = "4.2.2",
        ),
        "poise": crate.spec(
            version = "0.6.1-rc1",
        ),
        "prometheus": crate.spec(
            version = "0.13.3",
        ),
        "proc-macro2": crate.spec(
            version = "1.0.66",
        ),
        "prost": crate.spec(
            version = "0.12.0",
        ),
        "prost-types": crate.spec(
            version = "0.12.0",
        ),
        "protoc-gen-prost": crate.spec(
            version = "0.3.1",
        ),
        "protoc-gen-tonic": crate.spec(
            version = "0.3.0",
        ),
        "quote": crate.spec(
            version = "1.0.33",
        ),
        "rand": crate.spec(
            version = "0.8.5",
        ),
        "rayon": crate.spec(
            version = "1.10.0",
        ),
        "redis": crate.spec(
            features = [
                "tokio",
                "json",
                "aio",
                "tls-native-tls",
                "tokio-native-tls-comp",
                "r2d2",
                "connection-manager",
            ],
            version = "0.25.2",
        ),
        "regex": crate.spec(
            version = "1.9.5",
        ),
        "reqwest": crate.spec(
            default_features = False,
            features = [
                "json",
                "native-tls",
                "default-tls",
            ],
            version = "0.12.2",
        ),
        "rss": crate.spec(
            version = "2.0.6",
        ),
        "rusttype": crate.spec(
            version = "0.9.3",
        ),
        "serde": crate.spec(
            features = [
                "derive",
                "serde_derive",
            ],
            version = "1.0.188",
        ),
        "serde_json": crate.spec(
            version = "1.0.105",
        ),
        "sha2": crate.spec(
            version = "0.10.8",
        ),
        "sqlx": crate.spec(
            features = [
                "tls-native-tls",
                "runtime-tokio",
                "mysql",
                "macros",
                "migrate",
                "time",
                "json",
                "any",
            ],
            version = "0.7.4",
        ),
        "superconsole": crate.spec(
            version = "0.2.0",
        ),
        "syn": crate.spec(
            features = ["full"],
            version = "2.0.29",
        ),
        "tempfile": crate.spec(
            version = "3.9.0",
        ),
        "thiserror": crate.spec(
            version = "1.0.24",
        ),
        "time": crate.spec(
            version = "0.3.28",
        ),
        "tracing": crate.spec(
            version = "0.1.37",
        ),
        "tracing-subscriber": crate.spec(
            features = [
                "fmt",
                "json",
                "std",
                "env-filter",
            ],
            version = "0.3",
        ),
        "trait-variant": crate.spec(
            version = "0.1.2",
        ),
        "tokio": crate.spec(
            features = ["full"],
            version = "1.37.0",
        ),
        "tokio-util": crate.spec(
            version = "0.7.10",
        ),
        "toml_edit": crate.spec(
            version = "0.19.14",
        ),
        "tonic": crate.spec(
            version = "0.10.0",
        ),
        "tower-http": crate.spec(
            features = [
                "cors",
                "fs",
                "trace",
            ],
            version = "0.5.2",
        ),
        "url": crate.spec(
            version = "2.5.0",
        ),
        "uuid": crate.spec(
            version = "1.7.0",
        ),
        "walkdir": crate.spec(
            version = "2.3.3",
        ),
        "which": crate.spec(
            version = "6.0.1",
        ),
        "validator": crate.spec(
            features = ["derive"],
            version = "0.18.1",
        ),
    },
    rust_version = "nightly/2024-03-27",
)

load("@crate_index//:defs.bzl", "crate_repositories")

crate_repositories()

# Setup Rust analyzer support.
load("@rules_rust//tools/rust_analyzer:deps.bzl", "rust_analyzer_dependencies")

rust_analyzer_dependencies()

# Setup RPC support.
load("@rules_rust//proto/prost:repositories.bzl", "rust_prost_dependencies")

rust_prost_dependencies()

load("@rules_rust//proto/prost:transitive_repositories.bzl", "rust_prost_transitive_repositories")

rust_prost_transitive_repositories()

register_toolchains("//:prost_toolchain")

# Setup external make support.
http_archive(
    name = "rules_foreign_cc",
    # TODO: Get the latest sha256 value from a bazel debug message or the latest
    #       release on the releases page: https://github.com/bazelbuild/rules_foreign_cc/releases
    #
    # sha256 = "...",
    strip_prefix = "rules_foreign_cc-51152aac9d6d8b887802a47ec08a1a37ef2c4885",
    url = "https://github.com/bazelbuild/rules_foreign_cc/archive/51152aac9d6d8b887802a47ec08a1a37ef2c4885.tar.gz",
)

load("@rules_foreign_cc//foreign_cc:repositories.bzl", "rules_foreign_cc_dependencies")

rules_foreign_cc_dependencies()

# Setup bindgen support.
load("@rules_rust//bindgen:repositories.bzl", "rust_bindgen_dependencies", "rust_bindgen_register_toolchains")

rust_bindgen_dependencies()

rust_bindgen_register_toolchains()

load("@rules_rust//bindgen:transitive_repositories.bzl", "rust_bindgen_transitive_dependencies")

rust_bindgen_transitive_dependencies()

