package manupura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈哈莱MalhallaBarony struct {
	feud.BaseBarony
}

var BMalhalla迈哈莱 feud.Barony = &迈哈莱MalhallaBarony{}

func init() {
    f := BMalhalla迈哈莱.(*迈哈莱MalhallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malhalla",
		TitleName: "迈哈莱",
		TitleCode: "b_malhalla",
	}
}
