package paro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涉鲁囊TshalunangBarony struct {
	feud.BaseBarony
}

var BTshalunang涉鲁囊 feud.Barony = &涉鲁囊TshalunangBarony{}

func init() {
	f := BTshalunang涉鲁囊.(*涉鲁囊TshalunangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tshalunang",
		TitleName: "涉鲁囊",
		TitleCode: "b_tshalunang",
	}
}
