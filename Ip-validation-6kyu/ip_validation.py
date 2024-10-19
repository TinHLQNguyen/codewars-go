def is_valid_IP(input: str) -> bool:
    octets = input.split(".")
    if len(octets) != 4:
        return False
    for octet in octets:
        # check valid digits
        if not octet.isdigit():
            return False

        # check leading '0' case
        if octet.startswith("0") and len(octet) > 1:
            return False

        # check int conversion case
        try:
            octet = int(octet)
        except ValueError:
            return False

        # check value case
        if (octet < 0) or (octet > 255):
            return False

    return True
