package trigarta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉亨RahonBarony struct {
	feud.BaseBarony
}

var BRahon拉亨 feud.Barony = &拉亨RahonBarony{}

func init() {
	f := BRahon拉亨.(*拉亨RahonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rahon",
		TitleName: "拉亨",
		TitleCode: "b_rahon",
	}
}
