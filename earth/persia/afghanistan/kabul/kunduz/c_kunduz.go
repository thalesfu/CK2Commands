package kunduz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KunduzCounty interface {
	feud.County
	BAibak艾伊拜克() feud.Barony
	BBaghlan缚伽浪() feud.Barony
	BKhomri胡姆里() feud.Barony
	BKunduz昆都士() feud.Barony
	BMogholan蒙哥兰() feud.Barony
	BNahrain那赫林() feud.Barony
	BSiminjan悉泯健() feud.Barony
}

type 昆都士KunduzCounty struct {
	feud.BaseCounty
	艾伊拜克Aibak   feud.Barony
	缚伽浪Baghlan  feud.Barony
	胡姆里Khomri   feud.Barony
	昆都士Kunduz   feud.Barony
	蒙哥兰Mogholan feud.Barony
	那赫林Nahrain  feud.Barony
	悉泯健Siminjan feud.Barony
}

func (c *昆都士KunduzCounty) BAibak艾伊拜克() feud.Barony {
	return c.艾伊拜克Aibak
}

func (c *昆都士KunduzCounty) BBaghlan缚伽浪() feud.Barony {
	return c.缚伽浪Baghlan
}

func (c *昆都士KunduzCounty) BKhomri胡姆里() feud.Barony {
	return c.胡姆里Khomri
}

func (c *昆都士KunduzCounty) BKunduz昆都士() feud.Barony {
	return c.昆都士Kunduz
}

func (c *昆都士KunduzCounty) BMogholan蒙哥兰() feud.Barony {
	return c.蒙哥兰Mogholan
}

func (c *昆都士KunduzCounty) BNahrain那赫林() feud.Barony {
	return c.那赫林Nahrain
}

func (c *昆都士KunduzCounty) BSiminjan悉泯健() feud.Barony {
	return c.悉泯健Siminjan
}

var CKunduz昆都士 KunduzCounty = &昆都士KunduzCounty{}

func init() {
	f := CKunduz昆都士.(*昆都士KunduzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1187",
		Title:     "kunduz",
		TitleName: "昆都士",
		TitleCode: "c_kunduz",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾伊拜克Aibak = BAibak艾伊拜克
	f.艾伊拜克Aibak.SetParent(f)

	f.缚伽浪Baghlan = BBaghlan缚伽浪
	f.缚伽浪Baghlan.SetParent(f)

	f.胡姆里Khomri = BKhomri胡姆里
	f.胡姆里Khomri.SetParent(f)

	f.昆都士Kunduz = BKunduz昆都士
	f.昆都士Kunduz.SetParent(f)

	f.蒙哥兰Mogholan = BMogholan蒙哥兰
	f.蒙哥兰Mogholan.SetParent(f)

	f.那赫林Nahrain = BNahrain那赫林
	f.那赫林Nahrain.SetParent(f)

	f.悉泯健Siminjan = BSiminjan悉泯健
	f.悉泯健Siminjan.SetParent(f)

}
