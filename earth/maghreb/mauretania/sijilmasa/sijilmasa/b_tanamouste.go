package sijilmasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔纳穆斯特TanamousteBarony struct {
	feud.BaseBarony
}

var BTanamouste塔纳穆斯特 feud.Barony = &塔纳穆斯特TanamousteBarony{}

func init() {
    f := BTanamouste塔纳穆斯特.(*塔纳穆斯特TanamousteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tanamouste",
		TitleName: "塔纳穆斯特",
		TitleCode: "b_tanamouste",
	}
}
