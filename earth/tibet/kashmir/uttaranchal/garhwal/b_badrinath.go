package garhwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆陀梨那他BadrinathBarony struct {
	feud.BaseBarony
}

var BBadrinath婆陀梨那他 feud.Barony = &婆陀梨那他BadrinathBarony{}

func init() {
    f := BBadrinath婆陀梨那他.(*婆陀梨那他BadrinathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "badrinath",
		TitleName: "婆陀梨那他",
		TitleCode: "b_badrinath",
	}
}
