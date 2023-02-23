package hradyzk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HradyzkCounty interface {
	feud.County
	BHradyzk格拉季济克() feud.Barony
	BKhorol霍罗尔() feud.Barony
	BKobeliaky科别利亚基() feud.Barony
	BKremenchuk克列缅丘格() feud.Barony
	BLubny卢布内() feud.Barony
	BMyrhorod米尔哥罗德() feud.Barony
	BPoltava波尔塔瓦() feud.Barony
}

type 格拉季济克HradyzkCounty struct {
	feud.BaseCounty
	格拉季济克Hradyzk    feud.Barony
	霍罗尔Khorol       feud.Barony
	科别利亚基Kobeliaky  feud.Barony
	克列缅丘格Kremenchuk feud.Barony
	卢布内Lubny        feud.Barony
	米尔哥罗德Myrhorod   feud.Barony
	波尔塔瓦Poltava     feud.Barony
}

func (c *格拉季济克HradyzkCounty) BHradyzk格拉季济克() feud.Barony {
	return c.格拉季济克Hradyzk
}

func (c *格拉季济克HradyzkCounty) BKhorol霍罗尔() feud.Barony {
	return c.霍罗尔Khorol
}

func (c *格拉季济克HradyzkCounty) BKobeliaky科别利亚基() feud.Barony {
	return c.科别利亚基Kobeliaky
}

func (c *格拉季济克HradyzkCounty) BKremenchuk克列缅丘格() feud.Barony {
	return c.克列缅丘格Kremenchuk
}

func (c *格拉季济克HradyzkCounty) BLubny卢布内() feud.Barony {
	return c.卢布内Lubny
}

func (c *格拉季济克HradyzkCounty) BMyrhorod米尔哥罗德() feud.Barony {
	return c.米尔哥罗德Myrhorod
}

func (c *格拉季济克HradyzkCounty) BPoltava波尔塔瓦() feud.Barony {
	return c.波尔塔瓦Poltava
}

var CHradyzk格拉季济克 HradyzkCounty = &格拉季济克HradyzkCounty{}

func init() {
	f := CHradyzk格拉季济克.(*格拉季济克HradyzkCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1658",
		Title:     "hradyzk",
		TitleName: "格拉季济克",
		TitleCode: "c_hradyzk",
		Baronies:  map[string]feud.Barony{},
	}

	f.格拉季济克Hradyzk = BHradyzk格拉季济克
	f.格拉季济克Hradyzk.SetParent(f)

	f.霍罗尔Khorol = BKhorol霍罗尔
	f.霍罗尔Khorol.SetParent(f)

	f.科别利亚基Kobeliaky = BKobeliaky科别利亚基
	f.科别利亚基Kobeliaky.SetParent(f)

	f.克列缅丘格Kremenchuk = BKremenchuk克列缅丘格
	f.克列缅丘格Kremenchuk.SetParent(f)

	f.卢布内Lubny = BLubny卢布内
	f.卢布内Lubny.SetParent(f)

	f.米尔哥罗德Myrhorod = BMyrhorod米尔哥罗德
	f.米尔哥罗德Myrhorod.SetParent(f)

	f.波尔塔瓦Poltava = BPoltava波尔塔瓦
	f.波尔塔瓦Poltava.SetParent(f)

}
