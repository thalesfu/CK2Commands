package mecca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡鲁纳QarnBarony struct {
	feud.BaseBarony
}

var BQarn卡鲁纳 feud.Barony = &卡鲁纳QarnBarony{}

func init() {
	f := BQarn卡鲁纳.(*卡鲁纳QarnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qarn",
		TitleName: "卡鲁纳",
		TitleCode: "b_qarn",
	}
}
