package leix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂马胡TimahoeBarony struct {
	feud.BaseBarony
}

var BTimahoe蒂马胡 feud.Barony = &蒂马胡TimahoeBarony{}

func init() {
	f := BTimahoe蒂马胡.(*蒂马胡TimahoeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timahoe",
		TitleName: "蒂马胡",
		TitleCode: "b_timahoe",
	}
}
