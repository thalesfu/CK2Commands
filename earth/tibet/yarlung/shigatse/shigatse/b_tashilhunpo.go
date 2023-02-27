package shigatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎什伦布TashilhunpoBarony struct {
	feud.BaseBarony
}

var BTashilhunpo扎什伦布 feud.Barony = &扎什伦布TashilhunpoBarony{}

func init() {
    f := BTashilhunpo扎什伦布.(*扎什伦布TashilhunpoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tashilhunpo",
		TitleName: "扎什伦布",
		TitleCode: "b_tashilhunpo",
	}
}
