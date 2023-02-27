package agder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 许勒斯塔HylestadBarony struct {
	feud.BaseBarony
}

var BHylestad许勒斯塔 feud.Barony = &许勒斯塔HylestadBarony{}

func init() {
    f := BHylestad许勒斯塔.(*许勒斯塔HylestadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hylestad",
		TitleName: "许勒斯塔",
		TitleCode: "b_hylestad",
	}
}
