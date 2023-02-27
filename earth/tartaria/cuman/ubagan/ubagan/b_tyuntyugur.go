package ubagan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 琼秋古尔TyuntyugurBarony struct {
	feud.BaseBarony
}

var BTyuntyugur琼秋古尔 feud.Barony = &琼秋古尔TyuntyugurBarony{}

func init() {
    f := BTyuntyugur琼秋古尔.(*琼秋古尔TyuntyugurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyuntyugur",
		TitleName: "琼秋古尔",
		TitleCode: "b_tyuntyugur",
	}
}
