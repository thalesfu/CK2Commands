package bira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代里克DerikBarony struct {
	feud.BaseBarony
}

var BDerik代里克 feud.Barony = &代里克DerikBarony{}

func init() {
	f := BDerik代里克.(*代里克DerikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derik",
		TitleName: "代里克",
		TitleCode: "b_derik",
	}
}
