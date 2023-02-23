package lahur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉合尔LahurBarony struct {
	feud.BaseBarony
}

var BLahur拉合尔 feud.Barony = &拉合尔LahurBarony{}

func init() {
	f := BLahur拉合尔.(*拉合尔LahurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lahur",
		TitleName: "拉合尔",
		TitleCode: "b_lahur",
	}
}
