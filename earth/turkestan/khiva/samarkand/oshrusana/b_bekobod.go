package oshrusana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝科博德BekobodBarony struct {
	feud.BaseBarony
}

var BBekobod贝科博德 feud.Barony = &贝科博德BekobodBarony{}

func init() {
    f := BBekobod贝科博德.(*贝科博德BekobodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bekobod",
		TitleName: "贝科博德",
		TitleCode: "b_bekobod",
	}
}
