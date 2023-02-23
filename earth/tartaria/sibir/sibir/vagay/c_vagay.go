package vagay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VagayCounty interface {
	feud.County
	BBegitino别吉季诺() feud.Barony
	BChernoye切尔诺耶() feud.Barony
	BKularovo库拉罗沃() feud.Barony
	BTakhtagul塔赫塔古尔() feud.Barony
	BTukuz图库兹() feud.Barony
	BVagay瓦盖() feud.Barony
}

type 瓦盖VagayCounty struct {
	feud.BaseCounty
	别吉季诺Begitino   feud.Barony
	切尔诺耶Chernoye   feud.Barony
	库拉罗沃Kularovo   feud.Barony
	塔赫塔古尔Takhtagul feud.Barony
	图库兹Tukuz       feud.Barony
	瓦盖Vagay        feud.Barony
}

func (c *瓦盖VagayCounty) BBegitino别吉季诺() feud.Barony {
	return c.别吉季诺Begitino
}

func (c *瓦盖VagayCounty) BChernoye切尔诺耶() feud.Barony {
	return c.切尔诺耶Chernoye
}

func (c *瓦盖VagayCounty) BKularovo库拉罗沃() feud.Barony {
	return c.库拉罗沃Kularovo
}

func (c *瓦盖VagayCounty) BTakhtagul塔赫塔古尔() feud.Barony {
	return c.塔赫塔古尔Takhtagul
}

func (c *瓦盖VagayCounty) BTukuz图库兹() feud.Barony {
	return c.图库兹Tukuz
}

func (c *瓦盖VagayCounty) BVagay瓦盖() feud.Barony {
	return c.瓦盖Vagay
}

var CVagay瓦盖 VagayCounty = &瓦盖VagayCounty{}

func init() {
	f := CVagay瓦盖.(*瓦盖VagayCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1863",
		Title:     "vagay",
		TitleName: "瓦盖",
		TitleCode: "c_vagay",
		Baronies:  map[string]feud.Barony{},
	}

	f.别吉季诺Begitino = BBegitino别吉季诺
	f.别吉季诺Begitino.SetParent(f)

	f.切尔诺耶Chernoye = BChernoye切尔诺耶
	f.切尔诺耶Chernoye.SetParent(f)

	f.库拉罗沃Kularovo = BKularovo库拉罗沃
	f.库拉罗沃Kularovo.SetParent(f)

	f.塔赫塔古尔Takhtagul = BTakhtagul塔赫塔古尔
	f.塔赫塔古尔Takhtagul.SetParent(f)

	f.图库兹Tukuz = BTukuz图库兹
	f.图库兹Tukuz.SetParent(f)

	f.瓦盖Vagay = BVagay瓦盖
	f.瓦盖Vagay.SetParent(f)

}
