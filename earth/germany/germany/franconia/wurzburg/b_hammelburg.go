package wurzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈默尔堡HammelburgBarony struct {
	feud.BaseBarony
}

var BHammelburg哈默尔堡 feud.Barony = &哈默尔堡HammelburgBarony{}

func init() {
	f := BHammelburg哈默尔堡.(*哈默尔堡HammelburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hammelburg",
		TitleName: "哈默尔堡",
		TitleCode: "b_hammelburg",
	}
}
