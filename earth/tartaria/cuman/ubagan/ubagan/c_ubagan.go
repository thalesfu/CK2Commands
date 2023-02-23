package ubagan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UbaganCounty interface {
	feud.County
	BAuliekol奥利耶科利() feud.Barony
	BChernigovka切尔尼戈夫卡() feud.Barony
	BKarasu卡拉苏() feud.Barony
	BKozubay科祖拜() feud.Barony
	BTyuntyugur琼秋古尔() feud.Barony
	BUbagan乌巴甘() feud.Barony
	BZhumagul朱马古利() feud.Barony
}

type 乌巴甘UbaganCounty struct {
	feud.BaseCounty
	奥利耶科利Auliekol     feud.Barony
	切尔尼戈夫卡Chernigovka feud.Barony
	卡拉苏Karasu         feud.Barony
	科祖拜Kozubay        feud.Barony
	琼秋古尔Tyuntyugur    feud.Barony
	乌巴甘Ubagan         feud.Barony
	朱马古利Zhumagul      feud.Barony
}

func (c *乌巴甘UbaganCounty) BAuliekol奥利耶科利() feud.Barony {
	return c.奥利耶科利Auliekol
}

func (c *乌巴甘UbaganCounty) BChernigovka切尔尼戈夫卡() feud.Barony {
	return c.切尔尼戈夫卡Chernigovka
}

func (c *乌巴甘UbaganCounty) BKarasu卡拉苏() feud.Barony {
	return c.卡拉苏Karasu
}

func (c *乌巴甘UbaganCounty) BKozubay科祖拜() feud.Barony {
	return c.科祖拜Kozubay
}

func (c *乌巴甘UbaganCounty) BTyuntyugur琼秋古尔() feud.Barony {
	return c.琼秋古尔Tyuntyugur
}

func (c *乌巴甘UbaganCounty) BUbagan乌巴甘() feud.Barony {
	return c.乌巴甘Ubagan
}

func (c *乌巴甘UbaganCounty) BZhumagul朱马古利() feud.Barony {
	return c.朱马古利Zhumagul
}

var CUbagan乌巴甘 UbaganCounty = &乌巴甘UbaganCounty{}

func init() {
	f := CUbagan乌巴甘.(*乌巴甘UbaganCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1856",
		Title:     "ubagan",
		TitleName: "乌巴甘",
		TitleCode: "c_ubagan",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥利耶科利Auliekol = BAuliekol奥利耶科利
	f.奥利耶科利Auliekol.SetParent(f)

	f.切尔尼戈夫卡Chernigovka = BChernigovka切尔尼戈夫卡
	f.切尔尼戈夫卡Chernigovka.SetParent(f)

	f.卡拉苏Karasu = BKarasu卡拉苏
	f.卡拉苏Karasu.SetParent(f)

	f.科祖拜Kozubay = BKozubay科祖拜
	f.科祖拜Kozubay.SetParent(f)

	f.琼秋古尔Tyuntyugur = BTyuntyugur琼秋古尔
	f.琼秋古尔Tyuntyugur.SetParent(f)

	f.乌巴甘Ubagan = BUbagan乌巴甘
	f.乌巴甘Ubagan.SetParent(f)

	f.朱马古利Zhumagul = BZhumagul朱马古利
	f.朱马古利Zhumagul.SetParent(f)

}
