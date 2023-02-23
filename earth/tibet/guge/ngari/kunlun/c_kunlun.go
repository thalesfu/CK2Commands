package kunlun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KunlunCounty interface {
	feud.County
	BAshi阿其克() feud.Barony
	BBangdag邦达() feud.Barony
	BDulihri独立石() feud.Barony
	BPur普尔() feud.Barony
	BSeldang色当() feud.Barony
	BUlug乌鲁克() feud.Barony
	BYulya月牙() feud.Barony
}

type 昆仑KunlunCounty struct {
	feud.BaseCounty
	阿其克Ashi    feud.Barony
	邦达Bangdag  feud.Barony
	独立石Dulihri feud.Barony
	普尔Pur      feud.Barony
	色当Seldang  feud.Barony
	乌鲁克Ulug    feud.Barony
	月牙Yulya    feud.Barony
}

func (c *昆仑KunlunCounty) BAshi阿其克() feud.Barony {
	return c.阿其克Ashi
}

func (c *昆仑KunlunCounty) BBangdag邦达() feud.Barony {
	return c.邦达Bangdag
}

func (c *昆仑KunlunCounty) BDulihri独立石() feud.Barony {
	return c.独立石Dulihri
}

func (c *昆仑KunlunCounty) BPur普尔() feud.Barony {
	return c.普尔Pur
}

func (c *昆仑KunlunCounty) BSeldang色当() feud.Barony {
	return c.色当Seldang
}

func (c *昆仑KunlunCounty) BUlug乌鲁克() feud.Barony {
	return c.乌鲁克Ulug
}

func (c *昆仑KunlunCounty) BYulya月牙() feud.Barony {
	return c.月牙Yulya
}

var CKunlun昆仑 KunlunCounty = &昆仑KunlunCounty{}

func init() {
	f := CKunlun昆仑.(*昆仑KunlunCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1489",
		Title:     "kunlun",
		TitleName: "昆仑",
		TitleCode: "c_kunlun",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿其克Ashi = BAshi阿其克
	f.阿其克Ashi.SetParent(f)

	f.邦达Bangdag = BBangdag邦达
	f.邦达Bangdag.SetParent(f)

	f.独立石Dulihri = BDulihri独立石
	f.独立石Dulihri.SetParent(f)

	f.普尔Pur = BPur普尔
	f.普尔Pur.SetParent(f)

	f.色当Seldang = BSeldang色当
	f.色当Seldang.SetParent(f)

	f.乌鲁克Ulug = BUlug乌鲁克
	f.乌鲁克Ulug.SetParent(f)

	f.月牙Yulya = BYulya月牙
	f.月牙Yulya.SetParent(f)

}
