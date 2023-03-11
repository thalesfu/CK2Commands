package job

import "github.com/thalesfu/CK2Commands/family/fu"

func GetMyloads() []int {
	return []int{0,
		fu.Me,
		fu.MyWife,
	}
}
