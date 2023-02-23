package air

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库费KoufeBarony struct {
	feud.BaseBarony
}

var BKoufe库费 feud.Barony = &库费KoufeBarony{}

func init() {
	f := BKoufe库费.(*库费KoufeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koufe",
		TitleName: "库费",
		TitleCode: "b_koufe",
	}
}
