package trencin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日利纳ZilinaBarony struct {
	feud.BaseBarony
}

var BZilina日利纳 feud.Barony = &日利纳ZilinaBarony{}

func init() {
    f := BZilina日利纳.(*日利纳ZilinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zilina",
		TitleName: "日利纳",
		TitleCode: "b_zilina",
	}
}
