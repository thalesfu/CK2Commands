package tyrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊姆伯特城堡CasalimbertBarony struct {
	feud.BaseBarony
}

var BCasalimbert伊姆伯特城堡 feud.Barony = &伊姆伯特城堡CasalimbertBarony{}

func init() {
	f := BCasalimbert伊姆伯特城堡.(*伊姆伯特城堡CasalimbertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "casalimbert",
		TitleName: "伊姆伯特城堡",
		TitleCode: "b_casalimbert",
	}
}
