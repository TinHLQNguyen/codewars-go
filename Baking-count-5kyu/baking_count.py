def cakes(recipe: dict, available: dict) -> int:
    # NOTE: assuming recipe is not empty
    requiredIngredients = recipe.keys()
    availableIngredients = available.keys()

    numCake = 0
    # check enough requiredIngredients
    for item in requiredIngredients:
        if item not in availableIngredients or available[item] == 0:
            return 0
        else:
            thisIngredientCapacity = available[item] / recipe[item]
            numCake = (
                thisIngredientCapacity
                if (thisIngredientCapacity < numCake or numCake == 0)
                else numCake
            )
    return int(numCake)
