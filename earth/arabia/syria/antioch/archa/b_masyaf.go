package archa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈斯亚夫MasyafBarony struct {
	feud.BaseBarony
}

var BMasyaf迈斯亚夫 feud.Barony = &迈斯亚夫MasyafBarony{}

func init() {
	f := BMasyaf迈斯亚夫.(*迈斯亚夫MasyafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masyaf",
		TitleName: "迈斯亚夫",
		TitleCode: "b_masyaf",
	}
}
