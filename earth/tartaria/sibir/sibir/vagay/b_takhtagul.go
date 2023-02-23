package vagay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔赫塔古尔TakhtagulBarony struct {
	feud.BaseBarony
}

var BTakhtagul塔赫塔古尔 feud.Barony = &塔赫塔古尔TakhtagulBarony{}

func init() {
	f := BTakhtagul塔赫塔古尔.(*塔赫塔古尔TakhtagulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "takhtagul",
		TitleName: "塔赫塔古尔",
		TitleCode: "b_takhtagul",
	}
}
