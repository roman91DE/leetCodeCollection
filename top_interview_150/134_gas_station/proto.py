#!/usr/bin/env python
# coding: utf-8


def canCompleteCircuitRecursive(gas: list[int], cost: list[int]) -> int:

    if sum(gas) < sum(cost):
        return -1

    def possible_starting_positions(gas, cost):
        valid = []
        for i, (payload, cost) in enumerate(zip(gas, cost)):
            if payload >= cost:
                valid.append(i)
        return valid

    def recurse(gas, cost, tank, start_pos=None):
        if len(gas) == 0:
            return start_pos
        elif (tank + gas[0]) < cost[0]:
            return -1
        else:
            new_tank = tank + gas[0] - cost[0]
            return recurse(gas[1:], cost[1:], new_tank, start_pos)

    starting_pos = possible_starting_positions(gas, cost)

    if len(starting_pos) == 1:
        return starting_pos[0]

    for pos in starting_pos:

        gas_new = gas[pos + 1 :] + gas[:pos]
        cost_new = cost[pos + 1 :] + cost[:pos]
        tank_new = gas[pos] - cost[pos]

        sol = recurse(gas_new, cost_new, tank_new, pos)
        if sol != -1:
            return sol

    return -1


def canCompleteCircuitIterative(gas: list[int], cost: list[int]) -> int:

    if sum(gas) < sum(cost):
        return -1

    def possible_starting_positions(gas, cost):
        valid = []
        for i, (payload, cost) in enumerate(zip(gas, cost)):
            if payload >= cost:
                valid.append(i)
        return valid

    def loop(gas, cost, tank, start_pos=None):

        while len(gas) > 0:

            if (tank + gas[0]) < cost[0]:
                return -1
            tank += gas[0] - cost[0]

            gas = gas[1:]
            cost = cost[1:]

        else:
            return start_pos

    starting_pos = possible_starting_positions(gas, cost)

    if len(starting_pos) == 1:
        return starting_pos[0]

    for pos in starting_pos:

        gas_new = gas[pos + 1 :] + gas[:pos]
        cost_new = cost[pos + 1 :] + cost[:pos]
        tank_new = gas[pos] - cost[pos]

        sol = loop(gas_new, cost_new, tank_new, pos)
        if sol != -1:
            return sol

    return -1


# Test cases
test_cases = [
    {"gas": [1, 2, 3, 4, 5], "cost": [3, 4, 5, 1, 2], "expected": 3},
    {"gas": [2, 3, 4], "cost": [3, 4, 3], "expected": -1},
    {"gas": [5], "cost": [4], "expected": 0},
    {"gas": [2, 1, 5, 3], "cost": [3, 2, 4, 5], "expected": -1},
]

# Run tests
for i, test in enumerate(test_cases, 1):
    gas: list[int] = test["gas"]  # type: ignore
    cost: list[int] = test["cost"]  # type: ignore
    expected = test["expected"]
    result = canCompleteCircuitIterative(gas, cost)
    print(
        f"ITERATIVE: Test Case {i}: {'PASSED' if result == expected else 'FAILED'} (Output: {result}, Expected: {expected})"
    )

    result = canCompleteCircuitRecursive(gas, cost)
    print(
        f"RECURSIVE: Test Case {i}: {'PASSED' if result == expected else 'FAILED'} (Output: {result}, Expected: {expected})"
    )
