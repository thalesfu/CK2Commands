package cherchen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CherchenCounty interface {
	feud.County
	BCalmandana折摩驮那() feud.Barony
	BCherchen车尔臣() feud.Barony
	BMocheng末城() feud.Barony
	BSutang苏塘() feud.Barony
	BYeyik叶亦克() feud.Barony
	BYuling扜零() feud.Barony
	BZuomo左末() feud.Barony
}

type 车尔臣CherchenCounty struct {
	feud.BaseCounty
	折摩驮那Calmandana feud.Barony
	车尔臣Cherchen    feud.Barony
	末城Mocheng      feud.Barony
	苏塘Sutang       feud.Barony
	叶亦克Yeyik       feud.Barony
	扜零Yuling       feud.Barony
	左末Zuomo        feud.Barony
}

func (c *车尔臣CherchenCounty) BCalmandana折摩驮那() feud.Barony {
	return c.折摩驮那Calmandana
}

func (c *车尔臣CherchenCounty) BCherchen车尔臣() feud.Barony {
	return c.车尔臣Cherchen
}

func (c *车尔臣CherchenCounty) BMocheng末城() feud.Barony {
	return c.末城Mocheng
}

func (c *车尔臣CherchenCounty) BSutang苏塘() feud.Barony {
	return c.苏塘Sutang
}

func (c *车尔臣CherchenCounty) BYeyik叶亦克() feud.Barony {
	return c.叶亦克Yeyik
}

func (c *车尔臣CherchenCounty) BYuling扜零() feud.Barony {
	return c.扜零Yuling
}

func (c *车尔臣CherchenCounty) BZuomo左末() feud.Barony {
	return c.左末Zuomo
}

var CCherchen车尔臣 CherchenCounty = &车尔臣CherchenCounty{}

func init() {
	f := CCherchen车尔臣.(*车尔臣CherchenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1441",
		Title:     "cherchen",
		TitleName: "车尔臣",
		TitleCode: "c_cherchen",
		Baronies:  map[string]feud.Barony{},
	}

	f.折摩驮那Calmandana = BCalmandana折摩驮那
	f.折摩驮那Calmandana.SetParent(f)

	f.车尔臣Cherchen = BCherchen车尔臣
	f.车尔臣Cherchen.SetParent(f)

	f.末城Mocheng = BMocheng末城
	f.末城Mocheng.SetParent(f)

	f.苏塘Sutang = BSutang苏塘
	f.苏塘Sutang.SetParent(f)

	f.叶亦克Yeyik = BYeyik叶亦克
	f.叶亦克Yeyik.SetParent(f)

	f.扜零Yuling = BYuling扜零
	f.扜零Yuling.SetParent(f)

	f.左末Zuomo = BZuomo左末
	f.左末Zuomo.SetParent(f)

}
