package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉尔KalarBarony struct {
	feud.BaseBarony
}

var BKalar卡拉尔 feud.Barony = &卡拉尔KalarBarony{}

func init() {
    f := BKalar卡拉尔.(*卡拉尔KalarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalar",
		TitleName: "卡拉尔",
		TitleCode: "b_kalar",
	}
}
