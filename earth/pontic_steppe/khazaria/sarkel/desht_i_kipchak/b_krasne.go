package desht_i_kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉斯内KrasneBarony struct {
	feud.BaseBarony
}

var BKrasne克拉斯内 feud.Barony = &克拉斯内KrasneBarony{}

func init() {
    f := BKrasne克拉斯内.(*克拉斯内KrasneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasne",
		TitleName: "克拉斯内",
		TitleCode: "b_krasne",
	}
}
