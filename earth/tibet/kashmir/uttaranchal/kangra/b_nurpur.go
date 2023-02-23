package kangra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努尔布尔NurpurBarony struct {
	feud.BaseBarony
}

var BNurpur努尔布尔 feud.Barony = &努尔布尔NurpurBarony{}

func init() {
	f := BNurpur努尔布尔.(*努尔布尔NurpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nurpur",
		TitleName: "努尔布尔",
		TitleCode: "b_nurpur",
	}
}
