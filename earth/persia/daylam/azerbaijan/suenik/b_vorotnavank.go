package suenik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃罗特纳万克VorotnavankBarony struct {
	feud.BaseBarony
}

var BVorotnavank沃罗特纳万克 feud.Barony = &沃罗特纳万克VorotnavankBarony{}

func init() {
	f := BVorotnavank沃罗特纳万克.(*沃罗特纳万克VorotnavankBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vorotnavank",
		TitleName: "沃罗特纳万克",
		TitleCode: "b_vorotnavank",
	}
}
