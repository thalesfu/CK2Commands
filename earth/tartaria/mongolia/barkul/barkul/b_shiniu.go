package barkul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 石纽ShiniuBarony struct {
	feud.BaseBarony
}

var BShiniu石纽 feud.Barony = &石纽ShiniuBarony{}

func init() {
	f := BShiniu石纽.(*石纽ShiniuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shiniu",
		TitleName: "石纽",
		TitleCode: "b_shiniu",
	}
}
