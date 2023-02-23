package karelen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯托穆克沙KostomukshaBarony struct {
	feud.BaseBarony
}

var BKostomuksha科斯托穆克沙 feud.Barony = &科斯托穆克沙KostomukshaBarony{}

func init() {
	f := BKostomuksha科斯托穆克沙.(*科斯托穆克沙KostomukshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kostomuksha",
		TitleName: "科斯托穆克沙",
		TitleCode: "b_kostomuksha",
	}
}
