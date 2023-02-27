package kanj_rustaq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥法AufahBarony struct {
	feud.BaseBarony
}

var BAufah奥法 feud.Barony = &奥法AufahBarony{}

func init() {
    f := BAufah奥法.(*奥法AufahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aufah",
		TitleName: "奥法",
		TitleCode: "b_aufah",
	}
}
