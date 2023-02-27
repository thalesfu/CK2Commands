package osnabruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本特海姆BentheimBarony struct {
	feud.BaseBarony
}

var BBentheim本特海姆 feud.Barony = &本特海姆BentheimBarony{}

func init() {
    f := BBentheim本特海姆.(*本特海姆BentheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bentheim",
		TitleName: "本特海姆",
		TitleCode: "b_bentheim",
	}
}
