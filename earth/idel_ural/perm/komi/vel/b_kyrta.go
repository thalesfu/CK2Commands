package vel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔塔KyrtaBarony struct {
	feud.BaseBarony
}

var BKyrta克尔塔 feud.Barony = &克尔塔KyrtaBarony{}

func init() {
    f := BKyrta克尔塔.(*克尔塔KyrtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyrta",
		TitleName: "克尔塔",
		TitleCode: "b_kyrta",
	}
}
