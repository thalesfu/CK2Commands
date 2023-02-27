package pangong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克什尔村KizilyezaBarony struct {
	feud.BaseBarony
}

var BKizilyeza克什尔村 feud.Barony = &克什尔村KizilyezaBarony{}

func init() {
    f := BKizilyeza克什尔村.(*克什尔村KizilyezaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kizilyeza",
		TitleName: "克什尔村",
		TitleCode: "b_kizilyeza",
	}
}
