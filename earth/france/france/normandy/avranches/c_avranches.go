package avranches

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AvranchesCounty interface {
	feud.County
	BAvranches阿夫朗什() feud.Barony
	BBarfleur巴夫勒尔() feud.Barony
	BCarentan卡朗唐() feud.Barony
	BCherbourg瑟堡() feud.Barony
	BCoutances库唐斯() feud.Barony
	BMortain莫尔坦() feud.Barony
}

type 莫尔坦AvranchesCounty struct {
	feud.BaseCounty
	阿夫朗什Avranches feud.Barony
	巴夫勒尔Barfleur  feud.Barony
	卡朗唐Carentan   feud.Barony
	瑟堡Cherbourg   feud.Barony
	库唐斯Coutances  feud.Barony
	莫尔坦Mortain    feud.Barony
}

func (c *莫尔坦AvranchesCounty) BAvranches阿夫朗什() feud.Barony {
	return c.阿夫朗什Avranches
}

func (c *莫尔坦AvranchesCounty) BBarfleur巴夫勒尔() feud.Barony {
	return c.巴夫勒尔Barfleur
}

func (c *莫尔坦AvranchesCounty) BCarentan卡朗唐() feud.Barony {
	return c.卡朗唐Carentan
}

func (c *莫尔坦AvranchesCounty) BCherbourg瑟堡() feud.Barony {
	return c.瑟堡Cherbourg
}

func (c *莫尔坦AvranchesCounty) BCoutances库唐斯() feud.Barony {
	return c.库唐斯Coutances
}

func (c *莫尔坦AvranchesCounty) BMortain莫尔坦() feud.Barony {
	return c.莫尔坦Mortain
}

var CAvranches莫尔坦 AvranchesCounty = &莫尔坦AvranchesCounty{}

func init() {
	f := CAvranches莫尔坦.(*莫尔坦AvranchesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "100",
		Title:     "avranches",
		TitleName: "莫尔坦",
		TitleCode: "c_avranches",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿夫朗什Avranches = BAvranches阿夫朗什
	f.阿夫朗什Avranches.SetParent(f)

	f.巴夫勒尔Barfleur = BBarfleur巴夫勒尔
	f.巴夫勒尔Barfleur.SetParent(f)

	f.卡朗唐Carentan = BCarentan卡朗唐
	f.卡朗唐Carentan.SetParent(f)

	f.瑟堡Cherbourg = BCherbourg瑟堡
	f.瑟堡Cherbourg.SetParent(f)

	f.库唐斯Coutances = BCoutances库唐斯
	f.库唐斯Coutances.SetParent(f)

	f.莫尔坦Mortain = BMortain莫尔坦
	f.莫尔坦Mortain.SetParent(f)

}
