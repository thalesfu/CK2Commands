package weimar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉GeraBarony struct {
	feud.BaseBarony
}

var BGera格拉 feud.Barony = &格拉GeraBarony{}

func init() {
	f := BGera格拉.(*格拉GeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gera",
		TitleName: "格拉",
		TitleCode: "b_gera",
	}
}
