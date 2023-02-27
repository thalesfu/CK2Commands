package nagormo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 清水QingshuiBarony struct {
	feud.BaseBarony
}

var BQingshui清水 feud.Barony = &清水QingshuiBarony{}

func init() {
    f := BQingshui清水.(*清水QingshuiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qingshui",
		TitleName: "清水",
		TitleCode: "b_qingshui",
	}
}
