package lower_silesia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥得河畔比托姆BytomodrzBarony struct {
	feud.BaseBarony
}

var BBytomodrz奥得河畔比托姆 feud.Barony = &奥得河畔比托姆BytomodrzBarony{}

func init() {
    f := BBytomodrz奥得河畔比托姆.(*奥得河畔比托姆BytomodrzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bytomodrz",
		TitleName: "奥得河畔比托姆",
		TitleCode: "b_bytomodrz",
	}
}
