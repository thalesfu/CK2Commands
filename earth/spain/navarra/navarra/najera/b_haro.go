package najera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿罗HaroBarony struct {
	feud.BaseBarony
}

var BHaro阿罗 feud.Barony = &阿罗HaroBarony{}

func init() {
	f := BHaro阿罗.(*阿罗HaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haro",
		TitleName: "阿罗",
		TitleCode: "b_haro",
	}
}
