package markakol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MarkakolCounty interface {
	feud.County
	BAltaya阿尔塔亚() feud.Barony
	BBykovo贝科沃() feud.Barony
	BMaleevsk马列耶夫斯克() feud.Barony
	BManat马纳特() feud.Barony
	BMarkakol马尔卡科尔() feud.Barony
	BOrnek奥尔内克() feud.Barony
	BZyryanovsk济良诺夫斯克() feud.Barony
}

type 马尔卡科尔MarkakolCounty struct {
	feud.BaseCounty
	阿尔塔亚Altaya       feud.Barony
	贝科沃Bykovo        feud.Barony
	马列耶夫斯克Maleevsk   feud.Barony
	马纳特Manat         feud.Barony
	马尔卡科尔Markakol    feud.Barony
	奥尔内克Ornek        feud.Barony
	济良诺夫斯克Zyryanovsk feud.Barony
}

func (c *马尔卡科尔MarkakolCounty) BAltaya阿尔塔亚() feud.Barony {
	return c.阿尔塔亚Altaya
}

func (c *马尔卡科尔MarkakolCounty) BBykovo贝科沃() feud.Barony {
	return c.贝科沃Bykovo
}

func (c *马尔卡科尔MarkakolCounty) BMaleevsk马列耶夫斯克() feud.Barony {
	return c.马列耶夫斯克Maleevsk
}

func (c *马尔卡科尔MarkakolCounty) BManat马纳特() feud.Barony {
	return c.马纳特Manat
}

func (c *马尔卡科尔MarkakolCounty) BMarkakol马尔卡科尔() feud.Barony {
	return c.马尔卡科尔Markakol
}

func (c *马尔卡科尔MarkakolCounty) BOrnek奥尔内克() feud.Barony {
	return c.奥尔内克Ornek
}

func (c *马尔卡科尔MarkakolCounty) BZyryanovsk济良诺夫斯克() feud.Barony {
	return c.济良诺夫斯克Zyryanovsk
}

var CMarkakol马尔卡科尔 MarkakolCounty = &马尔卡科尔MarkakolCounty{}

func init() {
	f := CMarkakol马尔卡科尔.(*马尔卡科尔MarkakolCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1898",
		Title:     "markakol",
		TitleName: "马尔卡科尔",
		TitleCode: "c_markakol",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔塔亚Altaya = BAltaya阿尔塔亚
	f.阿尔塔亚Altaya.SetParent(f)

	f.贝科沃Bykovo = BBykovo贝科沃
	f.贝科沃Bykovo.SetParent(f)

	f.马列耶夫斯克Maleevsk = BMaleevsk马列耶夫斯克
	f.马列耶夫斯克Maleevsk.SetParent(f)

	f.马纳特Manat = BManat马纳特
	f.马纳特Manat.SetParent(f)

	f.马尔卡科尔Markakol = BMarkakol马尔卡科尔
	f.马尔卡科尔Markakol.SetParent(f)

	f.奥尔内克Ornek = BOrnek奥尔内克
	f.奥尔内克Ornek.SetParent(f)

	f.济良诺夫斯克Zyryanovsk = BZyryanovsk济良诺夫斯克
	f.济良诺夫斯克Zyryanovsk.SetParent(f)

}
