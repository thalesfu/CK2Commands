package tyrnovo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TyrnovoCounty interface {
	feud.County
	BChirpan奇尔潘() feud.Barony
	BHisarya希萨里亚() feud.Barony
	BIrinopolis伊里诺波利斯() feud.Barony
	BKazanlak卡赞勒克() feud.Barony
	BKilifarevski基利法雷夫斯基() feud.Barony
	BMaglizh默格利日() feud.Barony
	BOpan奥潘() feud.Barony
	BTyrnovo特尔诺沃() feud.Barony
}

type 特尔诺沃TyrnovoCounty struct {
	feud.BaseCounty
	奇尔潘Chirpan          feud.Barony
	希萨里亚Hisarya         feud.Barony
	伊里诺波利斯Irinopolis    feud.Barony
	卡赞勒克Kazanlak        feud.Barony
	基利法雷夫斯基Kilifarevski feud.Barony
	默格利日Maglizh         feud.Barony
	奥潘Opan              feud.Barony
	特尔诺沃Tyrnovo         feud.Barony
}

func (c *特尔诺沃TyrnovoCounty) BChirpan奇尔潘() feud.Barony {
	return c.奇尔潘Chirpan
}

func (c *特尔诺沃TyrnovoCounty) BHisarya希萨里亚() feud.Barony {
	return c.希萨里亚Hisarya
}

func (c *特尔诺沃TyrnovoCounty) BIrinopolis伊里诺波利斯() feud.Barony {
	return c.伊里诺波利斯Irinopolis
}

func (c *特尔诺沃TyrnovoCounty) BKazanlak卡赞勒克() feud.Barony {
	return c.卡赞勒克Kazanlak
}

func (c *特尔诺沃TyrnovoCounty) BKilifarevski基利法雷夫斯基() feud.Barony {
	return c.基利法雷夫斯基Kilifarevski
}

func (c *特尔诺沃TyrnovoCounty) BMaglizh默格利日() feud.Barony {
	return c.默格利日Maglizh
}

func (c *特尔诺沃TyrnovoCounty) BOpan奥潘() feud.Barony {
	return c.奥潘Opan
}

func (c *特尔诺沃TyrnovoCounty) BTyrnovo特尔诺沃() feud.Barony {
	return c.特尔诺沃Tyrnovo
}

var CTyrnovo特尔诺沃 TyrnovoCounty = &特尔诺沃TyrnovoCounty{}

func init() {
	f := CTyrnovo特尔诺沃.(*特尔诺沃TyrnovoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "499",
		Title:     "tyrnovo",
		TitleName: "特尔诺沃",
		TitleCode: "c_tyrnovo",
		Baronies:  map[string]feud.Barony{},
	}

	f.奇尔潘Chirpan = BChirpan奇尔潘
	f.奇尔潘Chirpan.SetParent(f)

	f.希萨里亚Hisarya = BHisarya希萨里亚
	f.希萨里亚Hisarya.SetParent(f)

	f.伊里诺波利斯Irinopolis = BIrinopolis伊里诺波利斯
	f.伊里诺波利斯Irinopolis.SetParent(f)

	f.卡赞勒克Kazanlak = BKazanlak卡赞勒克
	f.卡赞勒克Kazanlak.SetParent(f)

	f.基利法雷夫斯基Kilifarevski = BKilifarevski基利法雷夫斯基
	f.基利法雷夫斯基Kilifarevski.SetParent(f)

	f.默格利日Maglizh = BMaglizh默格利日
	f.默格利日Maglizh.SetParent(f)

	f.奥潘Opan = BOpan奥潘
	f.奥潘Opan.SetParent(f)

	f.特尔诺沃Tyrnovo = BTyrnovo特尔诺沃
	f.特尔诺沃Tyrnovo.SetParent(f)

}
