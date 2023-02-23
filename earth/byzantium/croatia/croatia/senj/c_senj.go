package senj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SenjCounty interface {
	feud.County
	BBrinje布里涅() feud.Barony
	BDonjilapac多尼拉帕茨() feud.Barony
	BKarlobag卡尔洛巴格() feud.Barony
	BKaseg卡塞格() feud.Barony
	BOtocac奥托查茨() feud.Barony
	BPerusic佩鲁希奇() feud.Barony
	BSenj塞尼() feud.Barony
	BUdbina乌德比纳() feud.Barony
}

type 塞尼SenjCounty struct {
	feud.BaseCounty
	布里涅Brinje       feud.Barony
	多尼拉帕茨Donjilapac feud.Barony
	卡尔洛巴格Karlobag   feud.Barony
	卡塞格Kaseg        feud.Barony
	奥托查茨Otocac      feud.Barony
	佩鲁希奇Perusic     feud.Barony
	塞尼Senj          feud.Barony
	乌德比纳Udbina      feud.Barony
}

func (c *塞尼SenjCounty) BBrinje布里涅() feud.Barony {
	return c.布里涅Brinje
}

func (c *塞尼SenjCounty) BDonjilapac多尼拉帕茨() feud.Barony {
	return c.多尼拉帕茨Donjilapac
}

func (c *塞尼SenjCounty) BKarlobag卡尔洛巴格() feud.Barony {
	return c.卡尔洛巴格Karlobag
}

func (c *塞尼SenjCounty) BKaseg卡塞格() feud.Barony {
	return c.卡塞格Kaseg
}

func (c *塞尼SenjCounty) BOtocac奥托查茨() feud.Barony {
	return c.奥托查茨Otocac
}

func (c *塞尼SenjCounty) BPerusic佩鲁希奇() feud.Barony {
	return c.佩鲁希奇Perusic
}

func (c *塞尼SenjCounty) BSenj塞尼() feud.Barony {
	return c.塞尼Senj
}

func (c *塞尼SenjCounty) BUdbina乌德比纳() feud.Barony {
	return c.乌德比纳Udbina
}

var CSenj塞尼 SenjCounty = &塞尼SenjCounty{}

func init() {
	f := CSenj塞尼.(*塞尼SenjCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "464",
		Title:     "senj",
		TitleName: "塞尼",
		TitleCode: "c_senj",
		Baronies:  map[string]feud.Barony{},
	}

	f.布里涅Brinje = BBrinje布里涅
	f.布里涅Brinje.SetParent(f)

	f.多尼拉帕茨Donjilapac = BDonjilapac多尼拉帕茨
	f.多尼拉帕茨Donjilapac.SetParent(f)

	f.卡尔洛巴格Karlobag = BKarlobag卡尔洛巴格
	f.卡尔洛巴格Karlobag.SetParent(f)

	f.卡塞格Kaseg = BKaseg卡塞格
	f.卡塞格Kaseg.SetParent(f)

	f.奥托查茨Otocac = BOtocac奥托查茨
	f.奥托查茨Otocac.SetParent(f)

	f.佩鲁希奇Perusic = BPerusic佩鲁希奇
	f.佩鲁希奇Perusic.SetParent(f)

	f.塞尼Senj = BSenj塞尼
	f.塞尼Senj.SetParent(f)

	f.乌德比纳Udbina = BUdbina乌德比纳
	f.乌德比纳Udbina.SetParent(f)

}
