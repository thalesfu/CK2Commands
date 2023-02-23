package beloozero

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BeloozeroCounty interface {
	feud.County
	BBaryevychi巴里耶维奇() feud.Barony
	BBelozersk别洛奥泽罗() feud.Barony
	BCherepovets切列波韦茨() feud.Barony
	BFedosyevo费多谢沃() feud.Barony
	BFerapontov费拉蓬托夫() feud.Barony
	BKinilov基尼洛夫() feud.Barony
	BKirillobelozersky基里尔别洛泽尔斯基() feud.Barony
}

type 别洛奥泽罗BeloozeroCounty struct {
	feud.BaseCounty
	巴里耶维奇Baryevychi            feud.Barony
	别洛奥泽罗Belozersk             feud.Barony
	切列波韦茨Cherepovets           feud.Barony
	费多谢沃Fedosyevo              feud.Barony
	费拉蓬托夫Ferapontov            feud.Barony
	基尼洛夫Kinilov                feud.Barony
	基里尔别洛泽尔斯基Kirillobelozersky feud.Barony
}

func (c *别洛奥泽罗BeloozeroCounty) BBaryevychi巴里耶维奇() feud.Barony {
	return c.巴里耶维奇Baryevychi
}

func (c *别洛奥泽罗BeloozeroCounty) BBelozersk别洛奥泽罗() feud.Barony {
	return c.别洛奥泽罗Belozersk
}

func (c *别洛奥泽罗BeloozeroCounty) BCherepovets切列波韦茨() feud.Barony {
	return c.切列波韦茨Cherepovets
}

func (c *别洛奥泽罗BeloozeroCounty) BFedosyevo费多谢沃() feud.Barony {
	return c.费多谢沃Fedosyevo
}

func (c *别洛奥泽罗BeloozeroCounty) BFerapontov费拉蓬托夫() feud.Barony {
	return c.费拉蓬托夫Ferapontov
}

func (c *别洛奥泽罗BeloozeroCounty) BKinilov基尼洛夫() feud.Barony {
	return c.基尼洛夫Kinilov
}

func (c *别洛奥泽罗BeloozeroCounty) BKirillobelozersky基里尔别洛泽尔斯基() feud.Barony {
	return c.基里尔别洛泽尔斯基Kirillobelozersky
}

var CBeloozero别洛奥泽罗 BeloozeroCounty = &别洛奥泽罗BeloozeroCounty{}

func init() {
	f := CBeloozero别洛奥泽罗.(*别洛奥泽罗BeloozeroCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "408",
		Title:     "beloozero",
		TitleName: "别洛奥泽罗",
		TitleCode: "c_beloozero",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴里耶维奇Baryevychi = BBaryevychi巴里耶维奇
	f.巴里耶维奇Baryevychi.SetParent(f)

	f.别洛奥泽罗Belozersk = BBelozersk别洛奥泽罗
	f.别洛奥泽罗Belozersk.SetParent(f)

	f.切列波韦茨Cherepovets = BCherepovets切列波韦茨
	f.切列波韦茨Cherepovets.SetParent(f)

	f.费多谢沃Fedosyevo = BFedosyevo费多谢沃
	f.费多谢沃Fedosyevo.SetParent(f)

	f.费拉蓬托夫Ferapontov = BFerapontov费拉蓬托夫
	f.费拉蓬托夫Ferapontov.SetParent(f)

	f.基尼洛夫Kinilov = BKinilov基尼洛夫
	f.基尼洛夫Kinilov.SetParent(f)

	f.基里尔别洛泽尔斯基Kirillobelozersky = BKirillobelozersky基里尔别洛泽尔斯基
	f.基里尔别洛泽尔斯基Kirillobelozersky.SetParent(f)

}
