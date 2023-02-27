package selenge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳乌什基NaushkiBarony struct {
	feud.BaseBarony
}

var BNaushki纳乌什基 feud.Barony = &纳乌什基NaushkiBarony{}

func init() {
    f := BNaushki纳乌什基.(*纳乌什基NaushkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naushki",
		TitleName: "纳乌什基",
		TitleCode: "b_naushki",
	}
}
