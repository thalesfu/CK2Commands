package zamora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨纳夫里亚SanabriaBarony struct {
	feud.BaseBarony
}

var BSanabria萨纳夫里亚 feud.Barony = &萨纳夫里亚SanabriaBarony{}

func init() {
    f := BSanabria萨纳夫里亚.(*萨纳夫里亚SanabriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanabria",
		TitleName: "萨纳夫里亚",
		TitleCode: "b_sanabria",
	}
}
