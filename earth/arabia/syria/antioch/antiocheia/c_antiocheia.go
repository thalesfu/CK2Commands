package antiocheia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AntiocheiaCounty interface {
	feud.County
	BAntiocheia安条克() feud.Barony
	BBaghras巴格拉斯() feud.Barony
	BDarbsak达巴萨克() feud.Barony
	BHarenc哈里姆() feud.Barony
	BHazart哈扎特() feud.Barony
	BLatakiah拉塔基亚() feud.Barony
	BSaone索恩() feud.Barony
	BStsymeon圣西梅昂() feud.Barony
}

type 安条克AntiocheiaCounty struct {
	feud.BaseCounty
	安条克Antiocheia feud.Barony
	巴格拉斯Baghras   feud.Barony
	达巴萨克Darbsak   feud.Barony
	哈里姆Harenc     feud.Barony
	哈扎特Hazart     feud.Barony
	拉塔基亚Latakiah  feud.Barony
	索恩Saone       feud.Barony
	圣西梅昂Stsymeon  feud.Barony
}

func (c *安条克AntiocheiaCounty) BAntiocheia安条克() feud.Barony {
	return c.安条克Antiocheia
}

func (c *安条克AntiocheiaCounty) BBaghras巴格拉斯() feud.Barony {
	return c.巴格拉斯Baghras
}

func (c *安条克AntiocheiaCounty) BDarbsak达巴萨克() feud.Barony {
	return c.达巴萨克Darbsak
}

func (c *安条克AntiocheiaCounty) BHarenc哈里姆() feud.Barony {
	return c.哈里姆Harenc
}

func (c *安条克AntiocheiaCounty) BHazart哈扎特() feud.Barony {
	return c.哈扎特Hazart
}

func (c *安条克AntiocheiaCounty) BLatakiah拉塔基亚() feud.Barony {
	return c.拉塔基亚Latakiah
}

func (c *安条克AntiocheiaCounty) BSaone索恩() feud.Barony {
	return c.索恩Saone
}

func (c *安条克AntiocheiaCounty) BStsymeon圣西梅昂() feud.Barony {
	return c.圣西梅昂Stsymeon
}

var CAntiocheia安条克 AntiocheiaCounty = &安条克AntiocheiaCounty{}

func init() {
	f := CAntiocheia安条克.(*安条克AntiocheiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "764",
		Title:     "antiocheia",
		TitleName: "安条克",
		TitleCode: "c_antiocheia",
		Baronies:  map[string]feud.Barony{},
	}

	f.安条克Antiocheia = BAntiocheia安条克
	f.安条克Antiocheia.SetParent(f)

	f.巴格拉斯Baghras = BBaghras巴格拉斯
	f.巴格拉斯Baghras.SetParent(f)

	f.达巴萨克Darbsak = BDarbsak达巴萨克
	f.达巴萨克Darbsak.SetParent(f)

	f.哈里姆Harenc = BHarenc哈里姆
	f.哈里姆Harenc.SetParent(f)

	f.哈扎特Hazart = BHazart哈扎特
	f.哈扎特Hazart.SetParent(f)

	f.拉塔基亚Latakiah = BLatakiah拉塔基亚
	f.拉塔基亚Latakiah.SetParent(f)

	f.索恩Saone = BSaone索恩
	f.索恩Saone.SetParent(f)

	f.圣西梅昂Stsymeon = BStsymeon圣西梅昂
	f.圣西梅昂Stsymeon.SetParent(f)

}
