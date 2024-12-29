package day9

import (
	"adventofcode-2025/common"
	"fmt"
	"strconv"
)

type diskBlock struct {
	isFreeSpace bool
	fileId      int
}

type blockInfo struct {
	initialPosition, size int
}

func parseDiskMap(lines []string) (diskMap []diskBlock, freeSpaceBlocks map[int]blockInfo, fileBlocks map[int]blockInfo) {
	rows := len(lines)

	if rows > 1 {
		panic("More than one disk map not allowed")
	}

	fileId := 0
	line := lines[0]
	diskMap = make([]diskBlock, 0)
	freeSpaceBlocks = make(map[int]blockInfo)
	fsi := 0
	fileBlocks = make(map[int]blockInfo)
	fbi := 0
	for index, char := range line {
		isFreeSpace := (index+1)%2 == 0
		size := common.ParseInt(char)
		for i := 0; i < size; i++ {
			if isFreeSpace {
				diskMap = append(diskMap, diskBlock{isFreeSpace: true, fileId: -1})
			} else {
				diskMap = append(diskMap, diskBlock{isFreeSpace: false, fileId: fileId})
			}
		}
		mapLen := len(diskMap) - size
		if !isFreeSpace {
			fileId++
			fileBlocks[fbi] = blockInfo{initialPosition: mapLen, size: size}
			fbi++
		} else {
			freeSpaceBlocks[fsi] = blockInfo{initialPosition: mapLen, size: size}
			fsi++
		}
	}

	return diskMap, freeSpaceBlocks, fileBlocks
}

func compactDiskMap(diskMap []diskBlock) []diskBlock {
	size := len(diskMap)

	for i := 0; i < size; i++ {
		freeBlock := diskMap[i]

		if !freeBlock.isFreeSpace {
			continue
		}

		for j := size - 1; j > i; j-- {
			right := diskMap[j]
			if right.isFreeSpace {
				continue
			}

			diskMap[i] = diskMap[j]
			diskMap[j] = freeBlock
			break
		}
	}

	return diskMap
}

func calculateChecksum(diskMap []diskBlock) (checksum int) {
	checksum = 0
	for i := 0; i < len(diskMap); i++ {
		block := diskMap[i]
		if block.isFreeSpace {
			continue
		}

		checksum += block.fileId * i
	}

	return checksum
}

func printDiskMap(diskMap []diskBlock) {
	for i := 0; i < len(diskMap); i++ {
		block := diskMap[i]
		if block.isFreeSpace {
			fmt.Print(".")
		} else {
			fmt.Print(strconv.Itoa(block.fileId))
		}
	}
	fmt.Println()
}

func Puzzle1() int {
	lines := common.ReadFileByLines("./inputs/day9.test3.txt")

	diskMap, _, _ := parseDiskMap(lines)

	compactedDiskMap := compactDiskMap(diskMap)

	checksum := calculateChecksum(compactedDiskMap)

	return checksum
}
