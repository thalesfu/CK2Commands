package hull

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 怀克WykeBarony struct {
	feud.BaseBarony
}

var BWyke怀克 feud.Barony = &怀克WykeBarony{}

func init() {
	f := BWyke怀克.(*怀克WykeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wyke",
		TitleName: "怀克",
		TitleCode: "b_wyke",
	}
}
