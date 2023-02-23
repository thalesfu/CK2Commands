package tripolitana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福杜格FunduqBarony struct {
	feud.BaseBarony
}

var BFunduq福杜格 feud.Barony = &福杜格FunduqBarony{}

func init() {
	f := BFunduq福杜格.(*福杜格FunduqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "funduq",
		TitleName: "福杜格",
		TitleCode: "b_funduq",
	}
}
