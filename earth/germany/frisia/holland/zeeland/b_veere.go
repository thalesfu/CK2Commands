package zeeland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费勒VeereBarony struct {
	feud.BaseBarony
}

var BVeere费勒 feud.Barony = &费勒VeereBarony{}

func init() {
	f := BVeere费勒.(*费勒VeereBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veere",
		TitleName: "费勒",
		TitleCode: "b_veere",
	}
}
