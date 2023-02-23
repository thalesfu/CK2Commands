package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SalzburgCounty interface {
	feud.County
	BBerchtesgaden贝希特斯加登() feud.Barony
	BDurmberg迪姆贝格() feud.Barony
	BGastein加斯泰因() feud.Barony
	BLaufen劳芬() feud.Barony
	BMuhldorf米尔多夫() feud.Barony
	BPlain普莱恩() feud.Barony
	BSalzburg萨尔茨堡() feud.Barony
	BTittmoning蒂特莫宁() feud.Barony
	BWaging瓦京() feud.Barony
}

type 萨尔茨堡SalzburgCounty struct {
	feud.BaseCounty
	贝希特斯加登Berchtesgaden feud.Barony
	迪姆贝格Durmberg        feud.Barony
	加斯泰因Gastein         feud.Barony
	劳芬Laufen            feud.Barony
	米尔多夫Muhldorf        feud.Barony
	普莱恩Plain            feud.Barony
	萨尔茨堡Salzburg        feud.Barony
	蒂特莫宁Tittmoning      feud.Barony
	瓦京Waging            feud.Barony
}

func (c *萨尔茨堡SalzburgCounty) BBerchtesgaden贝希特斯加登() feud.Barony {
	return c.贝希特斯加登Berchtesgaden
}

func (c *萨尔茨堡SalzburgCounty) BDurmberg迪姆贝格() feud.Barony {
	return c.迪姆贝格Durmberg
}

func (c *萨尔茨堡SalzburgCounty) BGastein加斯泰因() feud.Barony {
	return c.加斯泰因Gastein
}

func (c *萨尔茨堡SalzburgCounty) BLaufen劳芬() feud.Barony {
	return c.劳芬Laufen
}

func (c *萨尔茨堡SalzburgCounty) BMuhldorf米尔多夫() feud.Barony {
	return c.米尔多夫Muhldorf
}

func (c *萨尔茨堡SalzburgCounty) BPlain普莱恩() feud.Barony {
	return c.普莱恩Plain
}

func (c *萨尔茨堡SalzburgCounty) BSalzburg萨尔茨堡() feud.Barony {
	return c.萨尔茨堡Salzburg
}

func (c *萨尔茨堡SalzburgCounty) BTittmoning蒂特莫宁() feud.Barony {
	return c.蒂特莫宁Tittmoning
}

func (c *萨尔茨堡SalzburgCounty) BWaging瓦京() feud.Barony {
	return c.瓦京Waging
}

var CSalzburg萨尔茨堡 SalzburgCounty = &萨尔茨堡SalzburgCounty{}

func init() {
	f := CSalzburg萨尔茨堡.(*萨尔茨堡SalzburgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "448",
		Title:     "salzburg",
		TitleName: "萨尔茨堡",
		TitleCode: "c_salzburg",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝希特斯加登Berchtesgaden = BBerchtesgaden贝希特斯加登
	f.贝希特斯加登Berchtesgaden.SetParent(f)

	f.迪姆贝格Durmberg = BDurmberg迪姆贝格
	f.迪姆贝格Durmberg.SetParent(f)

	f.加斯泰因Gastein = BGastein加斯泰因
	f.加斯泰因Gastein.SetParent(f)

	f.劳芬Laufen = BLaufen劳芬
	f.劳芬Laufen.SetParent(f)

	f.米尔多夫Muhldorf = BMuhldorf米尔多夫
	f.米尔多夫Muhldorf.SetParent(f)

	f.普莱恩Plain = BPlain普莱恩
	f.普莱恩Plain.SetParent(f)

	f.萨尔茨堡Salzburg = BSalzburg萨尔茨堡
	f.萨尔茨堡Salzburg.SetParent(f)

	f.蒂特莫宁Tittmoning = BTittmoning蒂特莫宁
	f.蒂特莫宁Tittmoning.SetParent(f)

	f.瓦京Waging = BWaging瓦京
	f.瓦京Waging.SetParent(f)

}
