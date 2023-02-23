package almansa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托瓦拉TobarraBarony struct {
	feud.BaseBarony
}

var BTobarra托瓦拉 feud.Barony = &托瓦拉TobarraBarony{}

func init() {
	f := BTobarra托瓦拉.(*托瓦拉TobarraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tobarra",
		TitleName: "托瓦拉",
		TitleCode: "b_tobarra",
	}
}
