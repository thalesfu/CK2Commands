package el_rif

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克塔玛KetamaBarony struct {
	feud.BaseBarony
}

var BKetama克塔玛 feud.Barony = &克塔玛KetamaBarony{}

func init() {
    f := BKetama克塔玛.(*克塔玛KetamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ketama",
		TitleName: "克塔玛",
		TitleCode: "b_ketama",
	}
}
