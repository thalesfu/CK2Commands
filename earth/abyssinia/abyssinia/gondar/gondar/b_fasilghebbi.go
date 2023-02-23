package gondar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法西尔盖比FasilghebbiBarony struct {
	feud.BaseBarony
}

var BFasilghebbi法西尔盖比 feud.Barony = &法西尔盖比FasilghebbiBarony{}

func init() {
	f := BFasilghebbi法西尔盖比.(*法西尔盖比FasilghebbiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fasilghebbi",
		TitleName: "法西尔盖比",
		TitleCode: "b_fasilghebbi",
	}
}
