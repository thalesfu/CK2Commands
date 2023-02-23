package khuttal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhuttalCounty interface {
	feud.County
	BHalaward哈拉沃德() feud.Barony
	BKulob库洛布() feud.Barony
	BMunk蒙克() feud.Barony
	BQobadhiyan久越得犍() feud.Barony
	BRasht拉什特() feud.Barony
	BVakash镬沙() feud.Barony
	BWashjird瓦希尔德() feud.Barony
}

type 珂咄罗KhuttalCounty struct {
	feud.BaseCounty
	哈拉沃德Halaward   feud.Barony
	库洛布Kulob       feud.Barony
	蒙克Munk         feud.Barony
	久越得犍Qobadhiyan feud.Barony
	拉什特Rasht       feud.Barony
	镬沙Vakash       feud.Barony
	瓦希尔德Washjird   feud.Barony
}

func (c *珂咄罗KhuttalCounty) BHalaward哈拉沃德() feud.Barony {
	return c.哈拉沃德Halaward
}

func (c *珂咄罗KhuttalCounty) BKulob库洛布() feud.Barony {
	return c.库洛布Kulob
}

func (c *珂咄罗KhuttalCounty) BMunk蒙克() feud.Barony {
	return c.蒙克Munk
}

func (c *珂咄罗KhuttalCounty) BQobadhiyan久越得犍() feud.Barony {
	return c.久越得犍Qobadhiyan
}

func (c *珂咄罗KhuttalCounty) BRasht拉什特() feud.Barony {
	return c.拉什特Rasht
}

func (c *珂咄罗KhuttalCounty) BVakash镬沙() feud.Barony {
	return c.镬沙Vakash
}

func (c *珂咄罗KhuttalCounty) BWashjird瓦希尔德() feud.Barony {
	return c.瓦希尔德Washjird
}

var CKhuttal珂咄罗 KhuttalCounty = &珂咄罗KhuttalCounty{}

func init() {
	f := CKhuttal珂咄罗.(*珂咄罗KhuttalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1188",
		Title:     "khuttal",
		TitleName: "珂咄罗",
		TitleCode: "c_khuttal",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈拉沃德Halaward = BHalaward哈拉沃德
	f.哈拉沃德Halaward.SetParent(f)

	f.库洛布Kulob = BKulob库洛布
	f.库洛布Kulob.SetParent(f)

	f.蒙克Munk = BMunk蒙克
	f.蒙克Munk.SetParent(f)

	f.久越得犍Qobadhiyan = BQobadhiyan久越得犍
	f.久越得犍Qobadhiyan.SetParent(f)

	f.拉什特Rasht = BRasht拉什特
	f.拉什特Rasht.SetParent(f)

	f.镬沙Vakash = BVakash镬沙
	f.镬沙Vakash.SetParent(f)

	f.瓦希尔德Washjird = BWashjird瓦希尔德
	f.瓦希尔德Washjird.SetParent(f)

}
