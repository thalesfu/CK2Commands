package massawa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马塔拉MataraBarony struct {
	feud.BaseBarony
}

var BMatara马塔拉 feud.Barony = &马塔拉MataraBarony{}

func init() {
    f := BMatara马塔拉.(*马塔拉MataraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "matara",
		TitleName: "马塔拉",
		TitleCode: "b_matara",
	}
}
