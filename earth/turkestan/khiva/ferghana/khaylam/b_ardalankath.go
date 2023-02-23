package khaylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿达兰卡ArdalankathBarony struct {
	feud.BaseBarony
}

var BArdalankath阿达兰卡 feud.Barony = &阿达兰卡ArdalankathBarony{}

func init() {
	f := BArdalankath阿达兰卡.(*阿达兰卡ArdalankathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardalankath",
		TitleName: "阿达兰卡",
		TitleCode: "b_ardalankath",
	}
}
