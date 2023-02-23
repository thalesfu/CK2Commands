package taktse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉玉LayuBarony struct {
	feud.BaseBarony
}

var BLayu拉玉 feud.Barony = &拉玉LayuBarony{}

func init() {
	f := BLayu拉玉.(*拉玉LayuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "layu",
		TitleName: "拉玉",
		TitleCode: "b_layu",
	}
}
