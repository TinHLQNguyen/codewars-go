UNICODE0 = 48


def sum_strings(x: str, y: str) -> str:
    xbArray = PreProcessing(x)
    ybArray = PreProcessing(y)
    remainder = False
    digitsInCodePoint = []
    lenX = len(xbArray)
    lenY = len(ybArray)
    ## add padding so that don't care about out-of-bound anymore
    if lenX >= lenY:
        ybArray = ybArray.rjust(lenX, b"0")
        longerLen = lenX
    else:
        xbArray = xbArray.rjust(lenY, b"0")
        longerLen = lenY

    for idx in range(-1, -longerLen - 1, -1):
        tempSum = xbArray[idx] + ybArray[idx] - 2 * UNICODE0 + remainder
        remainder, digit = divmod(tempSum, 10)
        digitsInCodePoint.insert(0, digit + UNICODE0)
    if remainder:
        digitsInCodePoint.insert(0, UNICODE0 + 1)
    return bytes(digitsInCodePoint).decode("utf-8")


def PreProcessing(input: str) -> bytearray:
    # make "" into 0
    # remove leading 0
    barrays = bytearray(input, "utf-8")
    lStripted = barrays.lstrip(b"0")
    if lStripted == b"":
        return bytearray("0", "utf-8")
    else:
        return lStripted
