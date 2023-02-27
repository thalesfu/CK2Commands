package kanara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莽葛奴儿MangalurBarony struct {
	feud.BaseBarony
}

var BMangalur莽葛奴儿 feud.Barony = &莽葛奴儿MangalurBarony{}

func init() {
    f := BMangalur莽葛奴儿.(*莽葛奴儿MangalurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mangalur",
		TitleName: "莽葛奴儿",
		TitleCode: "b_mangalur",
	}
}
