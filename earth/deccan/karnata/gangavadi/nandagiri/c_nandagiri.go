package nandagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NandagiriCounty interface {
	feud.County
	BAvani阿瓦尼() feud.Barony
	BKundani军荼尼() feud.Barony
	BKurudumale俱卢厨摩隶() feud.Barony
	BKuvalala鸠婆剌罗() feud.Barony
	BMulavagil莫拉瓦吉尔() feud.Barony
	BNandagiri难陀山() feud.Barony
	BNangali囊伽梨() feud.Barony
}

type 难陀山NandagiriCounty struct {
	feud.BaseCounty
	阿瓦尼Avani        feud.Barony
	军荼尼Kundani      feud.Barony
	俱卢厨摩隶Kurudumale feud.Barony
	鸠婆剌罗Kuvalala    feud.Barony
	莫拉瓦吉尔Mulavagil  feud.Barony
	难陀山Nandagiri    feud.Barony
	囊伽梨Nangali      feud.Barony
}

func (c *难陀山NandagiriCounty) BAvani阿瓦尼() feud.Barony {
	return c.阿瓦尼Avani
}

func (c *难陀山NandagiriCounty) BKundani军荼尼() feud.Barony {
	return c.军荼尼Kundani
}

func (c *难陀山NandagiriCounty) BKurudumale俱卢厨摩隶() feud.Barony {
	return c.俱卢厨摩隶Kurudumale
}

func (c *难陀山NandagiriCounty) BKuvalala鸠婆剌罗() feud.Barony {
	return c.鸠婆剌罗Kuvalala
}

func (c *难陀山NandagiriCounty) BMulavagil莫拉瓦吉尔() feud.Barony {
	return c.莫拉瓦吉尔Mulavagil
}

func (c *难陀山NandagiriCounty) BNandagiri难陀山() feud.Barony {
	return c.难陀山Nandagiri
}

func (c *难陀山NandagiriCounty) BNangali囊伽梨() feud.Barony {
	return c.囊伽梨Nangali
}

var CNandagiri难陀山 NandagiriCounty = &难陀山NandagiriCounty{}

func init() {
	f := CNandagiri难陀山.(*难陀山NandagiriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1217",
		Title:     "nandagiri",
		TitleName: "难陀山",
		TitleCode: "c_nandagiri",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿瓦尼Avani = BAvani阿瓦尼
	f.阿瓦尼Avani.SetParent(f)

	f.军荼尼Kundani = BKundani军荼尼
	f.军荼尼Kundani.SetParent(f)

	f.俱卢厨摩隶Kurudumale = BKurudumale俱卢厨摩隶
	f.俱卢厨摩隶Kurudumale.SetParent(f)

	f.鸠婆剌罗Kuvalala = BKuvalala鸠婆剌罗
	f.鸠婆剌罗Kuvalala.SetParent(f)

	f.莫拉瓦吉尔Mulavagil = BMulavagil莫拉瓦吉尔
	f.莫拉瓦吉尔Mulavagil.SetParent(f)

	f.难陀山Nandagiri = BNandagiri难陀山
	f.难陀山Nandagiri.SetParent(f)

	f.囊伽梨Nangali = BNangali囊伽梨
	f.囊伽梨Nangali.SetParent(f)

}
