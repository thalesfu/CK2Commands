package anxi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新城XinchengBarony struct {
	feud.BaseBarony
}

var BXincheng新城 feud.Barony = &新城XinchengBarony{}

func init() {
	f := BXincheng新城.(*新城XinchengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xincheng",
		TitleName: "新城",
		TitleCode: "b_xincheng",
	}
}
