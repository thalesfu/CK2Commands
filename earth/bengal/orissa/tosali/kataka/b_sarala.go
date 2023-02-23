package kataka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙罗拉SaralaBarony struct {
	feud.BaseBarony
}

var BSarala沙罗拉 feud.Barony = &沙罗拉SaralaBarony{}

func init() {
	f := BSarala沙罗拉.(*沙罗拉SaralaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarala",
		TitleName: "沙罗拉",
		TitleCode: "b_sarala",
	}
}
