package clydesdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ClydesdaleCounty interface {
	feud.County
	BBothwell博斯韦尔() feud.Barony
	BCadzow卡奏() feud.Barony
	BDumbarton邓巴顿() feud.Barony
	BGlasgow格拉斯哥() feud.Barony
	BLanark拉纳克() feud.Barony
	BLesmahagow莱斯马黑戈() feud.Barony
	BRenfrew伦弗鲁() feud.Barony
}

type 克莱兹代尔ClydesdaleCounty struct {
	feud.BaseCounty
	博斯韦尔Bothwell    feud.Barony
	卡奏Cadzow        feud.Barony
	邓巴顿Dumbarton    feud.Barony
	格拉斯哥Glasgow     feud.Barony
	拉纳克Lanark       feud.Barony
	莱斯马黑戈Lesmahagow feud.Barony
	伦弗鲁Renfrew      feud.Barony
}

func (c *克莱兹代尔ClydesdaleCounty) BBothwell博斯韦尔() feud.Barony {
	return c.博斯韦尔Bothwell
}

func (c *克莱兹代尔ClydesdaleCounty) BCadzow卡奏() feud.Barony {
	return c.卡奏Cadzow
}

func (c *克莱兹代尔ClydesdaleCounty) BDumbarton邓巴顿() feud.Barony {
	return c.邓巴顿Dumbarton
}

func (c *克莱兹代尔ClydesdaleCounty) BGlasgow格拉斯哥() feud.Barony {
	return c.格拉斯哥Glasgow
}

func (c *克莱兹代尔ClydesdaleCounty) BLanark拉纳克() feud.Barony {
	return c.拉纳克Lanark
}

func (c *克莱兹代尔ClydesdaleCounty) BLesmahagow莱斯马黑戈() feud.Barony {
	return c.莱斯马黑戈Lesmahagow
}

func (c *克莱兹代尔ClydesdaleCounty) BRenfrew伦弗鲁() feud.Barony {
	return c.伦弗鲁Renfrew
}

var CClydesdale克莱兹代尔 ClydesdaleCounty = &克莱兹代尔ClydesdaleCounty{}

func init() {
	f := CClydesdale克莱兹代尔.(*克莱兹代尔ClydesdaleCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "47",
		Title:     "clydesdale",
		TitleName: "克莱兹代尔",
		TitleCode: "c_clydesdale",
		Baronies:  map[string]feud.Barony{},
	}

	f.博斯韦尔Bothwell = BBothwell博斯韦尔
	f.博斯韦尔Bothwell.SetParent(f)

	f.卡奏Cadzow = BCadzow卡奏
	f.卡奏Cadzow.SetParent(f)

	f.邓巴顿Dumbarton = BDumbarton邓巴顿
	f.邓巴顿Dumbarton.SetParent(f)

	f.格拉斯哥Glasgow = BGlasgow格拉斯哥
	f.格拉斯哥Glasgow.SetParent(f)

	f.拉纳克Lanark = BLanark拉纳克
	f.拉纳克Lanark.SetParent(f)

	f.莱斯马黑戈Lesmahagow = BLesmahagow莱斯马黑戈
	f.莱斯马黑戈Lesmahagow.SetParent(f)

	f.伦弗鲁Renfrew = BRenfrew伦弗鲁
	f.伦弗鲁Renfrew.SetParent(f)

}
