package blekinge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕科LyckaBarony struct {
	feud.BaseBarony
}

var BLycka吕科 feud.Barony = &吕科LyckaBarony{}

func init() {
	f := BLycka吕科.(*吕科LyckaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lycka",
		TitleName: "吕科",
		TitleCode: "b_lycka",
	}
}
