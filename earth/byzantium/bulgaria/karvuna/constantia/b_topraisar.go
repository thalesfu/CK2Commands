package constantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托普拉伊萨尔TopraisarBarony struct {
	feud.BaseBarony
}

var BTopraisar托普拉伊萨尔 feud.Barony = &托普拉伊萨尔TopraisarBarony{}

func init() {
	f := BTopraisar托普拉伊萨尔.(*托普拉伊萨尔TopraisarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "topraisar",
		TitleName: "托普拉伊萨尔",
		TitleCode: "b_topraisar",
	}
}
