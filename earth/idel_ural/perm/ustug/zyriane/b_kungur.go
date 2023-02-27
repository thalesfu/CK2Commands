package zyriane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆古尔KungurBarony struct {
	feud.BaseBarony
}

var BKungur昆古尔 feud.Barony = &昆古尔KungurBarony{}

func init() {
    f := BKungur昆古尔.(*昆古尔KungurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kungur",
		TitleName: "昆古尔",
		TitleCode: "b_kungur",
	}
}
