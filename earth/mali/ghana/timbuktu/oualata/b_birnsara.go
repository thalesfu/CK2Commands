package oualata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔恩萨拉BirnsaraBarony struct {
	feud.BaseBarony
}

var BBirnsara比尔恩萨拉 feud.Barony = &比尔恩萨拉BirnsaraBarony{}

func init() {
	f := BBirnsara比尔恩萨拉.(*比尔恩萨拉BirnsaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birnsara",
		TitleName: "比尔恩萨拉",
		TitleCode: "b_birnsara",
	}
}
