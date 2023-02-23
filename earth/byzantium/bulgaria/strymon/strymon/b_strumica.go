package strymon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特鲁米察StrumicaBarony struct {
	feud.BaseBarony
}

var BStrumica斯特鲁米察 feud.Barony = &斯特鲁米察StrumicaBarony{}

func init() {
	f := BStrumica斯特鲁米察.(*斯特鲁米察StrumicaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strumica",
		TitleName: "斯特鲁米察",
		TitleCode: "b_strumica",
	}
}
