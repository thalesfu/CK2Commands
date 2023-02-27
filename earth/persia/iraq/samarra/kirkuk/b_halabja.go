package kirkuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈莱卜杰HalabjaBarony struct {
	feud.BaseBarony
}

var BHalabja哈莱卜杰 feud.Barony = &哈莱卜杰HalabjaBarony{}

func init() {
    f := BHalabja哈莱卜杰.(*哈莱卜杰HalabjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halabja",
		TitleName: "哈莱卜杰",
		TitleCode: "b_halabja",
	}
}
