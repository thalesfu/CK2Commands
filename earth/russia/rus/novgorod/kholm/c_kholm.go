package kholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KholmCounty interface {
	feud.County
	BFishovo菲绍沃() feud.Barony
	BKholm霍尔姆() feud.Barony
	BLychkovo雷奇科沃() feud.Barony
	BMolvotitsy莫尔沃季齐() feud.Barony
	BNevskaya涅夫斯卡亚() feud.Barony
	BSuchki苏奇基() feud.Barony
	BYazhelbitsy亚热尔比齐() feud.Barony
}

type 霍尔姆KholmCounty struct {
	feud.BaseCounty
	菲绍沃Fishovo       feud.Barony
	霍尔姆Kholm         feud.Barony
	雷奇科沃Lychkovo     feud.Barony
	莫尔沃季齐Molvotitsy  feud.Barony
	涅夫斯卡亚Nevskaya    feud.Barony
	苏奇基Suchki        feud.Barony
	亚热尔比齐Yazhelbitsy feud.Barony
}

func (c *霍尔姆KholmCounty) BFishovo菲绍沃() feud.Barony {
	return c.菲绍沃Fishovo
}

func (c *霍尔姆KholmCounty) BKholm霍尔姆() feud.Barony {
	return c.霍尔姆Kholm
}

func (c *霍尔姆KholmCounty) BLychkovo雷奇科沃() feud.Barony {
	return c.雷奇科沃Lychkovo
}

func (c *霍尔姆KholmCounty) BMolvotitsy莫尔沃季齐() feud.Barony {
	return c.莫尔沃季齐Molvotitsy
}

func (c *霍尔姆KholmCounty) BNevskaya涅夫斯卡亚() feud.Barony {
	return c.涅夫斯卡亚Nevskaya
}

func (c *霍尔姆KholmCounty) BSuchki苏奇基() feud.Barony {
	return c.苏奇基Suchki
}

func (c *霍尔姆KholmCounty) BYazhelbitsy亚热尔比齐() feud.Barony {
	return c.亚热尔比齐Yazhelbitsy
}

var CKholm霍尔姆 KholmCounty = &霍尔姆KholmCounty{}

func init() {
	f := CKholm霍尔姆.(*霍尔姆KholmCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1667",
		Title:     "kholm",
		TitleName: "霍尔姆",
		TitleCode: "c_kholm",
		Baronies:  map[string]feud.Barony{},
	}

	f.菲绍沃Fishovo = BFishovo菲绍沃
	f.菲绍沃Fishovo.SetParent(f)

	f.霍尔姆Kholm = BKholm霍尔姆
	f.霍尔姆Kholm.SetParent(f)

	f.雷奇科沃Lychkovo = BLychkovo雷奇科沃
	f.雷奇科沃Lychkovo.SetParent(f)

	f.莫尔沃季齐Molvotitsy = BMolvotitsy莫尔沃季齐
	f.莫尔沃季齐Molvotitsy.SetParent(f)

	f.涅夫斯卡亚Nevskaya = BNevskaya涅夫斯卡亚
	f.涅夫斯卡亚Nevskaya.SetParent(f)

	f.苏奇基Suchki = BSuchki苏奇基
	f.苏奇基Suchki.SetParent(f)

	f.亚热尔比齐Yazhelbitsy = BYazhelbitsy亚热尔比齐
	f.亚热尔比齐Yazhelbitsy.SetParent(f)

}
