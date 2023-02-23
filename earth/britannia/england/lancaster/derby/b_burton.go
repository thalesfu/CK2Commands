package derby

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯顿BurtonBarony struct {
	feud.BaseBarony
}

var BBurton伯顿 feud.Barony = &伯顿BurtonBarony{}

func init() {
	f := BBurton伯顿.(*伯顿BurtonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burton",
		TitleName: "伯顿",
		TitleCode: "b_burton",
	}
}
