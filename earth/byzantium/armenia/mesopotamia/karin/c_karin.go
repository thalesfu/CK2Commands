package karin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarinCounty interface {
	feud.County
	BBaghesh比特利斯() feud.Barony
	BDaroynk达洛克() feud.Barony
	BKarin卡林() feud.Barony
	BKirklarkilisesi基克拉克吉利瑟斯() feud.Barony
	BMardin马尔丁() feud.Barony
	BOwshank奥乌斯汉克() feud.Barony
	BPasen帕森() feud.Barony
	BTercan泰尔詹() feud.Barony
}

type 卡林KarinCounty struct {
	feud.BaseCounty
	比特利斯Baghesh             feud.Barony
	达洛克Daroynk              feud.Barony
	卡林Karin                 feud.Barony
	基克拉克吉利瑟斯Kirklarkilisesi feud.Barony
	马尔丁Mardin               feud.Barony
	奥乌斯汉克Owshank            feud.Barony
	帕森Pasen                 feud.Barony
	泰尔詹Tercan               feud.Barony
}

func (c *卡林KarinCounty) BBaghesh比特利斯() feud.Barony {
	return c.比特利斯Baghesh
}

func (c *卡林KarinCounty) BDaroynk达洛克() feud.Barony {
	return c.达洛克Daroynk
}

func (c *卡林KarinCounty) BKarin卡林() feud.Barony {
	return c.卡林Karin
}

func (c *卡林KarinCounty) BKirklarkilisesi基克拉克吉利瑟斯() feud.Barony {
	return c.基克拉克吉利瑟斯Kirklarkilisesi
}

func (c *卡林KarinCounty) BMardin马尔丁() feud.Barony {
	return c.马尔丁Mardin
}

func (c *卡林KarinCounty) BOwshank奥乌斯汉克() feud.Barony {
	return c.奥乌斯汉克Owshank
}

func (c *卡林KarinCounty) BPasen帕森() feud.Barony {
	return c.帕森Pasen
}

func (c *卡林KarinCounty) BTercan泰尔詹() feud.Barony {
	return c.泰尔詹Tercan
}

var CKarin卡林 KarinCounty = &卡林KarinCounty{}

func init() {
	f := CKarin卡林.(*卡林KarinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "703",
		Title:     "karin",
		TitleName: "卡林",
		TitleCode: "c_karin",
		Baronies:  map[string]feud.Barony{},
	}

	f.比特利斯Baghesh = BBaghesh比特利斯
	f.比特利斯Baghesh.SetParent(f)

	f.达洛克Daroynk = BDaroynk达洛克
	f.达洛克Daroynk.SetParent(f)

	f.卡林Karin = BKarin卡林
	f.卡林Karin.SetParent(f)

	f.基克拉克吉利瑟斯Kirklarkilisesi = BKirklarkilisesi基克拉克吉利瑟斯
	f.基克拉克吉利瑟斯Kirklarkilisesi.SetParent(f)

	f.马尔丁Mardin = BMardin马尔丁
	f.马尔丁Mardin.SetParent(f)

	f.奥乌斯汉克Owshank = BOwshank奥乌斯汉克
	f.奥乌斯汉克Owshank.SetParent(f)

	f.帕森Pasen = BPasen帕森
	f.帕森Pasen.SetParent(f)

	f.泰尔詹Tercan = BTercan泰尔詹
	f.泰尔詹Tercan.SetParent(f)

}
