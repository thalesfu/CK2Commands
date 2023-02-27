package nurnberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔姆巴赫KulmbachBarony struct {
	feud.BaseBarony
}

var BKulmbach库尔姆巴赫 feud.Barony = &库尔姆巴赫KulmbachBarony{}

func init() {
    f := BKulmbach库尔姆巴赫.(*库尔姆巴赫KulmbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kulmbach",
		TitleName: "库尔姆巴赫",
		TitleCode: "b_kulmbach",
	}
}
