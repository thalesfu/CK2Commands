package tadjrift

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉舍雷GicherraBarony struct {
	feud.BaseBarony
}

var BGicherra吉舍雷 feud.Barony = &吉舍雷GicherraBarony{}

func init() {
    f := BGicherra吉舍雷.(*吉舍雷GicherraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gicherra",
		TitleName: "吉舍雷",
		TitleCode: "b_gicherra",
	}
}
