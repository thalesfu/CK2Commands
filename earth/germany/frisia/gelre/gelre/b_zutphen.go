package gelre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 聚特芬ZutphenBarony struct {
	feud.BaseBarony
}

var BZutphen聚特芬 feud.Barony = &聚特芬ZutphenBarony{}

func init() {
    f := BZutphen聚特芬.(*聚特芬ZutphenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zutphen",
		TitleName: "聚特芬",
		TitleCode: "b_zutphen",
	}
}
