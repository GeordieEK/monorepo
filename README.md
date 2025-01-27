# The Monorepo

This is a monorepo of various personal projects, with Bazel managing the build, test, and dependency processes.

## Structure

- `apps/`: Contains the main applications (e.g., web, API, CLI).
- `packages/`: Shared libraries and modules.
- `tools/`: Utility scripts and tools.
- `docs/`: Project documentation.

## How to Use

1. Install Bazel: https://bazel.build/
2. Run builds: `bazel build //...`
3. Run tests: `bazel test //...`

## Credits / References

[These Bazel examples](https://github.com/aspect-build/bazel-examples/blob/main/MODULE.bazel) by @aspect-build.
[The official bazel examples](https://github.com/bazelbuild/examples/tree/main) are pretty good starts too.
