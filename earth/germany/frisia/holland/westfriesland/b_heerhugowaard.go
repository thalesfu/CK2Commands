package westfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海尔许霍瓦德HeerhugowaardBarony struct {
	feud.BaseBarony
}

var BHeerhugowaard海尔许霍瓦德 feud.Barony = &海尔许霍瓦德HeerhugowaardBarony{}

func init() {
    f := BHeerhugowaard海尔许霍瓦德.(*海尔许霍瓦德HeerhugowaardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heerhugowaard",
		TitleName: "海尔许霍瓦德",
		TitleCode: "b_heerhugowaard",
	}
}
