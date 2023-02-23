package balkonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BalkondaCounty interface {
	feud.County
	BBalkonda婆罗军荼() feud.Barony
	BBasara巴萨拉() feud.Barony
	BDharmapuri达摩补利() feud.Barony
	BKhairaput夏罗布特() feud.Barony
	BMahur玛赫尔() feud.Barony
	BManikdurg摩尼迦突伽() feud.Barony
	BNirmal尼尔默尔() feud.Barony
}

type 婆罗军荼BalkondaCounty struct {
	feud.BaseCounty
	婆罗军荼Balkonda   feud.Barony
	巴萨拉Basara      feud.Barony
	达摩补利Dharmapuri feud.Barony
	夏罗布特Khairaput  feud.Barony
	玛赫尔Mahur       feud.Barony
	摩尼迦突伽Manikdurg feud.Barony
	尼尔默尔Nirmal     feud.Barony
}

func (c *婆罗军荼BalkondaCounty) BBalkonda婆罗军荼() feud.Barony {
	return c.婆罗军荼Balkonda
}

func (c *婆罗军荼BalkondaCounty) BBasara巴萨拉() feud.Barony {
	return c.巴萨拉Basara
}

func (c *婆罗军荼BalkondaCounty) BDharmapuri达摩补利() feud.Barony {
	return c.达摩补利Dharmapuri
}

func (c *婆罗军荼BalkondaCounty) BKhairaput夏罗布特() feud.Barony {
	return c.夏罗布特Khairaput
}

func (c *婆罗军荼BalkondaCounty) BMahur玛赫尔() feud.Barony {
	return c.玛赫尔Mahur
}

func (c *婆罗军荼BalkondaCounty) BManikdurg摩尼迦突伽() feud.Barony {
	return c.摩尼迦突伽Manikdurg
}

func (c *婆罗军荼BalkondaCounty) BNirmal尼尔默尔() feud.Barony {
	return c.尼尔默尔Nirmal
}

var CBalkonda婆罗军荼 BalkondaCounty = &婆罗军荼BalkondaCounty{}

func init() {
	f := CBalkonda婆罗军荼.(*婆罗军荼BalkondaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1157",
		Title:     "balkonda",
		TitleName: "婆罗军荼",
		TitleCode: "c_balkonda",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆罗军荼Balkonda = BBalkonda婆罗军荼
	f.婆罗军荼Balkonda.SetParent(f)

	f.巴萨拉Basara = BBasara巴萨拉
	f.巴萨拉Basara.SetParent(f)

	f.达摩补利Dharmapuri = BDharmapuri达摩补利
	f.达摩补利Dharmapuri.SetParent(f)

	f.夏罗布特Khairaput = BKhairaput夏罗布特
	f.夏罗布特Khairaput.SetParent(f)

	f.玛赫尔Mahur = BMahur玛赫尔
	f.玛赫尔Mahur.SetParent(f)

	f.摩尼迦突伽Manikdurg = BManikdurg摩尼迦突伽
	f.摩尼迦突伽Manikdurg.SetParent(f)

	f.尼尔默尔Nirmal = BNirmal尼尔默尔
	f.尼尔默尔Nirmal.SetParent(f)

}
