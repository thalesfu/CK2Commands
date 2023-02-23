package kashin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KashinCounty interface {
	feud.County
	BGoritsy戈里齐() feud.Barony
	BKaliazin卡利亚津() feud.Barony
	BKashin卡申() feud.Barony
	BKimry基姆雷() feud.Barony
	BLikhoslavl利霍斯拉夫尔() feud.Barony
	BUstinovo乌斯季诺沃() feud.Barony
}

type 卡申KashinCounty struct {
	feud.BaseCounty
	戈里齐Goritsy       feud.Barony
	卡利亚津Kaliazin     feud.Barony
	卡申Kashin         feud.Barony
	基姆雷Kimry         feud.Barony
	利霍斯拉夫尔Likhoslavl feud.Barony
	乌斯季诺沃Ustinovo    feud.Barony
}

func (c *卡申KashinCounty) BGoritsy戈里齐() feud.Barony {
	return c.戈里齐Goritsy
}

func (c *卡申KashinCounty) BKaliazin卡利亚津() feud.Barony {
	return c.卡利亚津Kaliazin
}

func (c *卡申KashinCounty) BKashin卡申() feud.Barony {
	return c.卡申Kashin
}

func (c *卡申KashinCounty) BKimry基姆雷() feud.Barony {
	return c.基姆雷Kimry
}

func (c *卡申KashinCounty) BLikhoslavl利霍斯拉夫尔() feud.Barony {
	return c.利霍斯拉夫尔Likhoslavl
}

func (c *卡申KashinCounty) BUstinovo乌斯季诺沃() feud.Barony {
	return c.乌斯季诺沃Ustinovo
}

var CKashin卡申 KashinCounty = &卡申KashinCounty{}

func init() {
	f := CKashin卡申.(*卡申KashinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1669",
		Title:     "kashin",
		TitleName: "卡申",
		TitleCode: "c_kashin",
		Baronies:  map[string]feud.Barony{},
	}

	f.戈里齐Goritsy = BGoritsy戈里齐
	f.戈里齐Goritsy.SetParent(f)

	f.卡利亚津Kaliazin = BKaliazin卡利亚津
	f.卡利亚津Kaliazin.SetParent(f)

	f.卡申Kashin = BKashin卡申
	f.卡申Kashin.SetParent(f)

	f.基姆雷Kimry = BKimry基姆雷
	f.基姆雷Kimry.SetParent(f)

	f.利霍斯拉夫尔Likhoslavl = BLikhoslavl利霍斯拉夫尔
	f.利霍斯拉夫尔Likhoslavl.SetParent(f)

	f.乌斯季诺沃Ustinovo = BUstinovo乌斯季诺沃
	f.乌斯季诺沃Ustinovo.SetParent(f)

}
