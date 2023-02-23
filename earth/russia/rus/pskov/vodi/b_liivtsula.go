package vodi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利夫特许莱LiivtsulaBarony struct {
	feud.BaseBarony
}

var BLiivtsula利夫特许莱 feud.Barony = &利夫特许莱LiivtsulaBarony{}

func init() {
	f := BLiivtsula利夫特许莱.(*利夫特许莱LiivtsulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liivtsula",
		TitleName: "利夫特许莱",
		TitleCode: "b_liivtsula",
	}
}
