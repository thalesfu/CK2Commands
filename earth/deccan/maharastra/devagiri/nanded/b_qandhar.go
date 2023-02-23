package nanded

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎达哈QandharBarony struct {
	feud.BaseBarony
}

var BQandhar坎达哈 feud.Barony = &坎达哈QandharBarony{}

func init() {
	f := BQandhar坎达哈.(*坎达哈QandharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qandhar",
		TitleName: "坎达哈",
		TitleCode: "b_qandhar",
	}
}
