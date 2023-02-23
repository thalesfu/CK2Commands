package taza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尼亚ManiaBarony struct {
	feud.BaseBarony
}

var BMania马尼亚 feud.Barony = &马尼亚ManiaBarony{}

func init() {
	f := BMania马尼亚.(*马尼亚ManiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mania",
		TitleName: "马尼亚",
		TitleCode: "b_mania",
	}
}
