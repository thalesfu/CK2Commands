package farama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗马尼RomaniBarony struct {
	feud.BaseBarony
}

var BRomani罗马尼 feud.Barony = &罗马尼RomaniBarony{}

func init() {
    f := BRomani罗马尼.(*罗马尼RomaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "romani",
		TitleName: "罗马尼",
		TitleCode: "b_romani",
	}
}
