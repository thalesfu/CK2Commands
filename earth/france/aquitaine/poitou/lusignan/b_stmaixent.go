package lusignan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣迈克桑StmaixentBarony struct {
	feud.BaseBarony
}

var BStmaixent圣迈克桑 feud.Barony = &圣迈克桑StmaixentBarony{}

func init() {
	f := BStmaixent圣迈克桑.(*圣迈克桑StmaixentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stmaixent",
		TitleName: "圣迈克桑",
		TitleCode: "b_stmaixent",
	}
}
