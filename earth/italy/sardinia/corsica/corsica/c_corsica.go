package corsica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CorsicaCounty interface {
	feud.County
	BAlando阿兰多() feud.Barony
	BAleria阿莱里亚() feud.Barony
	BAlgajola阿尔加约拉() feud.Barony
	BBastia巴斯蒂亚() feud.Barony
	BCorte科尔特() feud.Barony
	BLuri卢里() feud.Barony
	BMorosaglia莫罗萨利亚() feud.Barony
}

type 科西嘉CorsicaCounty struct {
	feud.BaseCounty
	阿兰多Alando       feud.Barony
	阿莱里亚Aleria      feud.Barony
	阿尔加约拉Algajola   feud.Barony
	巴斯蒂亚Bastia      feud.Barony
	科尔特Corte        feud.Barony
	卢里Luri          feud.Barony
	莫罗萨利亚Morosaglia feud.Barony
}

func (c *科西嘉CorsicaCounty) BAlando阿兰多() feud.Barony {
	return c.阿兰多Alando
}

func (c *科西嘉CorsicaCounty) BAleria阿莱里亚() feud.Barony {
	return c.阿莱里亚Aleria
}

func (c *科西嘉CorsicaCounty) BAlgajola阿尔加约拉() feud.Barony {
	return c.阿尔加约拉Algajola
}

func (c *科西嘉CorsicaCounty) BBastia巴斯蒂亚() feud.Barony {
	return c.巴斯蒂亚Bastia
}

func (c *科西嘉CorsicaCounty) BCorte科尔特() feud.Barony {
	return c.科尔特Corte
}

func (c *科西嘉CorsicaCounty) BLuri卢里() feud.Barony {
	return c.卢里Luri
}

func (c *科西嘉CorsicaCounty) BMorosaglia莫罗萨利亚() feud.Barony {
	return c.莫罗萨利亚Morosaglia
}

var CCorsica科西嘉 CorsicaCounty = &科西嘉CorsicaCounty{}

func init() {
	f := CCorsica科西嘉.(*科西嘉CorsicaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "324",
		Title:     "corsica",
		TitleName: "科西嘉",
		TitleCode: "c_corsica",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿兰多Alando = BAlando阿兰多
	f.阿兰多Alando.SetParent(f)

	f.阿莱里亚Aleria = BAleria阿莱里亚
	f.阿莱里亚Aleria.SetParent(f)

	f.阿尔加约拉Algajola = BAlgajola阿尔加约拉
	f.阿尔加约拉Algajola.SetParent(f)

	f.巴斯蒂亚Bastia = BBastia巴斯蒂亚
	f.巴斯蒂亚Bastia.SetParent(f)

	f.科尔特Corte = BCorte科尔特
	f.科尔特Corte.SetParent(f)

	f.卢里Luri = BLuri卢里
	f.卢里Luri.SetParent(f)

	f.莫罗萨利亚Morosaglia = BMorosaglia莫罗萨利亚
	f.莫罗萨利亚Morosaglia.SetParent(f)

}
