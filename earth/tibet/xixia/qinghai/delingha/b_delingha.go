package delingha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德令哈DelinghaBarony struct {
	feud.BaseBarony
}

var BDelingha德令哈 feud.Barony = &德令哈DelinghaBarony{}

func init() {
    f := BDelingha德令哈.(*德令哈DelinghaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "delingha",
		TitleName: "德令哈",
		TitleCode: "b_delingha",
	}
}
