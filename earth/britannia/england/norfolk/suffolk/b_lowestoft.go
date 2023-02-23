package suffolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛斯托夫特LowestoftBarony struct {
	feud.BaseBarony
}

var BLowestoft洛斯托夫特 feud.Barony = &洛斯托夫特LowestoftBarony{}

func init() {
	f := BLowestoft洛斯托夫特.(*洛斯托夫特LowestoftBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lowestoft",
		TitleName: "洛斯托夫特",
		TitleCode: "b_lowestoft",
	}
}
