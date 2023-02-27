package cinarca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯特利卡西亚BastelicacciaBarony struct {
	feud.BaseBarony
}

var BBastelicaccia巴斯特利卡西亚 feud.Barony = &巴斯特利卡西亚BastelicacciaBarony{}

func init() {
    f := BBastelicaccia巴斯特利卡西亚.(*巴斯特利卡西亚BastelicacciaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bastelicaccia",
		TitleName: "巴斯特利卡西亚",
		TitleCode: "b_bastelicaccia",
	}
}
