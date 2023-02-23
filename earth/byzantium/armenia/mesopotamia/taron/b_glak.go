package taron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉卡GlakBarony struct {
	feud.BaseBarony
}

var BGlak格拉卡 feud.Barony = &格拉卡GlakBarony{}

func init() {
	f := BGlak格拉卡.(*格拉卡GlakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glak",
		TitleName: "格拉卡",
		TitleCode: "b_glak",
	}
}
