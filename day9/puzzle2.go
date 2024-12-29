package day9

import (
	"adventofcode-2025/common"
)

func compactDiskMapV2(diskMap []diskBlock, freeSpaceMap *map[int]blockInfo, fileMap *map[int]blockInfo) []diskBlock {

	for i := len(*fileMap) - 1; i > 0; i-- {
		fileBlock := (*fileMap)[i]

		for j := 0; j < i; j++ {
			freeSpaceBlock := (*freeSpaceMap)[j]

			if freeSpaceBlock.size < fileBlock.size {
				continue
			}

			for k := 0; k < fileBlock.size; k++ {
				diskMap[freeSpaceBlock.initialPosition], diskMap[fileBlock.initialPosition+k] =
					diskMap[fileBlock.initialPosition+k], diskMap[freeSpaceBlock.initialPosition]
				freeSpaceBlock.size--
				freeSpaceBlock.initialPosition++
			}

			(*freeSpaceMap)[j] = freeSpaceBlock

			break
		}
	}

	return diskMap
}

func Puzzle2() int {
	lines := common.ReadFileByLines("./inputs/day9.txt")

	diskMap, freeSpaceMap, fileMap := parseDiskMap(lines)

	printDiskMap(diskMap)

	compactedDiskMap := compactDiskMapV2(diskMap, &freeSpaceMap, &fileMap)

	printDiskMap(compactedDiskMap)

	checksum := calculateChecksum(compactedDiskMap)

	return checksum
}
