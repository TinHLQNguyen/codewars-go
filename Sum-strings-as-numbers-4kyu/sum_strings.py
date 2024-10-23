UNICODE0 = 48
UNICODE1 = 49
UNICODE9 = 57


def sum_strings(x: str, y: str) -> str:
    x = PreProcessing(x)
    y = PreProcessing(y)
    remainder = False
    digitsInCodePoint = []
    lenX = len(x)
    lenY = len(y)
    ## add padding so that don't care about out-of-bound anymore
    if lenX >= lenY:
        y = y.rjust(lenX, "0")
        longerLen = lenX
    else:
        x = x.rjust(lenY, "0")
        longerLen = lenY

    for idx in range(-1, -longerLen - 1, -1):
        sumCharCode = (
            bytes(x[idx], "utf-8")[0] + bytes(y[idx], "utf-8")[0] - UNICODE0 + remainder
        )
        if sumCharCode > UNICODE9:
            remainder = True
            sumCharCode -= 10
        else:
            remainder = False
        digitsInCodePoint.insert(0, sumCharCode)
    if remainder:
        digitsInCodePoint.insert(0, UNICODE1)
    return bytes(digitsInCodePoint).decode("utf-8")


def PreProcessing(input: str) -> str:
    # make "" into 0
    # remove leading 0
    lStripted = input.lstrip("0")
    if lStripted == "":
        return "0"
    else:
        return lStripted
