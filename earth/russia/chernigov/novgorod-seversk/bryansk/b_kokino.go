package bryansk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科基诺KokinoBarony struct {
	feud.BaseBarony
}

var BKokino科基诺 feud.Barony = &科基诺KokinoBarony{}

func init() {
    f := BKokino科基诺.(*科基诺KokinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kokino",
		TitleName: "科基诺",
		TitleCode: "b_kokino",
	}
}
