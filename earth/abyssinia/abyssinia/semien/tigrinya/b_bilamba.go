package tigrinya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比拉巴BilambaBarony struct {
	feud.BaseBarony
}

var BBilamba比拉巴 feud.Barony = &比拉巴BilambaBarony{}

func init() {
	f := BBilamba比拉巴.(*比拉巴BilambaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilamba",
		TitleName: "比拉巴",
		TitleCode: "b_bilamba",
	}
}
