package kara_khorum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆LunBarony struct {
	feud.BaseBarony
}

var BLun隆 feud.Barony = &隆LunBarony{}

func init() {
    f := BLun隆.(*隆LunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lun",
		TitleName: "隆",
		TitleCode: "b_lun",
	}
}
