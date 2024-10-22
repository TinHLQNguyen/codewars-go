from typing import Tuple

UNICODE0 = 48
UNICODE9 = 57


def sum_strings(x: str, y: str) -> str:
    x = PreProcessing(x)
    y = PreProcessing(y)
    digitX = ""
    digitY = ""
    remainder = False
    sum = ""
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
        digitX = x[idx]
        digitY = y[idx]
        digitSum, remainder = SumDigit(digitX, digitY, remainder)
        sum = digitSum + sum
    if remainder:
        sum = "1" + sum
    return sum


def SumDigit(a: str, b: str, remainder: bool) -> Tuple[str, bool]:
    sumCharCode = ord(a) + ord(b) - UNICODE0 + remainder
    if sumCharCode > UNICODE9:
        remainder = True
        sumCharCode -= 10
    else:
        remainder = False
    return chr(sumCharCode), remainder


def PreProcessing(input: str) -> str:
    # make "" into 0
    # remove leading 0
    lStripted = input.lstrip("0")
    if lStripted == "":
        return "0"
    else:
        return lStripted
