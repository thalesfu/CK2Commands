package itil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图马克TumakBarony struct {
	feud.BaseBarony
}

var BTumak图马克 feud.Barony = &图马克TumakBarony{}

func init() {
    f := BTumak图马克.(*图马克TumakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tumak",
		TitleName: "图马克",
		TitleCode: "b_tumak",
	}
}
