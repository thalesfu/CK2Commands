package wielunska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰尔诺日维CzarnozylyBarony struct {
	feud.BaseBarony
}

var BCzarnozyly恰尔诺日维 feud.Barony = &恰尔诺日维CzarnozylyBarony{}

func init() {
    f := BCzarnozyly恰尔诺日维.(*恰尔诺日维CzarnozylyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "czarnozyly",
		TitleName: "恰尔诺日维",
		TitleCode: "b_czarnozyly",
	}
}
