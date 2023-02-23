package dodugu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DoduguCounty interface {
	feud.County
	BBamako巴马科() feud.Barony
	BKangaba康加巴() feud.Barony
	BManankoro马南科罗() feud.Barony
	BNamandiguina纳曼迪吉纳() feud.Barony
	BNdantari恩丹塔里() feud.Barony
	BTichilit蒂希利特() feud.Barony
}

type 多杜古DoduguCounty struct {
	feud.BaseCounty
	巴马科Bamako         feud.Barony
	康加巴Kangaba        feud.Barony
	马南科罗Manankoro     feud.Barony
	纳曼迪吉纳Namandiguina feud.Barony
	恩丹塔里Ndantari      feud.Barony
	蒂希利特Tichilit      feud.Barony
}

func (c *多杜古DoduguCounty) BBamako巴马科() feud.Barony {
	return c.巴马科Bamako
}

func (c *多杜古DoduguCounty) BKangaba康加巴() feud.Barony {
	return c.康加巴Kangaba
}

func (c *多杜古DoduguCounty) BManankoro马南科罗() feud.Barony {
	return c.马南科罗Manankoro
}

func (c *多杜古DoduguCounty) BNamandiguina纳曼迪吉纳() feud.Barony {
	return c.纳曼迪吉纳Namandiguina
}

func (c *多杜古DoduguCounty) BNdantari恩丹塔里() feud.Barony {
	return c.恩丹塔里Ndantari
}

func (c *多杜古DoduguCounty) BTichilit蒂希利特() feud.Barony {
	return c.蒂希利特Tichilit
}

var CDodugu多杜古 DoduguCounty = &多杜古DoduguCounty{}

func init() {
	f := CDodugu多杜古.(*多杜古DoduguCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1763",
		Title:     "dodugu",
		TitleName: "多杜古",
		TitleCode: "c_dodugu",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴马科Bamako = BBamako巴马科
	f.巴马科Bamako.SetParent(f)

	f.康加巴Kangaba = BKangaba康加巴
	f.康加巴Kangaba.SetParent(f)

	f.马南科罗Manankoro = BManankoro马南科罗
	f.马南科罗Manankoro.SetParent(f)

	f.纳曼迪吉纳Namandiguina = BNamandiguina纳曼迪吉纳
	f.纳曼迪吉纳Namandiguina.SetParent(f)

	f.恩丹塔里Ndantari = BNdantari恩丹塔里
	f.恩丹塔里Ndantari.SetParent(f)

	f.蒂希利特Tichilit = BTichilit蒂希利特
	f.蒂希利特Tichilit.SetParent(f)

}
