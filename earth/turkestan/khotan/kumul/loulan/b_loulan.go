package loulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楼兰LoulanBarony struct {
	feud.BaseBarony
}

var BLoulan楼兰 feud.Barony = &楼兰LoulanBarony{}

func init() {
    f := BLoulan楼兰.(*楼兰LoulanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loulan",
		TitleName: "楼兰",
		TitleCode: "b_loulan",
	}
}
