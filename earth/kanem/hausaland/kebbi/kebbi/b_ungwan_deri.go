package kebbi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 翁万德里Ungwan_deriBarony struct {
	feud.BaseBarony
}

var BUngwan_deri翁万德里 feud.Barony = &翁万德里Ungwan_deriBarony{}

func init() {
    f := BUngwan_deri翁万德里.(*翁万德里Ungwan_deriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ungwan_deri",
		TitleName: "翁万德里",
		TitleCode: "b_ungwan_deri",
	}
}
