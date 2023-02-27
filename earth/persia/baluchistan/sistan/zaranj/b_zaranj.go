package zaranj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 疾陵ZaranjBarony struct {
	feud.BaseBarony
}

var BZaranj疾陵 feud.Barony = &疾陵ZaranjBarony{}

func init() {
    f := BZaranj疾陵.(*疾陵ZaranjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaranj",
		TitleName: "疾陵",
		TitleCode: "b_zaranj",
	}
}
