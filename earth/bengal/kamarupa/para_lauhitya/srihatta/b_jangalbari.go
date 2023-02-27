package srihatta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旬迦罗婆利JangalbariBarony struct {
	feud.BaseBarony
}

var BJangalbari旬迦罗婆利 feud.Barony = &旬迦罗婆利JangalbariBarony{}

func init() {
    f := BJangalbari旬迦罗婆利.(*旬迦罗婆利JangalbariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jangalbari",
		TitleName: "旬迦罗婆利",
		TitleCode: "b_jangalbari",
	}
}
