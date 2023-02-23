package karelen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarelenCounty interface {
	feud.County
	BKalevala卡列瓦拉() feud.Barony
	BKem凯姆() feud.Barony
	BKostomuksha科斯托穆克沙() feud.Barony
	BLoukhi洛乌希() feud.Barony
	BMuezersky穆耶泽尔斯基() feud.Barony
	BPitkyaranta皮特基亚兰塔() feud.Barony
	BSordavala索尔塔瓦拉() feud.Barony
}

type 凯姆KarelenCounty struct {
	feud.BaseCounty
	卡列瓦拉Kalevala      feud.Barony
	凯姆Kem             feud.Barony
	科斯托穆克沙Kostomuksha feud.Barony
	洛乌希Loukhi         feud.Barony
	穆耶泽尔斯基Muezersky   feud.Barony
	皮特基亚兰塔Pitkyaranta feud.Barony
	索尔塔瓦拉Sordavala    feud.Barony
}

func (c *凯姆KarelenCounty) BKalevala卡列瓦拉() feud.Barony {
	return c.卡列瓦拉Kalevala
}

func (c *凯姆KarelenCounty) BKem凯姆() feud.Barony {
	return c.凯姆Kem
}

func (c *凯姆KarelenCounty) BKostomuksha科斯托穆克沙() feud.Barony {
	return c.科斯托穆克沙Kostomuksha
}

func (c *凯姆KarelenCounty) BLoukhi洛乌希() feud.Barony {
	return c.洛乌希Loukhi
}

func (c *凯姆KarelenCounty) BMuezersky穆耶泽尔斯基() feud.Barony {
	return c.穆耶泽尔斯基Muezersky
}

func (c *凯姆KarelenCounty) BPitkyaranta皮特基亚兰塔() feud.Barony {
	return c.皮特基亚兰塔Pitkyaranta
}

func (c *凯姆KarelenCounty) BSordavala索尔塔瓦拉() feud.Barony {
	return c.索尔塔瓦拉Sordavala
}

var CKarelen凯姆 KarelenCounty = &凯姆KarelenCounty{}

func init() {
	f := CKarelen凯姆.(*凯姆KarelenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "388",
		Title:     "karelen",
		TitleName: "凯姆",
		TitleCode: "c_karelen",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡列瓦拉Kalevala = BKalevala卡列瓦拉
	f.卡列瓦拉Kalevala.SetParent(f)

	f.凯姆Kem = BKem凯姆
	f.凯姆Kem.SetParent(f)

	f.科斯托穆克沙Kostomuksha = BKostomuksha科斯托穆克沙
	f.科斯托穆克沙Kostomuksha.SetParent(f)

	f.洛乌希Loukhi = BLoukhi洛乌希
	f.洛乌希Loukhi.SetParent(f)

	f.穆耶泽尔斯基Muezersky = BMuezersky穆耶泽尔斯基
	f.穆耶泽尔斯基Muezersky.SetParent(f)

	f.皮特基亚兰塔Pitkyaranta = BPitkyaranta皮特基亚兰塔
	f.皮特基亚兰塔Pitkyaranta.SetParent(f)

	f.索尔塔瓦拉Sordavala = BSordavala索尔塔瓦拉
	f.索尔塔瓦拉Sordavala.SetParent(f)

}
