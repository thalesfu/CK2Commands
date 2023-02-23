package belgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BelgorodCounty interface {
	feud.County
	BBelgorod别尔哥罗德() feud.Barony
	BCapriana卡普里亚纳() feud.Barony
	BCauseni克乌谢尼() feud.Barony
	BCedaesti切达埃斯特() feud.Barony
	BChisinau基希讷乌() feud.Barony
	BTighina蒂吉纳() feud.Barony
}

type 别尔哥罗德BelgorodCounty struct {
	feud.BaseCounty
	别尔哥罗德Belgorod feud.Barony
	卡普里亚纳Capriana feud.Barony
	克乌谢尼Causeni   feud.Barony
	切达埃斯特Cedaesti feud.Barony
	基希讷乌Chisinau  feud.Barony
	蒂吉纳Tighina    feud.Barony
}

func (c *别尔哥罗德BelgorodCounty) BBelgorod别尔哥罗德() feud.Barony {
	return c.别尔哥罗德Belgorod
}

func (c *别尔哥罗德BelgorodCounty) BCapriana卡普里亚纳() feud.Barony {
	return c.卡普里亚纳Capriana
}

func (c *别尔哥罗德BelgorodCounty) BCauseni克乌谢尼() feud.Barony {
	return c.克乌谢尼Causeni
}

func (c *别尔哥罗德BelgorodCounty) BCedaesti切达埃斯特() feud.Barony {
	return c.切达埃斯特Cedaesti
}

func (c *别尔哥罗德BelgorodCounty) BChisinau基希讷乌() feud.Barony {
	return c.基希讷乌Chisinau
}

func (c *别尔哥罗德BelgorodCounty) BTighina蒂吉纳() feud.Barony {
	return c.蒂吉纳Tighina
}

var CBelgorod别尔哥罗德 BelgorodCounty = &别尔哥罗德BelgorodCounty{}

func init() {
	f := CBelgorod别尔哥罗德.(*别尔哥罗德BelgorodCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "512",
		Title:     "belgorod",
		TitleName: "别尔哥罗德",
		TitleCode: "c_belgorod",
		Baronies:  map[string]feud.Barony{},
	}

	f.别尔哥罗德Belgorod = BBelgorod别尔哥罗德
	f.别尔哥罗德Belgorod.SetParent(f)

	f.卡普里亚纳Capriana = BCapriana卡普里亚纳
	f.卡普里亚纳Capriana.SetParent(f)

	f.克乌谢尼Causeni = BCauseni克乌谢尼
	f.克乌谢尼Causeni.SetParent(f)

	f.切达埃斯特Cedaesti = BCedaesti切达埃斯特
	f.切达埃斯特Cedaesti.SetParent(f)

	f.基希讷乌Chisinau = BChisinau基希讷乌
	f.基希讷乌Chisinau.SetParent(f)

	f.蒂吉纳Tighina = BTighina蒂吉纳
	f.蒂吉纳Tighina.SetParent(f)

}
