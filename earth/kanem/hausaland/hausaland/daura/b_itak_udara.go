package daura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊塔克乌达拉Itak_udaraBarony struct {
	feud.BaseBarony
}

var BItak_udara伊塔克乌达拉 feud.Barony = &伊塔克乌达拉Itak_udaraBarony{}

func init() {
    f := BItak_udara伊塔克乌达拉.(*伊塔克乌达拉Itak_udaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "itak_udara",
		TitleName: "伊塔克乌达拉",
		TitleCode: "b_itak_udara",
	}
}
