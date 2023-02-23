package gaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GayaCounty interface {
	feud.County
	BBahloipur婆诃卢罗补罗() feud.Barony
	BElahabad埃拉哈巴德() feud.Barony
	BGaya伽耶() feud.Barony
	BNalanda那烂陀() feud.Barony
	BNawada那婆陀() feud.Barony
	BSherghati舍夷罗迦提() feud.Barony
}

type 伽耶GayaCounty struct {
	feud.BaseCounty
	婆诃卢罗补罗Bahloipur feud.Barony
	埃拉哈巴德Elahabad   feud.Barony
	伽耶Gaya          feud.Barony
	那烂陀Nalanda      feud.Barony
	那婆陀Nawada       feud.Barony
	舍夷罗迦提Sherghati  feud.Barony
}

func (c *伽耶GayaCounty) BBahloipur婆诃卢罗补罗() feud.Barony {
	return c.婆诃卢罗补罗Bahloipur
}

func (c *伽耶GayaCounty) BElahabad埃拉哈巴德() feud.Barony {
	return c.埃拉哈巴德Elahabad
}

func (c *伽耶GayaCounty) BGaya伽耶() feud.Barony {
	return c.伽耶Gaya
}

func (c *伽耶GayaCounty) BNalanda那烂陀() feud.Barony {
	return c.那烂陀Nalanda
}

func (c *伽耶GayaCounty) BNawada那婆陀() feud.Barony {
	return c.那婆陀Nawada
}

func (c *伽耶GayaCounty) BSherghati舍夷罗迦提() feud.Barony {
	return c.舍夷罗迦提Sherghati
}

var CGaya伽耶 GayaCounty = &伽耶GayaCounty{}

func init() {
	f := CGaya伽耶.(*伽耶GayaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1276",
		Title:     "gaya",
		TitleName: "伽耶",
		TitleCode: "c_gaya",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆诃卢罗补罗Bahloipur = BBahloipur婆诃卢罗补罗
	f.婆诃卢罗补罗Bahloipur.SetParent(f)

	f.埃拉哈巴德Elahabad = BElahabad埃拉哈巴德
	f.埃拉哈巴德Elahabad.SetParent(f)

	f.伽耶Gaya = BGaya伽耶
	f.伽耶Gaya.SetParent(f)

	f.那烂陀Nalanda = BNalanda那烂陀
	f.那烂陀Nalanda.SetParent(f)

	f.那婆陀Nawada = BNawada那婆陀
	f.那婆陀Nawada.SetParent(f)

	f.舍夷罗迦提Sherghati = BSherghati舍夷罗迦提
	f.舍夷罗迦提Sherghati.SetParent(f)

}
