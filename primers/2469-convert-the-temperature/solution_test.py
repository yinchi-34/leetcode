import pytest

from testutil import load_solution

convert_temperature = load_solution(__file__).Solution().convertTemperature


@pytest.mark.parametrize(
    "celsius, expected",
    [(36.50, [309.65, 97.70]), (122.11, [395.26, 251.798]), (0, [273.15, 32.00])],
    ids=["example_1", "example_2", "zero"],
)
def test_convert_temperature(celsius, expected):
    # Floating-point results are compared within the 1e-5 tolerance the
    # problem allows, not with ==.
    assert convert_temperature(celsius) == pytest.approx(expected, abs=1e-5)
