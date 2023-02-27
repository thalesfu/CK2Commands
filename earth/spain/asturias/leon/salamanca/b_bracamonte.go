package salamanca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉卡蒙特BracamonteBarony struct {
	feud.BaseBarony
}

var BBracamonte布拉卡蒙特 feud.Barony = &布拉卡蒙特BracamonteBarony{}

func init() {
    f := BBracamonte布拉卡蒙特.(*布拉卡蒙特BracamonteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bracamonte",
		TitleName: "布拉卡蒙特",
		TitleCode: "b_bracamonte",
	}
}
