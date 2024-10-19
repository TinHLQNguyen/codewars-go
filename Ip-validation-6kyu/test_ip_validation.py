from ip_validation import is_valid_IP


def test_goodIP():
    assert is_valid_IP("1.2.3.4") is True
    assert is_valid_IP("123.45.67.89") is True


def test_invalidOctets():
    assert is_valid_IP("1.2.3") is False
    assert is_valid_IP("1.2.3.4.5") is False
    assert is_valid_IP("") is False


def test_invalidTypes():
    assert is_valid_IP("abc.def.ghi.jk") is False
    assert is_valid_IP("...") is False


def test_invalidValueRange():
    assert is_valid_IP("123.500.4.90") is False


def test_trailing():
    assert is_valid_IP("123.045.067.089") is False
    assert is_valid_IP("1. 45.67.89") is False
    assert is_valid_IP("1.45.67 .89") is False
    assert is_valid_IP(" 1.2.3.4") is False
    assert is_valid_IP("1.2.3.4 ") is False
    assert is_valid_IP("\n1.2.3.4") is False
    assert is_valid_IP("1.2.3.4\n") is False
