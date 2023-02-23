package toulouse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉沃尔LavaurBarony struct {
	feud.BaseBarony
}

var BLavaur拉沃尔 feud.Barony = &拉沃尔LavaurBarony{}

func init() {
	f := BLavaur拉沃尔.(*拉沃尔LavaurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lavaur",
		TitleName: "拉沃尔",
		TitleCode: "b_lavaur",
	}
}
