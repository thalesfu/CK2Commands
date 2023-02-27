package damin_i_koh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆墯BhaduaBarony struct {
	feud.BaseBarony
}

var BBhadua婆墯 feud.Barony = &婆墯BhaduaBarony{}

func init() {
    f := BBhadua婆墯.(*婆墯BhaduaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhadua",
		TitleName: "婆墯",
		TitleCode: "b_bhadua",
	}
}
