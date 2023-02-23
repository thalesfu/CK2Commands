package barskhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂普TyupBarony struct {
	feud.BaseBarony
}

var BTyup蒂普 feud.Barony = &蒂普TyupBarony{}

func init() {
	f := BTyup蒂普.(*蒂普TyupBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyup",
		TitleName: "蒂普",
		TitleCode: "b_tyup",
	}
}
