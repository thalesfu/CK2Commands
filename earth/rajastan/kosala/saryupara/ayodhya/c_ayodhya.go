package ayodhya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AyodhyaCounty interface {
	feud.County
	BAyodhya阿踰陀() feud.Barony
	BFaizabad法扎巴德() feud.Barony
	BNageshwarnath那祇湿伐罗那他() feud.Barony
	BRamkot罗摩拘吒() feud.Barony
	BSarairasi萨罗依罗斯() feud.Barony
}

type 阿踰陀AyodhyaCounty struct {
	feud.BaseCounty
	阿踰陀Ayodhya           feud.Barony
	法扎巴德Faizabad         feud.Barony
	那祇湿伐罗那他Nageshwarnath feud.Barony
	罗摩拘吒Ramkot           feud.Barony
	萨罗依罗斯Sarairasi       feud.Barony
}

func (c *阿踰陀AyodhyaCounty) BAyodhya阿踰陀() feud.Barony {
	return c.阿踰陀Ayodhya
}

func (c *阿踰陀AyodhyaCounty) BFaizabad法扎巴德() feud.Barony {
	return c.法扎巴德Faizabad
}

func (c *阿踰陀AyodhyaCounty) BNageshwarnath那祇湿伐罗那他() feud.Barony {
	return c.那祇湿伐罗那他Nageshwarnath
}

func (c *阿踰陀AyodhyaCounty) BRamkot罗摩拘吒() feud.Barony {
	return c.罗摩拘吒Ramkot
}

func (c *阿踰陀AyodhyaCounty) BSarairasi萨罗依罗斯() feud.Barony {
	return c.萨罗依罗斯Sarairasi
}

var CAyodhya阿踰陀 AyodhyaCounty = &阿踰陀AyodhyaCounty{}

func init() {
	f := CAyodhya阿踰陀.(*阿踰陀AyodhyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1250",
		Title:     "ayodhya",
		TitleName: "阿踰陀",
		TitleCode: "c_ayodhya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿踰陀Ayodhya = BAyodhya阿踰陀
	f.阿踰陀Ayodhya.SetParent(f)

	f.法扎巴德Faizabad = BFaizabad法扎巴德
	f.法扎巴德Faizabad.SetParent(f)

	f.那祇湿伐罗那他Nageshwarnath = BNageshwarnath那祇湿伐罗那他
	f.那祇湿伐罗那他Nageshwarnath.SetParent(f)

	f.罗摩拘吒Ramkot = BRamkot罗摩拘吒
	f.罗摩拘吒Ramkot.SetParent(f)

	f.萨罗依罗斯Sarairasi = BSarairasi萨罗依罗斯
	f.萨罗依罗斯Sarairasi.SetParent(f)

}
