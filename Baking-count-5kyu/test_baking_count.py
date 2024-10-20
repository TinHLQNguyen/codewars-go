import pytest
from baking_count import cakes


@pytest.mark.parametrize(
    "recipe, ingredient, numCake",
    [
        (
            {"flour": 500, "sugar": 200, "eggs": 1},
            {"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200},
            2,
        ),
        (
            {"flour": 500, "sugar": 200, "eggs": 1},
            {"flour": 500, "sugar": 200, "eggs": 5, "milk": 200},
            1,
        ),
        (
            {"flour": 500, "sugar": 200, "eggs": 1},
            {"flour": 500, "sugar": 200, "eggs": 1},
            1,
        ),
    ],
)
def test_HaveCake(recipe, ingredient, numCake):
    assert cakes(recipe, ingredient) == numCake


@pytest.mark.parametrize(
    "recipe, ingredient, numCake",
    [
        (
            {"apples": 3, "flour": 3000, "sugar": 700, "milk": 3000, "oil": 100},
            {"sugar": 500, "flour": 2000, "milk": 2000},
            0,
        ),
    ],
)
def test_MissingIngredient(recipe, ingredient, numCake):
    assert cakes(recipe, ingredient) == numCake


@pytest.mark.parametrize(
    "recipe, ingredient, numCake",
    [
        (
            {"sugar": 150, "flour": 300, "milk": 100},
            {"sugar": 500, "flour": 2000, "milk": 2000},
            0,
        ),
    ],
)
def test_NotEnoughIngredient(recipe, ingredient, numCake):
    assert cakes(recipe, ingredient) == numCake


@pytest.mark.parametrize(
    "recipe, ingredient, numCake",
    [
        (
            {"sugar": 150, "flour": 300, "milk": 100},
            {},
            0,
        ),
    ],
)
def test_NoIngredient(recipe, ingredient, numCake):
    assert cakes(recipe, ingredient) == numCake
