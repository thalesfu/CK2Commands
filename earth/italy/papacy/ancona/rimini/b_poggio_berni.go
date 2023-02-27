package rimini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波焦贝尔尼Poggio_berniBarony struct {
	feud.BaseBarony
}

var BPoggio_berni波焦贝尔尼 feud.Barony = &波焦贝尔尼Poggio_berniBarony{}

func init() {
    f := BPoggio_berni波焦贝尔尼.(*波焦贝尔尼Poggio_berniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poggio_berni",
		TitleName: "波焦贝尔尼",
		TitleCode: "b_poggio_berni",
	}
}
