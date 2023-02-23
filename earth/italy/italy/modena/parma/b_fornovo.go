package parma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福尔诺沃FornovoBarony struct {
	feud.BaseBarony
}

var BFornovo福尔诺沃 feud.Barony = &福尔诺沃FornovoBarony{}

func init() {
	f := BFornovo福尔诺沃.(*福尔诺沃FornovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fornovo",
		TitleName: "福尔诺沃",
		TitleCode: "b_fornovo",
	}
}
