package limousin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂雷纳TurenneBarony struct {
	feud.BaseBarony
}

var BTurenne蒂雷纳 feud.Barony = &蒂雷纳TurenneBarony{}

func init() {
    f := BTurenne蒂雷纳.(*蒂雷纳TurenneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turenne",
		TitleName: "蒂雷纳",
		TitleCode: "b_turenne",
	}
}
