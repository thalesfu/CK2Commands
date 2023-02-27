package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马赫迪耶MahdiaBarony struct {
	feud.BaseBarony
}

var BMahdia马赫迪耶 feud.Barony = &马赫迪耶MahdiaBarony{}

func init() {
    f := BMahdia马赫迪耶.(*马赫迪耶MahdiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahdia",
		TitleName: "马赫迪耶",
		TitleCode: "b_mahdia",
	}
}
