package beloozero

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基里尔别洛泽尔斯基KirillobelozerskyBarony struct {
	feud.BaseBarony
}

var BKirillobelozersky基里尔别洛泽尔斯基 feud.Barony = &基里尔别洛泽尔斯基KirillobelozerskyBarony{}

func init() {
    f := BKirillobelozersky基里尔别洛泽尔斯基.(*基里尔别洛泽尔斯基KirillobelozerskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirillobelozersky",
		TitleName: "基里尔别洛泽尔斯基",
		TitleCode: "b_kirillobelozersky",
	}
}
