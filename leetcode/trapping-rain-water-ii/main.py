from collections import deque, defaultdict
from colorsys import hls_to_rgb
from email.policy import default
from re import A, L


class AreaLeeks(Exception):
    pass


class SolutionTypeHints(object):
    area_number_generator = 0

    def trapRainWater(self, heightMap):
        """
        :type heightMap: List[List[int]]
        :rtype: int
        """

        self.area_number_generator = 0

        LL, CC = len(heightMap), len(heightMap[0])
        visited = set()

        # areas and the maximum height of that area
        areas = {}
        areaheights = {}

        for i in range(LL):
            for j in range(CC):
                self.area_number_generator += 1
                areano = self.area_number_generator
                self.explore_area(
                    areas,
                    areaheights,
                    areano,
                    (i, j),
                    visited,
                    heightMap,
                    LL,
                    CC,
                )

        # print(areas, areaheights)
        s = 0
        for el in areas:
            areano = areas[el]
            i, j = el
            s += max(0, areaheights[areano] - heightMap[i][j])
        return s

        # for i in range(area_number_generator + 1):
        #     to_print = []
        #     for el in areas:
        #         if areas[el] == i:
        #             to_print.append(el)
        #     print(i, to_print)

    def explore_area(
        self,
        areas: dict[tuple[int, int], int],
        areaheights: dict[int, int],
        areano: int,
        starting_pos: tuple[int, int],
        visited: set[tuple[int, int]],
        heightMap,
        LL,
        CC,
    ):

        # ignore if it's already been explored, means it's in a different
        # area. check probably redundant
        if starting_pos in visited:
            return
        # ignore if it's already in a different area, probably the only
        # check needed
        if starting_pos in areas:
            return

        q = deque()
        i, j = starting_pos
        h = heightMap[i][j]
        q.append((i, j))

        areaheights[areano] = 99999999999999999999

        to_update_leaks = []
        while len(q) != 0:
            # set the location to the current area
            top = q.popleft()
            l, c = top
            visited.add((l, c))  # mark this as visited

            areas[top] = areano

            ns = [(l - 1, c), (l + 1, c), (l, c - 1), (l, c + 1)]

            for n in ns:
                ll, cc = n
                if not ((0 <= ll < LL) and (0 <= cc < CC)):
                    areaheights[areano] = -1  # will leak
                    continue

                # cannot leak to the neighbour, use its height as boundary
                if heightMap[ll][cc] > h:
                    areaheights[areano] = min(areaheights[areano], heightMap[ll][cc])
                    continue

                # neighbour already visited, ignore it
                if n in visited:
                    continue

                # regular neighbour
                if heightMap[ll][cc] == h:
                    q.append(n)
                # will leak in neighbour
                elif heightMap[ll][cc] < h:
                    # if not explored, explore
                    if n not in areas:
                        self.area_number_generator += 1
                        next_area_no = self.area_number_generator
                        self.explore_area(
                            areas,
                            areaheights,
                            next_area_no,
                            n,
                            visited,
                            heightMap,
                            LL,
                            CC,
                        )

                    # now that it's 100% explored, add to the leaky relationship
                    # this means, if the place it leaks into cannot hold to at
                    # least the current level, set it as leaky
                    if areaheights[next_area_no] < h:
                        areaheights[areano] = -1
                    else:
                        to_update_leaks.append(next_area_no)

        if areaheights[areano] != -1:
            for l in to_update_leaks:
                areaheights[l] = areaheights[areano]


from collections import deque, defaultdict
from colorsys import hls_to_rgb
from email.policy import default
from re import A, L


class AreaLeeks(Exception):
    pass


class SolutionTypeHints(object):
    area_number_generator = 0

    def trapRainWater(self, heightMap):
        """
        :type heightMap: List[List[int]]
        :rtype: int
        """

        self.area_number_generator = 0

        LL, CC = len(heightMap), len(heightMap[0])
        visited = set()

        # areas and the maximum height of that area
        areas = {}
        areaheights = {}

        for i in range(LL):
            for j in range(CC):
                self.area_number_generator += 1
                areano = self.area_number_generator
                self.explore_area(
                    areas,
                    areaheights,
                    areano,
                    (i, j),
                    visited,
                    heightMap,
                    LL,
                    CC,
                )

        # print(areas, areaheights)
        s = 0
        for el in areas:
            areano = areas[el]
            i, j = el
            s += max(0, areaheights[areano] - heightMap[i][j])
        return s

        # for i in range(area_number_generator + 1):
        #     to_print = []
        #     for el in areas:
        #         if areas[el] == i:
        #             to_print.append(el)
        #     print(i, to_print)

    def explore_area(
        self,
        areas: dict[tuple[int, int], int],
        areaheights: dict[int, int],
        areano: int,
        starting_pos: tuple[int, int],
        visited: set[tuple[int, int]],
        heightMap,
        LL,
        CC,
    ):

        # ignore if it's already been explored, means it's in a different
        # area. check probably redundant
        if starting_pos in visited:
            return
        # ignore if it's already in a different area, probably the only
        # check needed
        if starting_pos in areas:
            return

        q = deque()
        i, j = starting_pos
        h = heightMap[i][j]
        q.append((i, j))

        areaheights[areano] = 99999999999999999999

        to_update_leaks = []
        while len(q) != 0:
            # set the location to the current area
            top = q.popleft()
            l, c = top
            visited.add((l, c))  # mark this as visited

            areas[top] = areano

            ns = [(l - 1, c), (l + 1, c), (l, c - 1), (l, c + 1)]

            for n in ns:
                ll, cc = n
                if not ((0 <= ll < LL) and (0 <= cc < CC)):
                    areaheights[areano] = -1  # will leak
                    continue

                # cannot leak to the neighbour, use its height as boundary
                if heightMap[ll][cc] > h:
                    areaheights[areano] = min(areaheights[areano], heightMap[ll][cc])
                    continue

                # neighbour already visited, ignore it
                if n in visited:
                    continue

                # regular neighbour
                if heightMap[ll][cc] == h:
                    q.append(n)
                # will leak in neighbour
                elif heightMap[ll][cc] < h:
                    # if not explored, explore
                    if n not in areas:
                        self.area_number_generator += 1
                        next_area_no = self.area_number_generator
                        self.explore_area(
                            areas,
                            areaheights,
                            next_area_no,
                            n,
                            visited,
                            heightMap,
                            LL,
                            CC,
                        )

                    # now that it's 100% explored, add to the leaky relationship
                    # this means, if the place it leaks into cannot hold to at
                    # least the current level, set it as leaky
                    if areaheights[next_area_no] < h:
                        areaheights[areano] = -1
                    else:
                        to_update_leaks.append(next_area_no)

        if areaheights[areano] != -1:
            for l in to_update_leaks:
                areaheights[l] = areaheights[areano]


class Solution:
    area_number_generator = 0
    too_big = 99999999999999999999
    leaks_to = defaultdict(list)
    areaheights = {}
    can_hold = 0

    debug = False

    def trapRainWater(self, heightMap):
        """
        :type heightMap: List[List[int]]
        :rtype: int
        """

        # generate unique area numbers
        self.area_number_generator = 0

        LL, CC = len(heightMap), len(heightMap[0])
        visited = set()

        # areas and the maximum height of that area
        areas = {}
        areamaxpotentialheights = defaultdict(int)
        self.areaheights = {}

        self.leaks_to = defaultdict(set)
        self.leaks_from = defaultdict(set)

        self.can_hold = 0

        for i in range(LL):
            for j in range(CC):
                self.explore_area(
                    areas,
                    areamaxpotentialheights,
                    (i, j),
                    visited,
                    heightMap,
                    LL,
                    CC,
                )

        nodes_in_area = defaultdict(list)

        # we need to effectively find and "move" all nodes from an area to
        # a different area
        for k, v in areas.items():
            nodes_in_area[v].append(k)

        # basically the strategy is as follows:
        # pick any area with no outbound neihgbours (no area to leak to)
        # note that the "empty" area is an outbound neighbour
        # then basically pick the neighbour with smallest height that leaks to
        # it
        # fill all that area up to that location and merge it with all inbound
        # neihbours that have that size

        # note that if it doesn't leak anywhere, someone has to leak to it.
        # boundary nodes leak to "nowhere"

        while True:

            acc_point = -1

            found_acc_point = False
            for areano in nodes_in_area:
                if len(self.leaks_to[areano]) == 0:
                    acc_point = areano
                    found_acc_point = True
                    break

            if not found_acc_point:
                return self.can_hold

            leaker_sizes = []
            for leaker in self.leaks_from[acc_point]:
                leaker_sizes.append(self.areaheights[leaker])
            min_height_that_leaks_to_it = min(leaker_sizes)

            to_merge_without_acc_point = []
            for leaker in self.leaks_from[acc_point]:
                if self.areaheights[leaker] == min_height_that_leaks_to_it:
                    to_merge_without_acc_point.append(leaker)

            # merger step
            total_filled = (
                min_height_that_leaks_to_it - self.areaheights[acc_point]
            ) * len(nodes_in_area[acc_point])
            self.can_hold += total_filled

            # nodes_in_area, leaks_from, leaks_to, areaheights

            # increase area that we leak into's size
            self.areaheights[acc_point] = min_height_that_leaks_to_it

            # print("deleting", to_merge_without_acc_point)

            for tm in to_merge_without_acc_point:

                # if anything leaks to an area that we remove, adjust
                for lt, ltv in self.leaks_to.items():
                    if tm in ltv:
                        ltv.remove(tm)
                        ltv.add(acc_point)
                        self.leaks_from[acc_point].add(lt)

                self.leaks_from[acc_point].remove(tm)

                for remaining in self.leaks_to[tm]:
                    if (
                        remaining == acc_point
                    ):  # obvious that since we filled it up it's not relevant anymore
                        continue
                    self.leaks_to[acc_point].add(remaining)
                    if remaining != -1:
                        self.leaks_from[remaining].add(acc_point)

                for lf, lfv in self.leaks_from.items():
                    if tm in lfv:
                        lfv.remove(tm)

                del self.areaheights[tm]
                self.leaks_to.pop(tm, None)
                self.leaks_from.pop(tm, None)
                nodes_in_area[acc_point].extend(nodes_in_area[tm])
                del nodes_in_area[tm]

            # break

    def explore_area(
        self,
        areas,
        areaheights,
        starting_pos,
        visited,
        heightMap,
        LL,
        CC,
    ) -> int:

        # ignore if it's already been explored, means it's in a different
        # area. check probably redundant
        if starting_pos in visited:
            return
        # ignore if it's already in a different area, probably the only
        # check needed
        if starting_pos in areas:
            return
        # generate a unique area number for exploration
        self.area_number_generator += 1
        areano = self.area_number_generator

        q = deque()
        i, j = starting_pos
        h = heightMap[i][j]
        q.append((i, j))

        self.areaheights[areano] = h

        # if it's just the start of the exploration, use a big number as the
        # maximum potential water this area can hold
        areaheights[areano] = self.too_big

        # optimisation, a neighbour may be explored and visited by an inner
        # bfs but you still want to use it for the current exploration to
        # either mark it as a boundary or a neighbour you leak into
        used_as_neighbour = set()

        while len(q) != 0:
            # set the location to the current area
            top = q.popleft()
            l, c = top
            visited.add((l, c))  # mark this as visited

            # put this in the corresponding area number
            areas[top] = areano

            ns = [(l - 1, c), (l + 1, c), (l, c - 1), (l, c + 1)]

            for n in ns:
                ll, cc = n

                # neighbour already visited, ignore it
                if n in used_as_neighbour:
                    continue

                if not ((0 <= ll < LL) and (0 <= cc < CC)):
                    areaheights[areano] = -1  # will leak
                    self.leaks_to[areano].add(-1)  # leaks to "nowhere"
                    continue  # however continue to add all neighbours to the area

                # cannot leak to the neighbour, use its height as boundary
                if heightMap[ll][cc] > h:
                    areaheights[areano] = min(areaheights[areano], heightMap[ll][cc])
                    continue

                # regular neighbour that was not visited before
                if heightMap[ll][cc] == h and n not in visited:
                    q.append(n)

                # will leak in neighbour
                elif heightMap[ll][cc] < h:
                    # if not explored, explore
                    if n not in areas:
                        next_area_no = self.explore_area(
                            areas,
                            areaheights,
                            n,
                            visited,
                            heightMap,
                            LL,
                            CC,
                        )
                    else:
                        # will still leak into it even if already explored
                        next_area_no = areas[n]

                    # add it to the leakers relationship, if not already there
                    self.leaks_to[areano].add(next_area_no)
                    self.leaks_from[next_area_no].add(areano)

                    # now that it's 100% explored, add to the leaky relationship
                    # this means, if the place it leaks into cannot hold to at
                    # least the current level, set it as leaky
                    if areaheights[next_area_no] < h:
                        areaheights[areano] = -1

                used_as_neighbour.add(n)

        if areaheights[areano] == self.too_big:
            areaheights[areano] = -1  # no boundaries, won't hold.

        # print("to update leaks", to_update_leaks)
        return areano


# TODO: leaking in two areas, but only one is completely falling. the second
# area may still be full

# l = [[1, 2], [55, 12, 34]]
# p = l[1]
# p.sort()
# print(l)

s = Solution()
print(
    "can fill up to",
    s.trapRainWater(
        [
            [3, 3, 3, 3, 3],
            [3, 2, 2, 2, 3],
            [3, 2, 1, 2, 3],
            [3, 2, 2, 2, 3],
            [3, 3, 3, 3, 3],
        ]
    ),
)
print(
    "can fill up to",
    s.trapRainWater([[1, 4, 3, 1, 3, 2], [3, 2, 1, 3, 2, 4], [2, 3, 3, 2, 3, 1]]),
)
print(
    "can fill up to",
    s.trapRainWater([[1, 3, 3, 1, 3, 2], [3, 2, 1, 3, 2, 3], [3, 3, 3, 2, 3, 1]]),
)
print(
    "can fill up to",
    s.trapRainWater(
        [
            [12, 13, 1, 12],
            [13, 4, 13, 12],
            [13, 8, 10, 12],
            [12, 13, 12, 12],
            [13, 13, 13, 13],
        ]
    ),
)
print(
    "can fill up to",
    s.trapRainWater([[2, 3, 4], [5, 6, 7], [8, 9, 10], [11, 12, 13], [14, 15, 16]]),
)
print(
    "can fill up to",
    s.trapRainWater(
        [
            [9, 9, 9, 9, 9],
            [9, 2, 1, 2, 8],
            [9, 2, 6, 2, 8],
            [9, 2, 3, 2, 8],
            [9, 8, 8, 7, 8],
        ]
    ),
)
print(
    "can fill up to",
    s.trapRainWater(
        [
            [78, 16, 94, 36],
            [87, 93, 50, 22],
            [63, 28, 91, 60],
            [64, 27, 41, 27],
            [73, 37, 12, 69],
            [68, 30, 83, 31],
            [63, 24, 68, 36],
        ]
    ),
)


# class Solution(object):
#     area_number_generator = 0
#     too_big = 99999999999999999999

#     def trapRainWater(self, heightMap):
#         """
#         :type heightMap: List[List[int]]
#         :rtype: int
#         """

#         self.area_number_generator = 0

#         LL, CC = len(heightMap), len(heightMap[0])
#         visited = set()

#         # areas and the maximum height of that area
#         areas = {}
#         areaheights = {}

#         leaks_to = defaultdict(list)

#         for i in range(LL):
#             for j in range(CC):
#                 self.area_number_generator += 1
#                 areano = self.area_number_generator
#                 self.explore_area(
#                     areas,
#                     areaheights,
#                     areano,
#                     (i, j),
#                     visited,
#                     heightMap,
#                     LL,
#                     CC,
#                     leaks_to,
#                 )

#         print(areas, areaheights)
#         s = 0
#         for el in areas:
#             areano = areas[el]
#             i, j = el
#             s += max(0, areaheights[areano] - heightMap[i][j])
#         for i in range(self.area_number_generator + 1):
#             to_print = []
#             for el in areas:
#                 if areas[el] == i:
#                     to_print.append(el)
#             print(i, to_print)
#         print("leaks to", leaks_to)
#         return s

#     def explore_area(
#         self,
#         areas,
#         areaheights,
#         areano,
#         starting_pos,
#         visited,
#         heightMap,
#         LL,
#         CC,
#         leaks_to,
#     ):

#         # ignore if it's already been explored, means it's in a different
#         # area. check probably redundant
#         if starting_pos in visited:
#             return
#         # ignore if it's already in a different area, probably the only
#         # check needed
#         if starting_pos in areas:
#             return

#         q = deque()
#         i, j = starting_pos
#         h = heightMap[i][j]
#         q.append((i, j))

#         areaheights[areano] = self.too_big

#         while len(q) != 0:
#             # set the location to the current area
#             top = q.popleft()
#             l, c = top
#             visited.add((l, c))  # mark this as visited

#             areas[top] = areano

#             ns = [(l - 1, c), (l + 1, c), (l, c - 1), (l, c + 1)]

#             for n in ns:
#                 ll, cc = n
#                 if not ((0 <= ll < LL) and (0 <= cc < CC)):
#                     areaheights[areano] = -1  # will leak
#                     continue

#                 # cannot leak to the neighbour, use its height as boundary
#                 if heightMap[ll][cc] > h:
#                     areaheights[areano] = min(areaheights[areano], heightMap[ll][cc])
#                     continue

#                 # neighbour already visited, ignore it
#                 # however, you need to take care of the case where the neighbour
#                 # is already visited, but you could leak into it
#                 #
#                 # todo: also study a -> b -> c and a->c
#                 if n in visited:
#                     if heightMap[ll][cc] < h and n in areas:
#                         leaks_to[areano].append(areas[n])
#                         if areaheights[areas[n]] < h:
#                             areaheights[areano] = -1
#                     continue

#                 # regular neighbour
#                 if heightMap[ll][cc] == h:
#                     q.append(n)
#                 # will leak in neighbour
#                 elif heightMap[ll][cc] < h:
#                     next_area_no = -99
#                     # if not explored, explore
#                     if n not in areas:
#                         self.area_number_generator += 1
#                         next_area_no = self.area_number_generator
#                         self.explore_area(
#                             areas,
#                             areaheights,
#                             next_area_no,
#                             n,
#                             visited,
#                             heightMap,
#                             LL,
#                             CC,
#                             leaks_to,
#                         )
#                     else:
#                         next_area_no = areas[n]

#                     leaks_to[areano].append(next_area_no)

#                     # now that it's 100% explored, add to the leaky relationship
#                     # this means, if the place it leaks into cannot hold to at
#                     # least the current level, set it as leaky
#                     if areaheights[next_area_no] < h:
#                         areaheights[areano] = -1

#         if areaheights[areano] == self.too_big:
#             areaheights[areano] = -1  # no boundaries, won't hold.
#             # the ones for which it leaks to can stay the same

#         if areaheights[areano] != -1:
#             q = leaks_to[areano][:]
#             while (len(q)) != 0:
#                 top = q[0]
#                 q = q[1:]
#                 areaheights[top] = areaheights[areano]
#                 for nxt in leaks_to[top]:
#                     q.append(nxt)


# TODO: leaking in two areas, but only one is completely falling. the second
# area may still be full

# l = [[1, 2], [55, 12, 34]]
# p = l[1]
# p.sort()
# print(l)

s = Solution()
print(
    "can fill up to",
    s.trapRainWater(
        [
            [3, 3, 3, 3, 3],
            [3, 2, 2, 2, 3],
            [3, 2, 1, 2, 3],
            [3, 2, 2, 2, 3],
            [3, 3, 3, 3, 3],
        ]
    ),
)
print(
    "can fill up to",
    s.trapRainWater([[1, 4, 3, 1, 3, 2], [3, 2, 1, 3, 2, 4], [2, 3, 3, 2, 3, 1]]),
)
print(
    "can fill up to",
    s.trapRainWater([[1, 3, 3, 1, 3, 2], [3, 2, 1, 3, 2, 3], [3, 3, 3, 2, 3, 1]]),
)
print(
    "can fill up to",
    s.trapRainWater(
        [
            [12, 13, 1, 12],
            [13, 4, 13, 12],
            [13, 8, 10, 12],
            [12, 13, 12, 12],
            [13, 13, 13, 13],
        ]
    ),
)
print(
    "can fill up to",
    s.trapRainWater([[2, 3, 4], [5, 6, 7], [8, 9, 10], [11, 12, 13], [14, 15, 16]]),
)
print(
    "can fill up to",
    s.trapRainWater(
        [
            [9, 9, 9, 9, 9],
            [9, 2, 1, 2, 8],
            [9, 2, 6, 2, 8],
            [9, 2, 3, 2, 8],
            [9, 8, 8, 7, 8],
        ]
    ),
)
print(
    "can fill up to",
    s.trapRainWater(
        [
            [78, 16, 94, 36],
            [87, 93, 50, 22],
            [63, 28, 91, 60],
            [64, 27, 41, 27],
            [73, 37, 12, 69],
            [68, 30, 83, 31],
            [63, 24, 68, 36],
        ]
    ),
)


# class Solution(object):
#     area_number_generator = 0
#     too_big = 99999999999999999999

#     def trapRainWater(self, heightMap):
#         """
#         :type heightMap: List[List[int]]
#         :rtype: int
#         """

#         self.area_number_generator = 0

#         LL, CC = len(heightMap), len(heightMap[0])
#         visited = set()

#         # areas and the maximum height of that area
#         areas = {}
#         areaheights = {}

#         leaks_to = defaultdict(list)

#         for i in range(LL):
#             for j in range(CC):
#                 self.area_number_generator += 1
#                 areano = self.area_number_generator
#                 self.explore_area(
#                     areas,
#                     areaheights,
#                     areano,
#                     (i, j),
#                     visited,
#                     heightMap,
#                     LL,
#                     CC,
#                     leaks_to,
#                 )

#         print(areas, areaheights)
#         s = 0
#         for el in areas:
#             areano = areas[el]
#             i, j = el
#             s += max(0, areaheights[areano] - heightMap[i][j])
#         for i in range(self.area_number_generator + 1):
#             to_print = []
#             for el in areas:
#                 if areas[el] == i:
#                     to_print.append(el)
#             print(i, to_print)
#         print("leaks to", leaks_to)
#         return s

#     def explore_area(
#         self,
#         areas,
#         areaheights,
#         areano,
#         starting_pos,
#         visited,
#         heightMap,
#         LL,
#         CC,
#         leaks_to,
#     ):

#         # ignore if it's already been explored, means it's in a different
#         # area. check probably redundant
#         if starting_pos in visited:
#             return
#         # ignore if it's already in a different area, probably the only
#         # check needed
#         if starting_pos in areas:
#             return

#         q = deque()
#         i, j = starting_pos
#         h = heightMap[i][j]
#         q.append((i, j))

#         areaheights[areano] = self.too_big

#         while len(q) != 0:
#             # set the location to the current area
#             top = q.popleft()
#             l, c = top
#             visited.add((l, c))  # mark this as visited

#             areas[top] = areano

#             ns = [(l - 1, c), (l + 1, c), (l, c - 1), (l, c + 1)]

#             for n in ns:
#                 ll, cc = n
#                 if not ((0 <= ll < LL) and (0 <= cc < CC)):
#                     areaheights[areano] = -1  # will leak
#                     continue

#                 # cannot leak to the neighbour, use its height as boundary
#                 if heightMap[ll][cc] > h:
#                     areaheights[areano] = min(areaheights[areano], heightMap[ll][cc])
#                     continue

#                 # neighbour already visited, ignore it
#                 # however, you need to take care of the case where the neighbour
#                 # is already visited, but you could leak into it
#                 #
#                 # todo: also study a -> b -> c and a->c
#                 if n in visited:
#                     if heightMap[ll][cc] < h and n in areas:
#                         leaks_to[areano].append(areas[n])
#                         if areaheights[areas[n]] < h:
#                             areaheights[areano] = -1
#                     continue

#                 # regular neighbour
#                 if heightMap[ll][cc] == h:
#                     q.append(n)
#                 # will leak in neighbour
#                 elif heightMap[ll][cc] < h:
#                     next_area_no = -99
#                     # if not explored, explore
#                     if n not in areas:
#                         self.area_number_generator += 1
#                         next_area_no = self.area_number_generator
#                         self.explore_area(
#                             areas,
#                             areaheights,
#                             next_area_no,
#                             n,
#                             visited,
#                             heightMap,
#                             LL,
#                             CC,
#                             leaks_to,
#                         )
#                     else:
#                         next_area_no = areas[n]

#                     leaks_to[areano].append(next_area_no)

#                     # now that it's 100% explored, add to the leaky relationship
#                     # this means, if the place it leaks into cannot hold to at
#                     # least the current level, set it as leaky
#                     if areaheights[next_area_no] < h:
#                         areaheights[areano] = -1

#         if areaheights[areano] == self.too_big:
#             areaheights[areano] = -1  # no boundaries, won't hold.
#             # the ones for which it leaks to can stay the same

#         if areaheights[areano] != -1:
#             q = leaks_to[areano][:]
#             while (len(q)) != 0:
#                 top = q[0]
#                 q = q[1:]
#                 areaheights[top] = areaheights[areano]
#                 for nxt in leaks_to[top]:
#                     q.append(nxt)
