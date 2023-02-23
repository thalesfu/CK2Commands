package calatrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CalatravaCounty interface {
	feud.County
	BAlarcos阿拉尔科斯() feud.Barony
	BAlcazardesanjuan圣胡安堡() feud.Barony
	BAlmadeo阿尔马登() feud.Barony
	BAlmodovardelcampo阿尔莫多瓦尔德尔坎波() feud.Barony
	BCalatrava卡拉特拉瓦() feud.Barony
	BCaracuel卡拉库埃尔() feud.Barony
	BMedellin麦德林() feud.Barony
	BVillareal比利亚雷亚尔() feud.Barony
}

type 卡拉特拉瓦CalatravaCounty struct {
	feud.BaseCounty
	阿拉尔科斯Alarcos                feud.Barony
	圣胡安堡Alcazardesanjuan        feud.Barony
	阿尔马登Almadeo                 feud.Barony
	阿尔莫多瓦尔德尔坎波Almodovardelcampo feud.Barony
	卡拉特拉瓦Calatrava              feud.Barony
	卡拉库埃尔Caracuel               feud.Barony
	麦德林Medellin                 feud.Barony
	比利亚雷亚尔Villareal             feud.Barony
}

func (c *卡拉特拉瓦CalatravaCounty) BAlarcos阿拉尔科斯() feud.Barony {
	return c.阿拉尔科斯Alarcos
}

func (c *卡拉特拉瓦CalatravaCounty) BAlcazardesanjuan圣胡安堡() feud.Barony {
	return c.圣胡安堡Alcazardesanjuan
}

func (c *卡拉特拉瓦CalatravaCounty) BAlmadeo阿尔马登() feud.Barony {
	return c.阿尔马登Almadeo
}

func (c *卡拉特拉瓦CalatravaCounty) BAlmodovardelcampo阿尔莫多瓦尔德尔坎波() feud.Barony {
	return c.阿尔莫多瓦尔德尔坎波Almodovardelcampo
}

func (c *卡拉特拉瓦CalatravaCounty) BCalatrava卡拉特拉瓦() feud.Barony {
	return c.卡拉特拉瓦Calatrava
}

func (c *卡拉特拉瓦CalatravaCounty) BCaracuel卡拉库埃尔() feud.Barony {
	return c.卡拉库埃尔Caracuel
}

func (c *卡拉特拉瓦CalatravaCounty) BMedellin麦德林() feud.Barony {
	return c.麦德林Medellin
}

func (c *卡拉特拉瓦CalatravaCounty) BVillareal比利亚雷亚尔() feud.Barony {
	return c.比利亚雷亚尔Villareal
}

var CCalatrava卡拉特拉瓦 CalatravaCounty = &卡拉特拉瓦CalatravaCounty{}

func init() {
	f := CCalatrava卡拉特拉瓦.(*卡拉特拉瓦CalatravaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "196",
		Title:     "calatrava",
		TitleName: "卡拉特拉瓦",
		TitleCode: "c_calatrava",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉尔科斯Alarcos = BAlarcos阿拉尔科斯
	f.阿拉尔科斯Alarcos.SetParent(f)

	f.圣胡安堡Alcazardesanjuan = BAlcazardesanjuan圣胡安堡
	f.圣胡安堡Alcazardesanjuan.SetParent(f)

	f.阿尔马登Almadeo = BAlmadeo阿尔马登
	f.阿尔马登Almadeo.SetParent(f)

	f.阿尔莫多瓦尔德尔坎波Almodovardelcampo = BAlmodovardelcampo阿尔莫多瓦尔德尔坎波
	f.阿尔莫多瓦尔德尔坎波Almodovardelcampo.SetParent(f)

	f.卡拉特拉瓦Calatrava = BCalatrava卡拉特拉瓦
	f.卡拉特拉瓦Calatrava.SetParent(f)

	f.卡拉库埃尔Caracuel = BCaracuel卡拉库埃尔
	f.卡拉库埃尔Caracuel.SetParent(f)

	f.麦德林Medellin = BMedellin麦德林
	f.麦德林Medellin.SetParent(f)

	f.比利亚雷亚尔Villareal = BVillareal比利亚雷亚尔
	f.比利亚雷亚尔Villareal.SetParent(f)

}
