package kundina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KundinaCounty interface {
	feud.County
	BKundina军稚那() feud.Barony
	BNauli瑙力() feud.Barony
	BNijhar尼旬罗() feud.Barony
	BParavar钵罗跋罗() feud.Barony
	BPrithiminagar毕哩体弥那揭罗() feud.Barony
	BSadiya娑提耶() feud.Barony
	BTinsukia提那须吉阿() feud.Barony
}

type 军稚那KundinaCounty struct {
	feud.BaseCounty
	军稚那Kundina           feud.Barony
	瑙力Nauli              feud.Barony
	尼旬罗Nijhar            feud.Barony
	钵罗跋罗Paravar          feud.Barony
	毕哩体弥那揭罗Prithiminagar feud.Barony
	娑提耶Sadiya            feud.Barony
	提那须吉阿Tinsukia        feud.Barony
}

func (c *军稚那KundinaCounty) BKundina军稚那() feud.Barony {
	return c.军稚那Kundina
}

func (c *军稚那KundinaCounty) BNauli瑙力() feud.Barony {
	return c.瑙力Nauli
}

func (c *军稚那KundinaCounty) BNijhar尼旬罗() feud.Barony {
	return c.尼旬罗Nijhar
}

func (c *军稚那KundinaCounty) BParavar钵罗跋罗() feud.Barony {
	return c.钵罗跋罗Paravar
}

func (c *军稚那KundinaCounty) BPrithiminagar毕哩体弥那揭罗() feud.Barony {
	return c.毕哩体弥那揭罗Prithiminagar
}

func (c *军稚那KundinaCounty) BSadiya娑提耶() feud.Barony {
	return c.娑提耶Sadiya
}

func (c *军稚那KundinaCounty) BTinsukia提那须吉阿() feud.Barony {
	return c.提那须吉阿Tinsukia
}

var CKundina军稚那 KundinaCounty = &军稚那KundinaCounty{}

func init() {
	f := CKundina军稚那.(*军稚那KundinaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1177",
		Title:     "kundina",
		TitleName: "军稚那",
		TitleCode: "c_kundina",
		Baronies:  map[string]feud.Barony{},
	}

	f.军稚那Kundina = BKundina军稚那
	f.军稚那Kundina.SetParent(f)

	f.瑙力Nauli = BNauli瑙力
	f.瑙力Nauli.SetParent(f)

	f.尼旬罗Nijhar = BNijhar尼旬罗
	f.尼旬罗Nijhar.SetParent(f)

	f.钵罗跋罗Paravar = BParavar钵罗跋罗
	f.钵罗跋罗Paravar.SetParent(f)

	f.毕哩体弥那揭罗Prithiminagar = BPrithiminagar毕哩体弥那揭罗
	f.毕哩体弥那揭罗Prithiminagar.SetParent(f)

	f.娑提耶Sadiya = BSadiya娑提耶
	f.娑提耶Sadiya.SetParent(f)

	f.提那须吉阿Tinsukia = BTinsukia提那须吉阿
	f.提那须吉阿Tinsukia.SetParent(f)

}
