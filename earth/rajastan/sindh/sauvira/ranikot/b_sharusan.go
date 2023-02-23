package ranikot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍楼桑SharusanBarony struct {
	feud.BaseBarony
}

var BSharusan舍楼桑 feud.Barony = &舍楼桑SharusanBarony{}

func init() {
	f := BSharusan舍楼桑.(*舍楼桑SharusanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharusan",
		TitleName: "舍楼桑",
		TitleCode: "b_sharusan",
	}
}
