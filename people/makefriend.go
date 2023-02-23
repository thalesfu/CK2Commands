package people

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/CK2Commands/utils"
	"log"
)

func MakeFriends(group ...[]int) {
	file, err := utils.OpenFile("makefriends")
	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for gIndex, peoples := range group {
		peoples = utils.Filter[int](peoples, func(people int) bool {
			return people != 0
		})
		for i := 0; i < len(peoples); i++ {
			for j := i + 1; j < len(peoples); j++ {
				writer.WriteString(fmt.Sprintf("add_friend %d %d", peoples[i], peoples[j]))

				if !(gIndex == len(group)-1 && i == len(peoples)-2 && j == len(peoples)-1) {
					writer.WriteString("\n")
				}
			}
		}
	}

	writer.Flush()
}
