package orava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳迈斯托NamesztoBarony struct {
	feud.BaseBarony
}

var BNameszto纳迈斯托 feud.Barony = &纳迈斯托NamesztoBarony{}

func init() {
    f := BNameszto纳迈斯托.(*纳迈斯托NamesztoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nameszto",
		TitleName: "纳迈斯托",
		TitleCode: "b_nameszto",
	}
}
