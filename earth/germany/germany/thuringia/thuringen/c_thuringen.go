package thuringen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThuringenCounty interface {
	feud.County
	BArnstadt阿恩施塔特() feud.Barony
	BEichsfeld艾希斯费尔德() feud.Barony
	BErfurt埃尔福特() feud.Barony
	BHenneberg亨讷贝格() feud.Barony
	BMuhlhausen米尔豪森() feud.Barony
	BReuss罗伊斯() feud.Barony
	BSalzungen萨尔聪根() feud.Barony
	BSchmalkalden施马尔卡尔登() feud.Barony
}

type 图林根ThuringenCounty struct {
	feud.BaseCounty
	阿恩施塔特Arnstadt      feud.Barony
	艾希斯费尔德Eichsfeld    feud.Barony
	埃尔福特Erfurt         feud.Barony
	亨讷贝格Henneberg      feud.Barony
	米尔豪森Muhlhausen     feud.Barony
	罗伊斯Reuss           feud.Barony
	萨尔聪根Salzungen      feud.Barony
	施马尔卡尔登Schmalkalden feud.Barony
}

func (c *图林根ThuringenCounty) BArnstadt阿恩施塔特() feud.Barony {
	return c.阿恩施塔特Arnstadt
}

func (c *图林根ThuringenCounty) BEichsfeld艾希斯费尔德() feud.Barony {
	return c.艾希斯费尔德Eichsfeld
}

func (c *图林根ThuringenCounty) BErfurt埃尔福特() feud.Barony {
	return c.埃尔福特Erfurt
}

func (c *图林根ThuringenCounty) BHenneberg亨讷贝格() feud.Barony {
	return c.亨讷贝格Henneberg
}

func (c *图林根ThuringenCounty) BMuhlhausen米尔豪森() feud.Barony {
	return c.米尔豪森Muhlhausen
}

func (c *图林根ThuringenCounty) BReuss罗伊斯() feud.Barony {
	return c.罗伊斯Reuss
}

func (c *图林根ThuringenCounty) BSalzungen萨尔聪根() feud.Barony {
	return c.萨尔聪根Salzungen
}

func (c *图林根ThuringenCounty) BSchmalkalden施马尔卡尔登() feud.Barony {
	return c.施马尔卡尔登Schmalkalden
}

var CThuringen图林根 ThuringenCounty = &图林根ThuringenCounty{}

func init() {
	f := CThuringen图林根.(*图林根ThuringenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "255",
		Title:     "thuringen",
		TitleName: "图林根",
		TitleCode: "c_thuringen",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿恩施塔特Arnstadt = BArnstadt阿恩施塔特
	f.阿恩施塔特Arnstadt.SetParent(f)

	f.艾希斯费尔德Eichsfeld = BEichsfeld艾希斯费尔德
	f.艾希斯费尔德Eichsfeld.SetParent(f)

	f.埃尔福特Erfurt = BErfurt埃尔福特
	f.埃尔福特Erfurt.SetParent(f)

	f.亨讷贝格Henneberg = BHenneberg亨讷贝格
	f.亨讷贝格Henneberg.SetParent(f)

	f.米尔豪森Muhlhausen = BMuhlhausen米尔豪森
	f.米尔豪森Muhlhausen.SetParent(f)

	f.罗伊斯Reuss = BReuss罗伊斯
	f.罗伊斯Reuss.SetParent(f)

	f.萨尔聪根Salzungen = BSalzungen萨尔聪根
	f.萨尔聪根Salzungen.SetParent(f)

	f.施马尔卡尔登Schmalkalden = BSchmalkalden施马尔卡尔登
	f.施马尔卡尔登Schmalkalden.SetParent(f)

}
