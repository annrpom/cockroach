# DO NOT EDIT THIS FILE MANUALLY! Use `release update-releases-file`.
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

CONFIG_LINUX_AMD64 = "linux-amd64"
CONFIG_LINUX_ARM64 = "linux-arm64"
CONFIG_DARWIN_AMD64 = "darwin-10.9-amd64"
CONFIG_DARWIN_ARM64 = "darwin-11.0-arm64"

_CONFIGS = [
    ("23.1.23", [
        (CONFIG_DARWIN_AMD64, "ca05a8d2d58bbd2f8b3a32c64885cd55c8aba2357095f5c5000d6fa8e2bcffa2"),
        (CONFIG_DARWIN_ARM64, "158ee95b2545013cbe87dfb333cef758c0bf4944813988dca90156b18a481e6a"),
        (CONFIG_LINUX_AMD64, "f45370b65056bc5e1fa4d3d9c686b7587f7178b1aa5deb818dfc80d34b4979d8"),
        (CONFIG_LINUX_ARM64, "a6d6a8cb541b954f3023559882066599f30bfcf33c22c8acb8a26189a5d458c5"),
    ]),
    ("23.2.7", [
        (CONFIG_DARWIN_AMD64, "6db44b8c2dfe5801cb8cf6d29dc81409f388fb1b499724098b9bda33b95852bc"),
        (CONFIG_DARWIN_ARM64, "fc9dbfb0746da06652696aca3d0648e03e2d3c97e003657283b372320ce4ef42"),
        (CONFIG_LINUX_AMD64, "a57fcb5a6ecb23b8bccf883192973c2e4fd58a450513d0e74d4484f57e9e2e33"),
        (CONFIG_LINUX_ARM64, "bcc71ff7bd6f8004afd07bce601747511c878cbd304faa23f3ab9bba2f3992f5"),
    ]),
]

def _munge_name(s):
    return s.replace("-", "_").replace(".", "_")

def _repo_name(version, config_name):
    return "cockroach_binary_v{}_{}".format(
        _munge_name(version),
        _munge_name(config_name))

def _file_name(version, config_name):
    return "cockroach-v{}.{}/cockroach".format(
        version, config_name)

def target(config_name):
    targets = []
    for versionAndConfigs in _CONFIGS:
        version, _ = versionAndConfigs
        targets.append("@{}//:{}".format(_repo_name(version, config_name),
                                         _file_name(version, config_name)))
    return targets

def cockroach_binaries_for_testing():
    for versionAndConfigs in _CONFIGS:
        version, configs = versionAndConfigs
        for config in configs:
            config_name, shasum = config
            file_name = _file_name(version, config_name)
            http_archive(
                name = _repo_name(version, config_name),
                build_file_content = """exports_files(["{}"])""".format(file_name),
                sha256 = shasum,
                urls = [
                    "https://binaries.cockroachdb.com/{}".format(
                        file_name.removesuffix("/cockroach")) + ".tgz",
                ],
            )
