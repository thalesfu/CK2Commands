package tarragona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉戈纳TarragonaBarony struct {
	feud.BaseBarony
}

var BTarragona塔拉戈纳 feud.Barony = &塔拉戈纳TarragonaBarony{}

func init() {
    f := BTarragona塔拉戈纳.(*塔拉戈纳TarragonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarragona",
		TitleName: "塔拉戈纳",
		TitleCode: "b_tarragona",
	}
}
