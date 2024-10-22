from typing import Tuple

SumTable = {
    ## 0 group
    "00": ("0", False),
    "01": ("1", False),
    "02": ("2", False),
    "03": ("3", False),
    "04": ("4", False),
    "05": ("5", False),
    "06": ("6", False),
    "07": ("7", False),
    "08": ("8", False),
    "09": ("9", False),
    ## 1 group
    "11": ("2", False),
    "12": ("3", False),
    "13": ("4", False),
    "14": ("5", False),
    "15": ("6", False),
    "16": ("7", False),
    "17": ("8", False),
    "18": ("9", False),
    "19": ("0", True),
    ## 2 group
    "22": ("4", False),
    "23": ("5", False),
    "24": ("6", False),
    "25": ("7", False),
    "26": ("8", False),
    "27": ("9", False),
    "28": ("0", True),
    "29": ("1", True),
    ## 3 group
    "33": ("6", False),
    "34": ("7", False),
    "35": ("8", False),
    "36": ("9", False),
    "37": ("0", True),
    "38": ("1", True),
    "39": ("2", True),
    ## 4 group
    "44": ("8", False),
    "45": ("9", False),
    "46": ("0", True),
    "47": ("1", True),
    "48": ("2", True),
    "49": ("3", True),
    ## 5 group
    "55": ("0", True),
    "56": ("1", True),
    "57": ("2", True),
    "58": ("3", True),
    "59": ("4", True),
    ## 6 group
    "66": ("2", True),
    "67": ("3", True),
    "68": ("4", True),
    "69": ("5", True),
    ## 7 group
    "77": ("4", True),
    "78": ("5", True),
    "79": ("6", True),
    ## 8 group
    "88": ("6", True),
    "89": ("7", True),
    ## 9 group
    "99": ("8", True),
}


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
        idx -= 1
    if remainder:
        sum = "1" + sum
    return sum


def SumDigit(a: str, b: str, remainder: bool) -> Tuple[str, bool]:
    if (a + b) in SumTable:
        tempSum, newremainer = SumTable[a + b]
    else:
        tempSum, newremainer = SumTable[b + a]
    if remainder:
        if ("1" + tempSum) in SumTable:
            sum, returnRemainder = SumTable["1" + tempSum]
        else:
            sum, returnRemainder = SumTable[tempSum + "1"]
        return sum, returnRemainder or newremainer
    else:
        return tempSum, newremainer


def PreProcessing(input: str) -> str:
    # make "" into 0
    # remove leading 0
    lStripted = input.lstrip("0")
    if lStripted == "":
        return "0"
    else:
        return lStripted
