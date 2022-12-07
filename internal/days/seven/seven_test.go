package seven

import "testing"

var input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

var expectedDirSizes = map[string]int64{
	"/": 48381165,
	"d": 24933642,
	"a": 94853,
	"e": 584,
}

func TestProcessFileSystem(t *testing.T) {
	rootDir, allDirs := processFileSystem(input)
	rootDir.Print(0)
	if len(allDirs) != 4 {
		t.Fatalf("should be 4 dirs were %d", len(allDirs))
	}
	for _, d := range allDirs {
		expectedSize := expectedDirSizes[d.Path]
		if expectedSize != d.RecursiveSize {
			t.Fatalf("%s size should be %d but was %d", d.Path, expectedSize, d.RecursiveSize)
		}

	}

}
