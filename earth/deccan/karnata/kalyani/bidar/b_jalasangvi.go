package bidar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇罗僧毗JalasangviBarony struct {
	feud.BaseBarony
}

var BJalasangvi阇罗僧毗 feud.Barony = &阇罗僧毗JalasangviBarony{}

func init() {
	f := BJalasangvi阇罗僧毗.(*阇罗僧毗JalasangviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jalasangvi",
		TitleName: "阇罗僧毗",
		TitleCode: "b_jalasangvi",
	}
}
