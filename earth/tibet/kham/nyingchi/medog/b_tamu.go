package medog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达木TamuBarony struct {
	feud.BaseBarony
}

var BTamu达木 feud.Barony = &达木TamuBarony{}

func init() {
    f := BTamu达木.(*达木TamuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamu",
		TitleName: "达木",
		TitleCode: "b_tamu",
	}
}
