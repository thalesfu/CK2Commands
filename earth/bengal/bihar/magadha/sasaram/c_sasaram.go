package sasaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SasaramCounty interface {
	feud.County
	BArhu阿尔胡() feud.Barony
	BArrah阿罗() feud.Barony
	BHorilnagar呼利罗城() feud.Barony
	BJaund善陀() feud.Barony
	BKapadvanj伽波陀槃阇() feud.Barony
	BMundeshwari蒙提湿伐罗() feud.Barony
	BSasaram萨娑蓝() feud.Barony
}

type 萨娑蓝SasaramCounty struct {
	feud.BaseCounty
	阿尔胡Arhu          feud.Barony
	阿罗Arrah          feud.Barony
	呼利罗城Horilnagar   feud.Barony
	善陀Jaund          feud.Barony
	伽波陀槃阇Kapadvanj   feud.Barony
	蒙提湿伐罗Mundeshwari feud.Barony
	萨娑蓝Sasaram       feud.Barony
}

func (c *萨娑蓝SasaramCounty) BArhu阿尔胡() feud.Barony {
	return c.阿尔胡Arhu
}

func (c *萨娑蓝SasaramCounty) BArrah阿罗() feud.Barony {
	return c.阿罗Arrah
}

func (c *萨娑蓝SasaramCounty) BHorilnagar呼利罗城() feud.Barony {
	return c.呼利罗城Horilnagar
}

func (c *萨娑蓝SasaramCounty) BJaund善陀() feud.Barony {
	return c.善陀Jaund
}

func (c *萨娑蓝SasaramCounty) BKapadvanj伽波陀槃阇() feud.Barony {
	return c.伽波陀槃阇Kapadvanj
}

func (c *萨娑蓝SasaramCounty) BMundeshwari蒙提湿伐罗() feud.Barony {
	return c.蒙提湿伐罗Mundeshwari
}

func (c *萨娑蓝SasaramCounty) BSasaram萨娑蓝() feud.Barony {
	return c.萨娑蓝Sasaram
}

var CSasaram萨娑蓝 SasaramCounty = &萨娑蓝SasaramCounty{}

func init() {
	f := CSasaram萨娑蓝.(*萨娑蓝SasaramCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1251",
		Title:     "sasaram",
		TitleName: "萨娑蓝",
		TitleCode: "c_sasaram",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔胡Arhu = BArhu阿尔胡
	f.阿尔胡Arhu.SetParent(f)

	f.阿罗Arrah = BArrah阿罗
	f.阿罗Arrah.SetParent(f)

	f.呼利罗城Horilnagar = BHorilnagar呼利罗城
	f.呼利罗城Horilnagar.SetParent(f)

	f.善陀Jaund = BJaund善陀
	f.善陀Jaund.SetParent(f)

	f.伽波陀槃阇Kapadvanj = BKapadvanj伽波陀槃阇
	f.伽波陀槃阇Kapadvanj.SetParent(f)

	f.蒙提湿伐罗Mundeshwari = BMundeshwari蒙提湿伐罗
	f.蒙提湿伐罗Mundeshwari.SetParent(f)

	f.萨娑蓝Sasaram = BSasaram萨娑蓝
	f.萨娑蓝Sasaram.SetParent(f)

}
