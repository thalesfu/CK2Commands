package rafha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜韦德DuwaydBarony struct {
	feud.BaseBarony
}

var BDuwayd杜韦德 feud.Barony = &杜韦德DuwaydBarony{}

func init() {
	f := BDuwayd杜韦德.(*杜韦德DuwaydBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duwayd",
		TitleName: "杜韦德",
		TitleCode: "b_duwayd",
	}
}
