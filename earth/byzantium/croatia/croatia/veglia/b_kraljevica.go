package veglia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉列维察KraljevicaBarony struct {
	feud.BaseBarony
}

var BKraljevica克拉列维察 feud.Barony = &克拉列维察KraljevicaBarony{}

func init() {
	f := BKraljevica克拉列维察.(*克拉列维察KraljevicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kraljevica",
		TitleName: "克拉列维察",
		TitleCode: "b_kraljevica",
	}
}
