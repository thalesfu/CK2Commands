package dijon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西托CiteauxBarony struct {
	feud.BaseBarony
}

var BCiteaux西托 feud.Barony = &西托CiteauxBarony{}

func init() {
	f := BCiteaux西托.(*西托CiteauxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "citeaux",
		TitleName: "西托",
		TitleCode: "b_citeaux",
	}
}
