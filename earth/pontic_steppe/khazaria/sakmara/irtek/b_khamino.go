package irtek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈米诺KhaminoBarony struct {
	feud.BaseBarony
}

var BKhamino哈米诺 feud.Barony = &哈米诺KhaminoBarony{}

func init() {
    f := BKhamino哈米诺.(*哈米诺KhaminoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khamino",
		TitleName: "哈米诺",
		TitleCode: "b_khamino",
	}
}
