package ostfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔库姆BorkumBarony struct {
	feud.BaseBarony
}

var BBorkum博尔库姆 feud.Barony = &博尔库姆BorkumBarony{}

func init() {
    f := BBorkum博尔库姆.(*博尔库姆BorkumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borkum",
		TitleName: "博尔库姆",
		TitleCode: "b_borkum",
	}
}
