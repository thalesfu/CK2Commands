package benevento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷维科TrevicoBarony struct {
	feud.BaseBarony
}

var BTrevico特雷维科 feud.Barony = &特雷维科TrevicoBarony{}

func init() {
	f := BTrevico特雷维科.(*特雷维科TrevicoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trevico",
		TitleName: "特雷维科",
		TitleCode: "b_trevico",
	}
}
