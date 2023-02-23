package archa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔卡ArchaBarony struct {
	feud.BaseBarony
}

var BArcha阿尔卡 feud.Barony = &阿尔卡ArchaBarony{}

func init() {
	f := BArcha阿尔卡.(*阿尔卡ArchaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "archa",
		TitleName: "阿尔卡",
		TitleCode: "b_archa",
	}
}
