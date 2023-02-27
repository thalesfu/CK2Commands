package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日贾尔ZdarBarony struct {
	feud.BaseBarony
}

var BZdar日贾尔 feud.Barony = &日贾尔ZdarBarony{}

func init() {
    f := BZdar日贾尔.(*日贾尔ZdarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zdar",
		TitleName: "日贾尔",
		TitleCode: "b_zdar",
	}
}
