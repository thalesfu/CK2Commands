package achaia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔季察KarditzaBarony struct {
	feud.BaseBarony
}

var BKarditza卡尔季察 feud.Barony = &卡尔季察KarditzaBarony{}

func init() {
    f := BKarditza卡尔季察.(*卡尔季察KarditzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karditza",
		TitleName: "卡尔季察",
		TitleCode: "b_karditza",
	}
}
