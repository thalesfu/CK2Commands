package yperen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁瑟拉勒RoeselareBarony struct {
	feud.BaseBarony
}

var BRoeselare鲁瑟拉勒 feud.Barony = &鲁瑟拉勒RoeselareBarony{}

func init() {
	f := BRoeselare鲁瑟拉勒.(*鲁瑟拉勒RoeselareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roeselare",
		TitleName: "鲁瑟拉勒",
		TitleCode: "b_roeselare",
	}
}
