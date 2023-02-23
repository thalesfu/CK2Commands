package more

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MoreCounty interface {
	feud.County
	BAlem奥勒姆() feud.Barony
	BKalmar卡尔马() feud.Barony
	BMadesjo马德舍() feud.Barony
	BRostocka罗斯托卡() feud.Barony
	BSoderakra索德罗克拉() feud.Barony
	BTorsas图绍斯() feud.Barony
}

type 默雷MoreCounty struct {
	feud.BaseCounty
	奥勒姆Alem        feud.Barony
	卡尔马Kalmar      feud.Barony
	马德舍Madesjo     feud.Barony
	罗斯托卡Rostocka   feud.Barony
	索德罗克拉Soderakra feud.Barony
	图绍斯Torsas      feud.Barony
}

func (c *默雷MoreCounty) BAlem奥勒姆() feud.Barony {
	return c.奥勒姆Alem
}

func (c *默雷MoreCounty) BKalmar卡尔马() feud.Barony {
	return c.卡尔马Kalmar
}

func (c *默雷MoreCounty) BMadesjo马德舍() feud.Barony {
	return c.马德舍Madesjo
}

func (c *默雷MoreCounty) BRostocka罗斯托卡() feud.Barony {
	return c.罗斯托卡Rostocka
}

func (c *默雷MoreCounty) BSoderakra索德罗克拉() feud.Barony {
	return c.索德罗克拉Soderakra
}

func (c *默雷MoreCounty) BTorsas图绍斯() feud.Barony {
	return c.图绍斯Torsas
}

var CMore默雷 MoreCounty = &默雷MoreCounty{}

func init() {
	f := CMore默雷.(*默雷MoreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "931",
		Title:     "more",
		TitleName: "默雷",
		TitleCode: "c_more",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥勒姆Alem = BAlem奥勒姆
	f.奥勒姆Alem.SetParent(f)

	f.卡尔马Kalmar = BKalmar卡尔马
	f.卡尔马Kalmar.SetParent(f)

	f.马德舍Madesjo = BMadesjo马德舍
	f.马德舍Madesjo.SetParent(f)

	f.罗斯托卡Rostocka = BRostocka罗斯托卡
	f.罗斯托卡Rostocka.SetParent(f)

	f.索德罗克拉Soderakra = BSoderakra索德罗克拉
	f.索德罗克拉Soderakra.SetParent(f)

	f.图绍斯Torsas = BTorsas图绍斯
	f.图绍斯Torsas.SetParent(f)

}
