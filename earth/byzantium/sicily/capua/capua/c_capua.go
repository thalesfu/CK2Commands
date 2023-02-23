package capua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CapuaCounty interface {
	feud.County
	BAcerra阿切拉() feud.Barony
	BAquino阿奎诺() feud.Barony
	BCalvi卡尔维() feud.Barony
	BCapua卡普阿() feud.Barony
	BCaserta卡塞塔() feud.Barony
	BGaeta加埃塔() feud.Barony
	BMontecassino卡西诺山() feud.Barony
	BTeano泰阿诺() feud.Barony
}

type 卡普阿CapuaCounty struct {
	feud.BaseCounty
	阿切拉Acerra        feud.Barony
	阿奎诺Aquino        feud.Barony
	卡尔维Calvi         feud.Barony
	卡普阿Capua         feud.Barony
	卡塞塔Caserta       feud.Barony
	加埃塔Gaeta         feud.Barony
	卡西诺山Montecassino feud.Barony
	泰阿诺Teano         feud.Barony
}

func (c *卡普阿CapuaCounty) BAcerra阿切拉() feud.Barony {
	return c.阿切拉Acerra
}

func (c *卡普阿CapuaCounty) BAquino阿奎诺() feud.Barony {
	return c.阿奎诺Aquino
}

func (c *卡普阿CapuaCounty) BCalvi卡尔维() feud.Barony {
	return c.卡尔维Calvi
}

func (c *卡普阿CapuaCounty) BCapua卡普阿() feud.Barony {
	return c.卡普阿Capua
}

func (c *卡普阿CapuaCounty) BCaserta卡塞塔() feud.Barony {
	return c.卡塞塔Caserta
}

func (c *卡普阿CapuaCounty) BGaeta加埃塔() feud.Barony {
	return c.加埃塔Gaeta
}

func (c *卡普阿CapuaCounty) BMontecassino卡西诺山() feud.Barony {
	return c.卡西诺山Montecassino
}

func (c *卡普阿CapuaCounty) BTeano泰阿诺() feud.Barony {
	return c.泰阿诺Teano
}

var CCapua卡普阿 CapuaCounty = &卡普阿CapuaCounty{}

func init() {
	f := CCapua卡普阿.(*卡普阿CapuaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "851",
		Title:     "capua",
		TitleName: "卡普阿",
		TitleCode: "c_capua",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿切拉Acerra = BAcerra阿切拉
	f.阿切拉Acerra.SetParent(f)

	f.阿奎诺Aquino = BAquino阿奎诺
	f.阿奎诺Aquino.SetParent(f)

	f.卡尔维Calvi = BCalvi卡尔维
	f.卡尔维Calvi.SetParent(f)

	f.卡普阿Capua = BCapua卡普阿
	f.卡普阿Capua.SetParent(f)

	f.卡塞塔Caserta = BCaserta卡塞塔
	f.卡塞塔Caserta.SetParent(f)

	f.加埃塔Gaeta = BGaeta加埃塔
	f.加埃塔Gaeta.SetParent(f)

	f.卡西诺山Montecassino = BMontecassino卡西诺山
	f.卡西诺山Montecassino.SetParent(f)

	f.泰阿诺Teano = BTeano泰阿诺
	f.泰阿诺Teano.SetParent(f)

}
