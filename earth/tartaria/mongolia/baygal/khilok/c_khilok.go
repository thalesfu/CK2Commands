package khilok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhilokCounty interface {
	feud.County
	BBilyutay比柳泰() feud.Barony
	BKareliya卡列利亚() feud.Barony
	BKhilok勤勒豁() feud.Barony
	BSharalday沙喇勒岱() feud.Barony
	BStolga斯托尔加() feud.Barony
	BZakaznik扎卡兹尼克() feud.Barony
}

type 勤勒豁KhilokCounty struct {
	feud.BaseCounty
	比柳泰Bilyutay   feud.Barony
	卡列利亚Kareliya  feud.Barony
	勤勒豁Khilok     feud.Barony
	沙喇勒岱Sharalday feud.Barony
	斯托尔加Stolga    feud.Barony
	扎卡兹尼克Zakaznik feud.Barony
}

func (c *勤勒豁KhilokCounty) BBilyutay比柳泰() feud.Barony {
	return c.比柳泰Bilyutay
}

func (c *勤勒豁KhilokCounty) BKareliya卡列利亚() feud.Barony {
	return c.卡列利亚Kareliya
}

func (c *勤勒豁KhilokCounty) BKhilok勤勒豁() feud.Barony {
	return c.勤勒豁Khilok
}

func (c *勤勒豁KhilokCounty) BSharalday沙喇勒岱() feud.Barony {
	return c.沙喇勒岱Sharalday
}

func (c *勤勒豁KhilokCounty) BStolga斯托尔加() feud.Barony {
	return c.斯托尔加Stolga
}

func (c *勤勒豁KhilokCounty) BZakaznik扎卡兹尼克() feud.Barony {
	return c.扎卡兹尼克Zakaznik
}

var CKhilok勤勒豁 KhilokCounty = &勤勒豁KhilokCounty{}

func init() {
	f := CKhilok勤勒豁.(*勤勒豁KhilokCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1923",
		Title:     "khilok",
		TitleName: "勤勒豁",
		TitleCode: "c_khilok",
		Baronies:  map[string]feud.Barony{},
	}

	f.比柳泰Bilyutay = BBilyutay比柳泰
	f.比柳泰Bilyutay.SetParent(f)

	f.卡列利亚Kareliya = BKareliya卡列利亚
	f.卡列利亚Kareliya.SetParent(f)

	f.勤勒豁Khilok = BKhilok勤勒豁
	f.勤勒豁Khilok.SetParent(f)

	f.沙喇勒岱Sharalday = BSharalday沙喇勒岱
	f.沙喇勒岱Sharalday.SetParent(f)

	f.斯托尔加Stolga = BStolga斯托尔加
	f.斯托尔加Stolga.SetParent(f)

	f.扎卡兹尼克Zakaznik = BZakaznik扎卡兹尼克
	f.扎卡兹尼克Zakaznik.SetParent(f)

}
