package djimi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DjimiCounty interface {
	feud.County
	BAuno奥诺() feud.Barony
	BDagolu达戈卢() feud.Barony
	BDjimi吉米() feud.Barony
	BKagusti卡古斯提() feud.Barony
	BMao马奥() feud.Barony
	BSandu桑杜() feud.Barony
}

type 吉米DjimiCounty struct {
	feud.BaseCounty
	奥诺Auno      feud.Barony
	达戈卢Dagolu   feud.Barony
	吉米Djimi     feud.Barony
	卡古斯提Kagusti feud.Barony
	马奥Mao       feud.Barony
	桑杜Sandu     feud.Barony
}

func (c *吉米DjimiCounty) BAuno奥诺() feud.Barony {
	return c.奥诺Auno
}

func (c *吉米DjimiCounty) BDagolu达戈卢() feud.Barony {
	return c.达戈卢Dagolu
}

func (c *吉米DjimiCounty) BDjimi吉米() feud.Barony {
	return c.吉米Djimi
}

func (c *吉米DjimiCounty) BKagusti卡古斯提() feud.Barony {
	return c.卡古斯提Kagusti
}

func (c *吉米DjimiCounty) BMao马奥() feud.Barony {
	return c.马奥Mao
}

func (c *吉米DjimiCounty) BSandu桑杜() feud.Barony {
	return c.桑杜Sandu
}

var CDjimi吉米 DjimiCounty = &吉米DjimiCounty{}

func init() {
	f := CDjimi吉米.(*吉米DjimiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1739",
		Title:     "djimi",
		TitleName: "吉米",
		TitleCode: "c_djimi",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥诺Auno = BAuno奥诺
	f.奥诺Auno.SetParent(f)

	f.达戈卢Dagolu = BDagolu达戈卢
	f.达戈卢Dagolu.SetParent(f)

	f.吉米Djimi = BDjimi吉米
	f.吉米Djimi.SetParent(f)

	f.卡古斯提Kagusti = BKagusti卡古斯提
	f.卡古斯提Kagusti.SetParent(f)

	f.马奥Mao = BMao马奥
	f.马奥Mao.SetParent(f)

	f.桑杜Sandu = BSandu桑杜
	f.桑杜Sandu.SetParent(f)

}
