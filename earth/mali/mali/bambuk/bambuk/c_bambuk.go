package bambuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BambukCounty interface {
	feud.County
	BBambuk班布克() feud.Barony
	BDiouboye迪乌博耶() feud.Barony
	BGhiyaru吉雅鲁() feud.Barony
	BGoundafal贡达法尔() feud.Barony
	BKiffa基法() feud.Barony
	BSankaran桑卡朗() feud.Barony
	BYaresna亚雷斯纳() feud.Barony
}

type 班布克BambukCounty struct {
	feud.BaseCounty
	班布克Bambuk     feud.Barony
	迪乌博耶Diouboye  feud.Barony
	吉雅鲁Ghiyaru    feud.Barony
	贡达法尔Goundafal feud.Barony
	基法Kiffa       feud.Barony
	桑卡朗Sankaran   feud.Barony
	亚雷斯纳Yaresna   feud.Barony
}

func (c *班布克BambukCounty) BBambuk班布克() feud.Barony {
	return c.班布克Bambuk
}

func (c *班布克BambukCounty) BDiouboye迪乌博耶() feud.Barony {
	return c.迪乌博耶Diouboye
}

func (c *班布克BambukCounty) BGhiyaru吉雅鲁() feud.Barony {
	return c.吉雅鲁Ghiyaru
}

func (c *班布克BambukCounty) BGoundafal贡达法尔() feud.Barony {
	return c.贡达法尔Goundafal
}

func (c *班布克BambukCounty) BKiffa基法() feud.Barony {
	return c.基法Kiffa
}

func (c *班布克BambukCounty) BSankaran桑卡朗() feud.Barony {
	return c.桑卡朗Sankaran
}

func (c *班布克BambukCounty) BYaresna亚雷斯纳() feud.Barony {
	return c.亚雷斯纳Yaresna
}

var CBambuk班布克 BambukCounty = &班布克BambukCounty{}

func init() {
	f := CBambuk班布克.(*班布克BambukCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "926",
		Title:     "bambuk",
		TitleName: "班布克",
		TitleCode: "c_bambuk",
		Baronies:  map[string]feud.Barony{},
	}

	f.班布克Bambuk = BBambuk班布克
	f.班布克Bambuk.SetParent(f)

	f.迪乌博耶Diouboye = BDiouboye迪乌博耶
	f.迪乌博耶Diouboye.SetParent(f)

	f.吉雅鲁Ghiyaru = BGhiyaru吉雅鲁
	f.吉雅鲁Ghiyaru.SetParent(f)

	f.贡达法尔Goundafal = BGoundafal贡达法尔
	f.贡达法尔Goundafal.SetParent(f)

	f.基法Kiffa = BKiffa基法
	f.基法Kiffa.SetParent(f)

	f.桑卡朗Sankaran = BSankaran桑卡朗
	f.桑卡朗Sankaran.SetParent(f)

	f.亚雷斯纳Yaresna = BYaresna亚雷斯纳
	f.亚雷斯纳Yaresna.SetParent(f)

}
