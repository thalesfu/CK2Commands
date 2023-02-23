package leventina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法埃多FaedoBarony struct {
	feud.BaseBarony
}

var BFaedo法埃多 feud.Barony = &法埃多FaedoBarony{}

func init() {
	f := BFaedo法埃多.(*法埃多FaedoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faedo",
		TitleName: "法埃多",
		TitleCode: "b_faedo",
	}
}
