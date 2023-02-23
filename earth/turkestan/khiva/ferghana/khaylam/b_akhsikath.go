package khaylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西鞬AkhsikathBarony struct {
	feud.BaseBarony
}

var BAkhsikath西鞬 feud.Barony = &西鞬AkhsikathBarony{}

func init() {
	f := BAkhsikath西鞬.(*西鞬AkhsikathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhsikath",
		TitleName: "西鞬",
		TitleCode: "b_akhsikath",
	}
}
