package khokh_serkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图噶勒TugalBarony struct {
	feud.BaseBarony
}

var BTugal图噶勒 feud.Barony = &图噶勒TugalBarony{}

func init() {
    f := BTugal图噶勒.(*图噶勒TugalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tugal",
		TitleName: "图噶勒",
		TitleCode: "b_tugal",
	}
}
