package gelre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安亨ArnhemBarony struct {
	feud.BaseBarony
}

var BArnhem安亨 feud.Barony = &安亨ArnhemBarony{}

func init() {
	f := BArnhem安亨.(*安亨ArnhemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arnhem",
		TitleName: "安亨",
		TitleCode: "b_arnhem",
	}
}
