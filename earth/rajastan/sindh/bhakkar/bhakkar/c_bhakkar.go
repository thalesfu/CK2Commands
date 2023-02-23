package bhakkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BhakkarCounty interface {
	feud.County
	BBhakkar珀格尔() feud.Barony
	BKalari卡拉里() feud.Barony
	BLarkana拉卡纳() feud.Barony
	BMejia摩吉阿() feud.Barony
	BSarkara煞割令() feud.Barony
	BShikarpur尸佉罗补罗() feud.Barony
	BSialkot犀耶罗拘吒() feud.Barony
}

type 珀格尔BhakkarCounty struct {
	feud.BaseCounty
	珀格尔Bhakkar     feud.Barony
	卡拉里Kalari      feud.Barony
	拉卡纳Larkana     feud.Barony
	摩吉阿Mejia       feud.Barony
	煞割令Sarkara     feud.Barony
	尸佉罗补罗Shikarpur feud.Barony
	犀耶罗拘吒Sialkot   feud.Barony
}

func (c *珀格尔BhakkarCounty) BBhakkar珀格尔() feud.Barony {
	return c.珀格尔Bhakkar
}

func (c *珀格尔BhakkarCounty) BKalari卡拉里() feud.Barony {
	return c.卡拉里Kalari
}

func (c *珀格尔BhakkarCounty) BLarkana拉卡纳() feud.Barony {
	return c.拉卡纳Larkana
}

func (c *珀格尔BhakkarCounty) BMejia摩吉阿() feud.Barony {
	return c.摩吉阿Mejia
}

func (c *珀格尔BhakkarCounty) BSarkara煞割令() feud.Barony {
	return c.煞割令Sarkara
}

func (c *珀格尔BhakkarCounty) BShikarpur尸佉罗补罗() feud.Barony {
	return c.尸佉罗补罗Shikarpur
}

func (c *珀格尔BhakkarCounty) BSialkot犀耶罗拘吒() feud.Barony {
	return c.犀耶罗拘吒Sialkot
}

var CBhakkar珀格尔 BhakkarCounty = &珀格尔BhakkarCounty{}

func init() {
	f := CBhakkar珀格尔.(*珀格尔BhakkarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1138",
		Title:     "bhakkar",
		TitleName: "珀格尔",
		TitleCode: "c_bhakkar",
		Baronies:  map[string]feud.Barony{},
	}

	f.珀格尔Bhakkar = BBhakkar珀格尔
	f.珀格尔Bhakkar.SetParent(f)

	f.卡拉里Kalari = BKalari卡拉里
	f.卡拉里Kalari.SetParent(f)

	f.拉卡纳Larkana = BLarkana拉卡纳
	f.拉卡纳Larkana.SetParent(f)

	f.摩吉阿Mejia = BMejia摩吉阿
	f.摩吉阿Mejia.SetParent(f)

	f.煞割令Sarkara = BSarkara煞割令
	f.煞割令Sarkara.SetParent(f)

	f.尸佉罗补罗Shikarpur = BShikarpur尸佉罗补罗
	f.尸佉罗补罗Shikarpur.SetParent(f)

	f.犀耶罗拘吒Sialkot = BSialkot犀耶罗拘吒
	f.犀耶罗拘吒Sialkot.SetParent(f)

}
