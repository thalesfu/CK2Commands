package ormond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡舍尔CashelBarony struct {
	feud.BaseBarony
}

var BCashel卡舍尔 feud.Barony = &卡舍尔CashelBarony{}

func init() {
    f := BCashel卡舍尔.(*卡舍尔CashelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cashel",
		TitleName: "卡舍尔",
		TitleCode: "b_cashel",
	}
}
