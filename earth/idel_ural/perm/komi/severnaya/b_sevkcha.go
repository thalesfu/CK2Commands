package severnaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢夫克恰SevkchaBarony struct {
	feud.BaseBarony
}

var BSevkcha谢夫克恰 feud.Barony = &谢夫克恰SevkchaBarony{}

func init() {
    f := BSevkcha谢夫克恰.(*谢夫克恰SevkchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sevkcha",
		TitleName: "谢夫克恰",
		TitleCode: "b_sevkcha",
	}
}
