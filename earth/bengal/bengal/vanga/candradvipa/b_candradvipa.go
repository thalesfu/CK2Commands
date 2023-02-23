package candradvipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 月洲CandradvipaBarony struct {
	feud.BaseBarony
}

var BCandradvipa月洲 feud.Barony = &月洲CandradvipaBarony{}

func init() {
	f := BCandradvipa月洲.(*月洲CandradvipaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "candradvipa",
		TitleName: "月洲",
		TitleCode: "b_candradvipa",
	}
}
