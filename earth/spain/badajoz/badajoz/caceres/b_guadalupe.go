package caceres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓜达卢佩GuadalupeBarony struct {
	feud.BaseBarony
}

var BGuadalupe瓜达卢佩 feud.Barony = &瓜达卢佩GuadalupeBarony{}

func init() {
    f := BGuadalupe瓜达卢佩.(*瓜达卢佩GuadalupeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guadalupe",
		TitleName: "瓜达卢佩",
		TitleCode: "b_guadalupe",
	}
}
