package pecheneg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉克塔什SaraktashBarony struct {
	feud.BaseBarony
}

var BSaraktash萨拉克塔什 feud.Barony = &萨拉克塔什SaraktashBarony{}

func init() {
    f := BSaraktash萨拉克塔什.(*萨拉克塔什SaraktashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saraktash",
		TitleName: "萨拉克塔什",
		TitleCode: "b_saraktash",
	}
}
