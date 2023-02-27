package ryn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坚季克TendikBarony struct {
	feud.BaseBarony
}

var BTendik坚季克 feud.Barony = &坚季克TendikBarony{}

func init() {
    f := BTendik坚季克.(*坚季克TendikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tendik",
		TitleName: "坚季克",
		TitleCode: "b_tendik",
	}
}
