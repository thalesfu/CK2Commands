package birjand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内赫班丹NehbandanBarony struct {
	feud.BaseBarony
}

var BNehbandan内赫班丹 feud.Barony = &内赫班丹NehbandanBarony{}

func init() {
	f := BNehbandan内赫班丹.(*内赫班丹NehbandanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nehbandan",
		TitleName: "内赫班丹",
		TitleCode: "b_nehbandan",
	}
}
