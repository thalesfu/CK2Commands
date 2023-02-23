package kerman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KermanCounty interface {
	feud.County
	BBardsir巴尔德斯尔() feud.Barony
	BBehdesir比哈德斯尔() feud.Barony
	BGhaleganj吉哈勒甘吉() feud.Barony
	BKerman克尔曼() feud.Barony
	BMaymand迈曼德() feud.Barony
	BNough努格() feud.Barony
	BRayen拉延() feud.Barony
	BZarand扎兰德() feud.Barony
}

type 克尔曼KermanCounty struct {
	feud.BaseCounty
	巴尔德斯尔Bardsir   feud.Barony
	比哈德斯尔Behdesir  feud.Barony
	吉哈勒甘吉Ghaleganj feud.Barony
	克尔曼Kerman      feud.Barony
	迈曼德Maymand     feud.Barony
	努格Nough        feud.Barony
	拉延Rayen        feud.Barony
	扎兰德Zarand      feud.Barony
}

func (c *克尔曼KermanCounty) BBardsir巴尔德斯尔() feud.Barony {
	return c.巴尔德斯尔Bardsir
}

func (c *克尔曼KermanCounty) BBehdesir比哈德斯尔() feud.Barony {
	return c.比哈德斯尔Behdesir
}

func (c *克尔曼KermanCounty) BGhaleganj吉哈勒甘吉() feud.Barony {
	return c.吉哈勒甘吉Ghaleganj
}

func (c *克尔曼KermanCounty) BKerman克尔曼() feud.Barony {
	return c.克尔曼Kerman
}

func (c *克尔曼KermanCounty) BMaymand迈曼德() feud.Barony {
	return c.迈曼德Maymand
}

func (c *克尔曼KermanCounty) BNough努格() feud.Barony {
	return c.努格Nough
}

func (c *克尔曼KermanCounty) BRayen拉延() feud.Barony {
	return c.拉延Rayen
}

func (c *克尔曼KermanCounty) BZarand扎兰德() feud.Barony {
	return c.扎兰德Zarand
}

var CKerman克尔曼 KermanCounty = &克尔曼KermanCounty{}

func init() {
	f := CKerman克尔曼.(*克尔曼KermanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "639",
		Title:     "kerman",
		TitleName: "克尔曼",
		TitleCode: "c_kerman",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔德斯尔Bardsir = BBardsir巴尔德斯尔
	f.巴尔德斯尔Bardsir.SetParent(f)

	f.比哈德斯尔Behdesir = BBehdesir比哈德斯尔
	f.比哈德斯尔Behdesir.SetParent(f)

	f.吉哈勒甘吉Ghaleganj = BGhaleganj吉哈勒甘吉
	f.吉哈勒甘吉Ghaleganj.SetParent(f)

	f.克尔曼Kerman = BKerman克尔曼
	f.克尔曼Kerman.SetParent(f)

	f.迈曼德Maymand = BMaymand迈曼德
	f.迈曼德Maymand.SetParent(f)

	f.努格Nough = BNough努格
	f.努格Nough.SetParent(f)

	f.拉延Rayen = BRayen拉延
	f.拉延Rayen.SetParent(f)

	f.扎兰德Zarand = BZarand扎兰德
	f.扎兰德Zarand.SetParent(f)

}
