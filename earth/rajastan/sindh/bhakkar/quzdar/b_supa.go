package quzdar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏波SupaBarony struct {
	feud.BaseBarony
}

var BSupa苏波 feud.Barony = &苏波SupaBarony{}

func init() {
    f := BSupa苏波.(*苏波SupaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "supa",
		TitleName: "苏波",
		TitleCode: "b_supa",
	}
}
