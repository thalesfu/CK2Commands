package zamindawar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈莱KhalaiBarony struct {
	feud.BaseBarony
}

var BKhalai哈莱 feud.Barony = &哈莱KhalaiBarony{}

func init() {
    f := BKhalai哈莱.(*哈莱KhalaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khalai",
		TitleName: "哈莱",
		TitleCode: "b_khalai",
	}
}
