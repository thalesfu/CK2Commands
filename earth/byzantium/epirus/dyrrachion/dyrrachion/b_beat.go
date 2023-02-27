package dyrrachion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝特BeatBarony struct {
	feud.BaseBarony
}

var BBeat贝特 feud.Barony = &贝特BeatBarony{}

func init() {
    f := BBeat贝特.(*贝特BeatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beat",
		TitleName: "贝特",
		TitleCode: "b_beat",
	}
}
