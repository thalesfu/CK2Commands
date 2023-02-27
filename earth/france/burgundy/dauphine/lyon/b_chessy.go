package lyon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢西ChessyBarony struct {
	feud.BaseBarony
}

var BChessy谢西 feud.Barony = &谢西ChessyBarony{}

func init() {
    f := BChessy谢西.(*谢西ChessyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chessy",
		TitleName: "谢西",
		TitleCode: "b_chessy",
	}
}
