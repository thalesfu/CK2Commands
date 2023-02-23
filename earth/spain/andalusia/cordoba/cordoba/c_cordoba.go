package cordoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CordobaCounty interface {
	feud.County
	BAlcolea阿尔科莱阿() feud.Barony
	BBelalcazar贝拉尔卡萨() feud.Barony
	BCabra卡夫拉() feud.Barony
	BCanetedelastorres卡涅特德拉斯托雷斯() feud.Barony
	BCordoba科尔多瓦() feud.Barony
	BLucena卢塞纳() feud.Barony
	BMartos马尔托斯() feud.Barony
}

type 科尔多瓦CordobaCounty struct {
	feud.BaseCounty
	阿尔科莱阿Alcolea               feud.Barony
	贝拉尔卡萨Belalcazar            feud.Barony
	卡夫拉Cabra                   feud.Barony
	卡涅特德拉斯托雷斯Canetedelastorres feud.Barony
	科尔多瓦Cordoba                feud.Barony
	卢塞纳Lucena                  feud.Barony
	马尔托斯Martos                 feud.Barony
}

func (c *科尔多瓦CordobaCounty) BAlcolea阿尔科莱阿() feud.Barony {
	return c.阿尔科莱阿Alcolea
}

func (c *科尔多瓦CordobaCounty) BBelalcazar贝拉尔卡萨() feud.Barony {
	return c.贝拉尔卡萨Belalcazar
}

func (c *科尔多瓦CordobaCounty) BCabra卡夫拉() feud.Barony {
	return c.卡夫拉Cabra
}

func (c *科尔多瓦CordobaCounty) BCanetedelastorres卡涅特德拉斯托雷斯() feud.Barony {
	return c.卡涅特德拉斯托雷斯Canetedelastorres
}

func (c *科尔多瓦CordobaCounty) BCordoba科尔多瓦() feud.Barony {
	return c.科尔多瓦Cordoba
}

func (c *科尔多瓦CordobaCounty) BLucena卢塞纳() feud.Barony {
	return c.卢塞纳Lucena
}

func (c *科尔多瓦CordobaCounty) BMartos马尔托斯() feud.Barony {
	return c.马尔托斯Martos
}

var CCordoba科尔多瓦 CordobaCounty = &科尔多瓦CordobaCounty{}

func init() {
	f := CCordoba科尔多瓦.(*科尔多瓦CordobaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "181",
		Title:     "cordoba",
		TitleName: "科尔多瓦",
		TitleCode: "c_cordoba",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔科莱阿Alcolea = BAlcolea阿尔科莱阿
	f.阿尔科莱阿Alcolea.SetParent(f)

	f.贝拉尔卡萨Belalcazar = BBelalcazar贝拉尔卡萨
	f.贝拉尔卡萨Belalcazar.SetParent(f)

	f.卡夫拉Cabra = BCabra卡夫拉
	f.卡夫拉Cabra.SetParent(f)

	f.卡涅特德拉斯托雷斯Canetedelastorres = BCanetedelastorres卡涅特德拉斯托雷斯
	f.卡涅特德拉斯托雷斯Canetedelastorres.SetParent(f)

	f.科尔多瓦Cordoba = BCordoba科尔多瓦
	f.科尔多瓦Cordoba.SetParent(f)

	f.卢塞纳Lucena = BLucena卢塞纳
	f.卢塞纳Lucena.SetParent(f)

	f.马尔托斯Martos = BMartos马尔托斯
	f.马尔托斯Martos.SetParent(f)

}
