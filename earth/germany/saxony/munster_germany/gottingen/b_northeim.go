package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺特海姆NortheimBarony struct {
	feud.BaseBarony
}

var BNortheim诺特海姆 feud.Barony = &诺特海姆NortheimBarony{}

func init() {
    f := BNortheim诺特海姆.(*诺特海姆NortheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "northeim",
		TitleName: "诺特海姆",
		TitleCode: "b_northeim",
	}
}
