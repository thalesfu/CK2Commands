package vestisland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊萨菲厄泽IsafjordurBarony struct {
	feud.BaseBarony
}

var BIsafjordur伊萨菲厄泽 feud.Barony = &伊萨菲厄泽IsafjordurBarony{}

func init() {
    f := BIsafjordur伊萨菲厄泽.(*伊萨菲厄泽IsafjordurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isafjordur",
		TitleName: "伊萨菲厄泽",
		TitleCode: "b_isafjordur",
	}
}
