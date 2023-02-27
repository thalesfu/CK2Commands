package hormuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米纳卜MinabBarony struct {
	feud.BaseBarony
}

var BMinab米纳卜 feud.Barony = &米纳卜MinabBarony{}

func init() {
    f := BMinab米纳卜.(*米纳卜MinabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "minab",
		TitleName: "米纳卜",
		TitleCode: "b_minab",
	}
}
