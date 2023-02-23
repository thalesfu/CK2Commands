package matamma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇利基ChiriBarony struct {
	feud.BaseBarony
}

var BChiri奇利基 feud.Barony = &奇利基ChiriBarony{}

func init() {
	f := BChiri奇利基.(*奇利基ChiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chiri",
		TitleName: "奇利基",
		TitleCode: "b_chiri",
	}
}
