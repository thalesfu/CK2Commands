package belukha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琼古尔TyungurBarony struct {
	feud.BaseBarony
}

var BTyungur琼古尔 feud.Barony = &琼古尔TyungurBarony{}

func init() {
    f := BTyungur琼古尔.(*琼古尔TyungurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyungur",
		TitleName: "琼古尔",
		TitleCode: "b_tyungur",
	}
}
