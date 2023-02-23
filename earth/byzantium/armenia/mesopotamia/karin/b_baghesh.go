package karin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比特利斯BagheshBarony struct {
	feud.BaseBarony
}

var BBaghesh比特利斯 feud.Barony = &比特利斯BagheshBarony{}

func init() {
	f := BBaghesh比特利斯.(*比特利斯BagheshBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baghesh",
		TitleName: "比特利斯",
		TitleCode: "b_baghesh",
	}
}
