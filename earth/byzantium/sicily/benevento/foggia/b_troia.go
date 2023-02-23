package foggia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗亚TroiaBarony struct {
	feud.BaseBarony
}

var BTroia特罗亚 feud.Barony = &特罗亚TroiaBarony{}

func init() {
	f := BTroia特罗亚.(*特罗亚TroiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "troia",
		TitleName: "特罗亚",
		TitleCode: "b_troia",
	}
}
