package spoleto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SpoletoCounty interface {
	feud.County
	BCascia卡夏() feud.Barony
	BFoligno福利尼奥() feud.Barony
	BNursia努尔西亚() feud.Barony
	BSpoleto斯波莱托() feud.Barony
	BTerni特尔尼() feud.Barony
	BTodi托迪() feud.Barony
	BValva瓦尔瓦() feud.Barony
}

type 斯波莱托SpoletoCounty struct {
	feud.BaseCounty
	卡夏Cascia    feud.Barony
	福利尼奥Foligno feud.Barony
	努尔西亚Nursia  feud.Barony
	斯波莱托Spoleto feud.Barony
	特尔尼Terni    feud.Barony
	托迪Todi      feud.Barony
	瓦尔瓦Valva    feud.Barony
}

func (c *斯波莱托SpoletoCounty) BCascia卡夏() feud.Barony {
	return c.卡夏Cascia
}

func (c *斯波莱托SpoletoCounty) BFoligno福利尼奥() feud.Barony {
	return c.福利尼奥Foligno
}

func (c *斯波莱托SpoletoCounty) BNursia努尔西亚() feud.Barony {
	return c.努尔西亚Nursia
}

func (c *斯波莱托SpoletoCounty) BSpoleto斯波莱托() feud.Barony {
	return c.斯波莱托Spoleto
}

func (c *斯波莱托SpoletoCounty) BTerni特尔尼() feud.Barony {
	return c.特尔尼Terni
}

func (c *斯波莱托SpoletoCounty) BTodi托迪() feud.Barony {
	return c.托迪Todi
}

func (c *斯波莱托SpoletoCounty) BValva瓦尔瓦() feud.Barony {
	return c.瓦尔瓦Valva
}

var CSpoleto斯波莱托 SpoletoCounty = &斯波莱托SpoletoCounty{}

func init() {
	f := CSpoleto斯波莱托.(*斯波莱托SpoletoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "349",
		Title:     "spoleto",
		TitleName: "斯波莱托",
		TitleCode: "c_spoleto",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡夏Cascia = BCascia卡夏
	f.卡夏Cascia.SetParent(f)

	f.福利尼奥Foligno = BFoligno福利尼奥
	f.福利尼奥Foligno.SetParent(f)

	f.努尔西亚Nursia = BNursia努尔西亚
	f.努尔西亚Nursia.SetParent(f)

	f.斯波莱托Spoleto = BSpoleto斯波莱托
	f.斯波莱托Spoleto.SetParent(f)

	f.特尔尼Terni = BTerni特尔尼
	f.特尔尼Terni.SetParent(f)

	f.托迪Todi = BTodi托迪
	f.托迪Todi.SetParent(f)

	f.瓦尔瓦Valva = BValva瓦尔瓦
	f.瓦尔瓦Valva.SetParent(f)

}
