package dunois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙托丹ChateaudunBarony struct {
	feud.BaseBarony
}

var BChateaudun沙托丹 feud.Barony = &沙托丹ChateaudunBarony{}

func init() {
    f := BChateaudun沙托丹.(*沙托丹ChateaudunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateaudun",
		TitleName: "沙托丹",
		TitleCode: "b_chateaudun",
	}
}
