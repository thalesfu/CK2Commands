package piemonte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷瓦科雷CrevacuoreBarony struct {
	feud.BaseBarony
}

var BCrevacuore克雷瓦科雷 feud.Barony = &克雷瓦科雷CrevacuoreBarony{}

func init() {
	f := BCrevacuore克雷瓦科雷.(*克雷瓦科雷CrevacuoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crevacuore",
		TitleName: "克雷瓦科雷",
		TitleCode: "b_crevacuore",
	}
}
