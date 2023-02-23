package lhatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曲下QuxarBarony struct {
	feud.BaseBarony
}

var BQuxar曲下 feud.Barony = &曲下QuxarBarony{}

func init() {
	f := BQuxar曲下.(*曲下QuxarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quxar",
		TitleName: "曲下",
		TitleCode: "b_quxar",
	}
}
