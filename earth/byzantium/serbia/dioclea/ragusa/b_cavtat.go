package ragusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察夫塔特CavtatBarony struct {
	feud.BaseBarony
}

var BCavtat察夫塔特 feud.Barony = &察夫塔特CavtatBarony{}

func init() {
	f := BCavtat察夫塔特.(*察夫塔特CavtatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cavtat",
		TitleName: "察夫塔特",
		TitleCode: "b_cavtat",
	}
}
