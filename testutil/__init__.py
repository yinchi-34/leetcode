"""Shared helpers for the Python tests."""

import importlib.util
from pathlib import Path
from types import ModuleType


def load_solution(test_file: str, subdir: str = "") -> ModuleType:
    """Load solution.py from the test file's directory, or from one of its
    subdirectories when a problem keeps several solutions side by side.

    Problem directories are not importable package names (they start with a
    digit and contain hyphens), and every one of them has a solution.py, so
    each module is loaded by path under a name derived from that path.
    """
    path = Path(test_file).resolve().parent / subdir / "solution.py"
    name = "solution_" + "_".join(path.parent.parts[-3:]).replace("-", "_")
    spec = importlib.util.spec_from_file_location(name, path)
    if spec is None or spec.loader is None:
        raise ImportError(f"cannot load {path}")
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module
