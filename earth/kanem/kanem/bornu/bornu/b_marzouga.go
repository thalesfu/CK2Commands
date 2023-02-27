package bornu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔祖加MarzougaBarony struct {
	feud.BaseBarony
}

var BMarzouga梅尔祖加 feud.Barony = &梅尔祖加MarzougaBarony{}

func init() {
    f := BMarzouga梅尔祖加.(*梅尔祖加MarzougaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marzouga",
		TitleName: "梅尔祖加",
		TitleCode: "b_marzouga",
	}
}
