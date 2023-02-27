package smyrna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 士麦拿SmyrnaBarony struct {
	feud.BaseBarony
}

var BSmyrna士麦拿 feud.Barony = &士麦拿SmyrnaBarony{}

func init() {
    f := BSmyrna士麦拿.(*士麦拿SmyrnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "smyrna",
		TitleName: "士麦拿",
		TitleCode: "b_smyrna",
	}
}
