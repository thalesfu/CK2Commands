package osma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜罗DueroBarony struct {
	feud.BaseBarony
}

var BDuero杜罗 feud.Barony = &杜罗DueroBarony{}

func init() {
	f := BDuero杜罗.(*杜罗DueroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duero",
		TitleName: "杜罗",
		TitleCode: "b_duero",
	}
}
