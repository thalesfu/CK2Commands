package vestisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔塔菲厄泽AlftafjordurBarony struct {
	feud.BaseBarony
}

var BAlftafjordur奥尔塔菲厄泽 feud.Barony = &奥尔塔菲厄泽AlftafjordurBarony{}

func init() {
    f := BAlftafjordur奥尔塔菲厄泽.(*奥尔塔菲厄泽AlftafjordurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alftafjordur",
		TitleName: "奥尔塔菲厄泽",
		TitleCode: "b_alftafjordur",
	}
}
