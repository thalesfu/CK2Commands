package matamma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉什阿拜GishabayBarony struct {
	feud.BaseBarony
}

var BGishabay吉什阿拜 feud.Barony = &吉什阿拜GishabayBarony{}

func init() {
	f := BGishabay吉什阿拜.(*吉什阿拜GishabayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gishabay",
		TitleName: "吉什阿拜",
		TitleCode: "b_gishabay",
	}
}
