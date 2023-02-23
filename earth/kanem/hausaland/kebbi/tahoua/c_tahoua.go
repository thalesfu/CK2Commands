package tahoua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TahouaCounty interface {
	feud.County
	BGurumi古鲁米() feud.Barony
	BKuanawa夸纳瓦() feud.Barony
	BOkuaro奥夸罗() feud.Barony
	BRurum鲁鲁姆() feud.Barony
	BTahoua塔瓦() feud.Barony
}

type 塔瓦TahouaCounty struct {
	feud.BaseCounty
	古鲁米Gurumi  feud.Barony
	夸纳瓦Kuanawa feud.Barony
	奥夸罗Okuaro  feud.Barony
	鲁鲁姆Rurum   feud.Barony
	塔瓦Tahoua   feud.Barony
}

func (c *塔瓦TahouaCounty) BGurumi古鲁米() feud.Barony {
	return c.古鲁米Gurumi
}

func (c *塔瓦TahouaCounty) BKuanawa夸纳瓦() feud.Barony {
	return c.夸纳瓦Kuanawa
}

func (c *塔瓦TahouaCounty) BOkuaro奥夸罗() feud.Barony {
	return c.奥夸罗Okuaro
}

func (c *塔瓦TahouaCounty) BRurum鲁鲁姆() feud.Barony {
	return c.鲁鲁姆Rurum
}

func (c *塔瓦TahouaCounty) BTahoua塔瓦() feud.Barony {
	return c.塔瓦Tahoua
}

var CTahoua塔瓦 TahouaCounty = &塔瓦TahouaCounty{}

func init() {
	f := CTahoua塔瓦.(*塔瓦TahouaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1749",
		Title:     "tahoua",
		TitleName: "塔瓦",
		TitleCode: "c_tahoua",
		Baronies:  map[string]feud.Barony{},
	}

	f.古鲁米Gurumi = BGurumi古鲁米
	f.古鲁米Gurumi.SetParent(f)

	f.夸纳瓦Kuanawa = BKuanawa夸纳瓦
	f.夸纳瓦Kuanawa.SetParent(f)

	f.奥夸罗Okuaro = BOkuaro奥夸罗
	f.奥夸罗Okuaro.SetParent(f)

	f.鲁鲁姆Rurum = BRurum鲁鲁姆
	f.鲁鲁姆Rurum.SetParent(f)

	f.塔瓦Tahoua = BTahoua塔瓦
	f.塔瓦Tahoua.SetParent(f)

}
