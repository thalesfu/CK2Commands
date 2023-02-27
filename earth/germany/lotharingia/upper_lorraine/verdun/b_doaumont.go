package verdun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜奥蒙DoaumontBarony struct {
	feud.BaseBarony
}

var BDoaumont杜奥蒙 feud.Barony = &杜奥蒙DoaumontBarony{}

func init() {
    f := BDoaumont杜奥蒙.(*杜奥蒙DoaumontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doaumont",
		TitleName: "杜奥蒙",
		TitleCode: "b_doaumont",
	}
}
