package day8

func Puzzle2() int {
	uniqueAntinodes := make(map[Coordinate]struct{}) // Clear for Part 2
	for _, antenna := range antennas {
		sameFrequencyAntennas := antennaToCoordinates[antenna.Freq]
		for _, coord := range sameFrequencyAntennas {
			if coord.Row == antenna.Coord.Row && coord.Col == antenna.Coord.Col {
				continue
			}

			delta := antenna.Coord.Sub(coord)
			potentialAntinodeA := antenna.Coord.Add(delta)
			potentialAntinodeB := antenna.Coord.Sub(delta)

			// Traverse in both directions until out of bounds
			for IsInGrid(potentialAntinodeA, rows, cols) {
				uniqueAntinodes[potentialAntinodeA] = struct{}{}
				potentialAntinodeA = potentialAntinodeA.Add(delta)
			}
			for IsInGrid(potentialAntinodeB, rows, cols) {
				uniqueAntinodes[potentialAntinodeB] = struct{}{}
				potentialAntinodeB = potentialAntinodeB.Sub(delta)
			}
		}
	}

	return len(uniqueAntinodes)
}
