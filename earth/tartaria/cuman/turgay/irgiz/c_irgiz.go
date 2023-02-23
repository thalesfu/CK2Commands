package irgiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IrgizCounty interface {
	feud.County
	BAlabas阿拉巴斯() feud.Barony
	BBerchogur别尔乔古尔() feud.Barony
	BIrgiz伊尔吉兹() feud.Barony
	BKair凯尔() feud.Barony
	BShaktha沙赫塔() feud.Barony
	BShalkar沙尔卡尔() feud.Barony
	BUlpan乌尔潘() feud.Barony
}

type 伊尔吉兹IrgizCounty struct {
	feud.BaseCounty
	阿拉巴斯Alabas     feud.Barony
	别尔乔古尔Berchogur feud.Barony
	伊尔吉兹Irgiz      feud.Barony
	凯尔Kair         feud.Barony
	沙赫塔Shaktha     feud.Barony
	沙尔卡尔Shalkar    feud.Barony
	乌尔潘Ulpan       feud.Barony
}

func (c *伊尔吉兹IrgizCounty) BAlabas阿拉巴斯() feud.Barony {
	return c.阿拉巴斯Alabas
}

func (c *伊尔吉兹IrgizCounty) BBerchogur别尔乔古尔() feud.Barony {
	return c.别尔乔古尔Berchogur
}

func (c *伊尔吉兹IrgizCounty) BIrgiz伊尔吉兹() feud.Barony {
	return c.伊尔吉兹Irgiz
}

func (c *伊尔吉兹IrgizCounty) BKair凯尔() feud.Barony {
	return c.凯尔Kair
}

func (c *伊尔吉兹IrgizCounty) BShaktha沙赫塔() feud.Barony {
	return c.沙赫塔Shaktha
}

func (c *伊尔吉兹IrgizCounty) BShalkar沙尔卡尔() feud.Barony {
	return c.沙尔卡尔Shalkar
}

func (c *伊尔吉兹IrgizCounty) BUlpan乌尔潘() feud.Barony {
	return c.乌尔潘Ulpan
}

var CIrgiz伊尔吉兹 IrgizCounty = &伊尔吉兹IrgizCounty{}

func init() {
	f := CIrgiz伊尔吉兹.(*伊尔吉兹IrgizCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1869",
		Title:     "irgiz",
		TitleName: "伊尔吉兹",
		TitleCode: "c_irgiz",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉巴斯Alabas = BAlabas阿拉巴斯
	f.阿拉巴斯Alabas.SetParent(f)

	f.别尔乔古尔Berchogur = BBerchogur别尔乔古尔
	f.别尔乔古尔Berchogur.SetParent(f)

	f.伊尔吉兹Irgiz = BIrgiz伊尔吉兹
	f.伊尔吉兹Irgiz.SetParent(f)

	f.凯尔Kair = BKair凯尔
	f.凯尔Kair.SetParent(f)

	f.沙赫塔Shaktha = BShaktha沙赫塔
	f.沙赫塔Shaktha.SetParent(f)

	f.沙尔卡尔Shalkar = BShalkar沙尔卡尔
	f.沙尔卡尔Shalkar.SetParent(f)

	f.乌尔潘Ulpan = BUlpan乌尔潘
	f.乌尔潘Ulpan.SetParent(f)

}
