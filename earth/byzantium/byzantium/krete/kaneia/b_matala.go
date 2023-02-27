package kaneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马塔拉MatalaBarony struct {
	feud.BaseBarony
}

var BMatala马塔拉 feud.Barony = &马塔拉MatalaBarony{}

func init() {
    f := BMatala马塔拉.(*马塔拉MatalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "matala",
		TitleName: "马塔拉",
		TitleCode: "b_matala",
	}
}
