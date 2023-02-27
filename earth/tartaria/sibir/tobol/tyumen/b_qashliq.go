package tyumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀什里克QashliqBarony struct {
	feud.BaseBarony
}

var BQashliq喀什里克 feud.Barony = &喀什里克QashliqBarony{}

func init() {
    f := BQashliq喀什里克.(*喀什里克QashliqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qashliq",
		TitleName: "喀什里克",
		TitleCode: "b_qashliq",
	}
}
