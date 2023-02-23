package bukhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布哈拉BukharaBarony struct {
	feud.BaseBarony
}

var BBukhara布哈拉 feud.Barony = &布哈拉BukharaBarony{}

func init() {
	f := BBukhara布哈拉.(*布哈拉BukharaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bukhara",
		TitleName: "布哈拉",
		TitleCode: "b_bukhara",
	}
}
