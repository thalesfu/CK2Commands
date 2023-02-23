package krain

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KrainCounty interface {
	feud.County
	BAuersperg奥埃尔斯佩格() feud.Barony
	BGorz格尔茨() feud.Barony
	BGuetenegg古埃特内格() feud.Barony
	BGurk古尔克() feud.Barony
	BKrainburg克赖因堡() feud.Barony
	BStain施泰因() feud.Barony
	BStveit圣法伊特() feud.Barony
	BZerknitz齐尔克尼茨() feud.Barony
}

type 克雷纳KrainCounty struct {
	feud.BaseCounty
	奥埃尔斯佩格Auersperg feud.Barony
	格尔茨Gorz         feud.Barony
	古埃特内格Guetenegg  feud.Barony
	古尔克Gurk         feud.Barony
	克赖因堡Krainburg   feud.Barony
	施泰因Stain        feud.Barony
	圣法伊特Stveit      feud.Barony
	齐尔克尼茨Zerknitz   feud.Barony
}

func (c *克雷纳KrainCounty) BAuersperg奥埃尔斯佩格() feud.Barony {
	return c.奥埃尔斯佩格Auersperg
}

func (c *克雷纳KrainCounty) BGorz格尔茨() feud.Barony {
	return c.格尔茨Gorz
}

func (c *克雷纳KrainCounty) BGuetenegg古埃特内格() feud.Barony {
	return c.古埃特内格Guetenegg
}

func (c *克雷纳KrainCounty) BGurk古尔克() feud.Barony {
	return c.古尔克Gurk
}

func (c *克雷纳KrainCounty) BKrainburg克赖因堡() feud.Barony {
	return c.克赖因堡Krainburg
}

func (c *克雷纳KrainCounty) BStain施泰因() feud.Barony {
	return c.施泰因Stain
}

func (c *克雷纳KrainCounty) BStveit圣法伊特() feud.Barony {
	return c.圣法伊特Stveit
}

func (c *克雷纳KrainCounty) BZerknitz齐尔克尼茨() feud.Barony {
	return c.齐尔克尼茨Zerknitz
}

var CKrain克雷纳 KrainCounty = &克雷纳KrainCounty{}

func init() {
	f := CKrain克雷纳.(*克雷纳KrainCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "457",
		Title:     "krain",
		TitleName: "克雷纳",
		TitleCode: "c_krain",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥埃尔斯佩格Auersperg = BAuersperg奥埃尔斯佩格
	f.奥埃尔斯佩格Auersperg.SetParent(f)

	f.格尔茨Gorz = BGorz格尔茨
	f.格尔茨Gorz.SetParent(f)

	f.古埃特内格Guetenegg = BGuetenegg古埃特内格
	f.古埃特内格Guetenegg.SetParent(f)

	f.古尔克Gurk = BGurk古尔克
	f.古尔克Gurk.SetParent(f)

	f.克赖因堡Krainburg = BKrainburg克赖因堡
	f.克赖因堡Krainburg.SetParent(f)

	f.施泰因Stain = BStain施泰因
	f.施泰因Stain.SetParent(f)

	f.圣法伊特Stveit = BStveit圣法伊特
	f.圣法伊特Stveit.SetParent(f)

	f.齐尔克尼茨Zerknitz = BZerknitz齐尔克尼茨
	f.齐尔克尼茨Zerknitz.SetParent(f)

}
