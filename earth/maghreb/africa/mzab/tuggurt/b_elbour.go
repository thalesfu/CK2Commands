package tuggurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔ElbourBarony struct {
	feud.BaseBarony
}

var BElbour布尔 feud.Barony = &布尔ElbourBarony{}

func init() {
	f := BElbour布尔.(*布尔ElbourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elbour",
		TitleName: "布尔",
		TitleCode: "b_elbour",
	}
}
