package romerike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟吕姆桑SorumsandBarony struct {
	feud.BaseBarony
}

var BSorumsand瑟吕姆桑 feud.Barony = &瑟吕姆桑SorumsandBarony{}

func init() {
    f := BSorumsand瑟吕姆桑.(*瑟吕姆桑SorumsandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sorumsand",
		TitleName: "瑟吕姆桑",
		TitleCode: "b_sorumsand",
	}
}
