import pytest
from ip_validation import is_valid_IP


@pytest.mark.parametrize(
    "given",
    [("1.2.3.4"), ("123.45.67.89")],
)
def test_goodIP(given):
    assert is_valid_IP(given) is True


@pytest.mark.parametrize(
    "given",
    [("1.2.3"), ("1.2.3.4.5"), ("")],
)
def test_invalidOctets(given):
    assert is_valid_IP(given) is False


@pytest.mark.parametrize(
    "given",
    [("abc.def.ghi.jk"), ("...")],
)
def test_invalidTypes(given):
    assert is_valid_IP(given) is False


@pytest.mark.parametrize(
    "given",
    [("123.500.4.90"), ("-10.254.4.90")],
)
def test_invalidValueRange(given):
    assert is_valid_IP(given) is False


@pytest.mark.parametrize(
    "given",
    [
        ("123.045.067.089"),
        ("1. 45.67.89"),
        ("1.45.67 .89"),
        (" 1.2.3.4"),
        ("1.2.3.4 "),
        ("\n1.2.3.4"),
        ("1.2.3.4\n"),
    ],
)
def test_trailing(given):
    assert is_valid_IP(given) is False
