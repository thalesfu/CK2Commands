package tuul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆Lun_tuulBarony struct {
	feud.BaseBarony
}

var BLun_tuul隆 feud.Barony = &隆Lun_tuulBarony{}

func init() {
    f := BLun_tuul隆.(*隆Lun_tuulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lun_tuul",
		TitleName: "隆",
		TitleCode: "b_lun_tuul",
	}
}
