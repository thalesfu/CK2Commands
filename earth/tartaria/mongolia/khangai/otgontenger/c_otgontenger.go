package otgontenger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OtgontengerCounty interface {
	feud.County
	BBayanbulag巴彦布拉格() feud.Barony
	BGurvanbulag古尔班布拉格() feud.Barony
	BGuulin古林() feud.Barony
	BOtgontenger鄂特冈腾格里() feud.Barony
	BTaishir泰希尔() feud.Barony
	BTsagaanchuulut察干楚卢特() feud.Barony
	BZag扎格() feud.Barony
}

type 鄂特冈腾格里OtgontengerCounty struct {
	feud.BaseCounty
	巴彦布拉格Bayanbulag     feud.Barony
	古尔班布拉格Gurvanbulag   feud.Barony
	古林Guulin            feud.Barony
	鄂特冈腾格里Otgontenger   feud.Barony
	泰希尔Taishir          feud.Barony
	察干楚卢特Tsagaanchuulut feud.Barony
	扎格Zag               feud.Barony
}

func (c *鄂特冈腾格里OtgontengerCounty) BBayanbulag巴彦布拉格() feud.Barony {
	return c.巴彦布拉格Bayanbulag
}

func (c *鄂特冈腾格里OtgontengerCounty) BGurvanbulag古尔班布拉格() feud.Barony {
	return c.古尔班布拉格Gurvanbulag
}

func (c *鄂特冈腾格里OtgontengerCounty) BGuulin古林() feud.Barony {
	return c.古林Guulin
}

func (c *鄂特冈腾格里OtgontengerCounty) BOtgontenger鄂特冈腾格里() feud.Barony {
	return c.鄂特冈腾格里Otgontenger
}

func (c *鄂特冈腾格里OtgontengerCounty) BTaishir泰希尔() feud.Barony {
	return c.泰希尔Taishir
}

func (c *鄂特冈腾格里OtgontengerCounty) BTsagaanchuulut察干楚卢特() feud.Barony {
	return c.察干楚卢特Tsagaanchuulut
}

func (c *鄂特冈腾格里OtgontengerCounty) BZag扎格() feud.Barony {
	return c.扎格Zag
}

var COtgontenger鄂特冈腾格里 OtgontengerCounty = &鄂特冈腾格里OtgontengerCounty{}

func init() {
	f := COtgontenger鄂特冈腾格里.(*鄂特冈腾格里OtgontengerCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1912",
		Title:     "otgontenger",
		TitleName: "鄂特冈腾格里",
		TitleCode: "c_otgontenger",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴彦布拉格Bayanbulag = BBayanbulag巴彦布拉格
	f.巴彦布拉格Bayanbulag.SetParent(f)

	f.古尔班布拉格Gurvanbulag = BGurvanbulag古尔班布拉格
	f.古尔班布拉格Gurvanbulag.SetParent(f)

	f.古林Guulin = BGuulin古林
	f.古林Guulin.SetParent(f)

	f.鄂特冈腾格里Otgontenger = BOtgontenger鄂特冈腾格里
	f.鄂特冈腾格里Otgontenger.SetParent(f)

	f.泰希尔Taishir = BTaishir泰希尔
	f.泰希尔Taishir.SetParent(f)

	f.察干楚卢特Tsagaanchuulut = BTsagaanchuulut察干楚卢特
	f.察干楚卢特Tsagaanchuulut.SetParent(f)

	f.扎格Zag = BZag扎格
	f.扎格Zag.SetParent(f)

}
