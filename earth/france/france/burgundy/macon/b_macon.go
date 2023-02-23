package macon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马孔MaconBarony struct {
	feud.BaseBarony
}

var BMacon马孔 feud.Barony = &马孔MaconBarony{}

func init() {
	f := BMacon马孔.(*马孔MaconBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "macon",
		TitleName: "马孔",
		TitleCode: "b_macon",
	}
}
