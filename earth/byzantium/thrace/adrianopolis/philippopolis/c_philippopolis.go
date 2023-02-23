package philippopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PhilippopolisCounty interface {
	feud.County
	BBachkovo巴奇科沃() feud.Barony
	BHaskovo哈斯科沃() feud.Barony
	BKlokoknitsa克罗克尼萨() feud.Barony
	BPetrich彼得里奇() feud.Barony
	BPhilippopolis腓利波波利斯() feud.Barony
	BProdromos普罗德摩斯() feud.Barony
	BStenimachos斯忒尼马科斯() feud.Barony
}

type 腓利波波利斯PhilippopolisCounty struct {
	feud.BaseCounty
	巴奇科沃Bachkovo        feud.Barony
	哈斯科沃Haskovo         feud.Barony
	克罗克尼萨Klokoknitsa    feud.Barony
	彼得里奇Petrich         feud.Barony
	腓利波波利斯Philippopolis feud.Barony
	普罗德摩斯Prodromos      feud.Barony
	斯忒尼马科斯Stenimachos   feud.Barony
}

func (c *腓利波波利斯PhilippopolisCounty) BBachkovo巴奇科沃() feud.Barony {
	return c.巴奇科沃Bachkovo
}

func (c *腓利波波利斯PhilippopolisCounty) BHaskovo哈斯科沃() feud.Barony {
	return c.哈斯科沃Haskovo
}

func (c *腓利波波利斯PhilippopolisCounty) BKlokoknitsa克罗克尼萨() feud.Barony {
	return c.克罗克尼萨Klokoknitsa
}

func (c *腓利波波利斯PhilippopolisCounty) BPetrich彼得里奇() feud.Barony {
	return c.彼得里奇Petrich
}

func (c *腓利波波利斯PhilippopolisCounty) BPhilippopolis腓利波波利斯() feud.Barony {
	return c.腓利波波利斯Philippopolis
}

func (c *腓利波波利斯PhilippopolisCounty) BProdromos普罗德摩斯() feud.Barony {
	return c.普罗德摩斯Prodromos
}

func (c *腓利波波利斯PhilippopolisCounty) BStenimachos斯忒尼马科斯() feud.Barony {
	return c.斯忒尼马科斯Stenimachos
}

var CPhilippopolis腓利波波利斯 PhilippopolisCounty = &腓利波波利斯PhilippopolisCounty{}

func init() {
	f := CPhilippopolis腓利波波利斯.(*腓利波波利斯PhilippopolisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "493",
		Title:     "philippopolis",
		TitleName: "腓利波波利斯",
		TitleCode: "c_philippopolis",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴奇科沃Bachkovo = BBachkovo巴奇科沃
	f.巴奇科沃Bachkovo.SetParent(f)

	f.哈斯科沃Haskovo = BHaskovo哈斯科沃
	f.哈斯科沃Haskovo.SetParent(f)

	f.克罗克尼萨Klokoknitsa = BKlokoknitsa克罗克尼萨
	f.克罗克尼萨Klokoknitsa.SetParent(f)

	f.彼得里奇Petrich = BPetrich彼得里奇
	f.彼得里奇Petrich.SetParent(f)

	f.腓利波波利斯Philippopolis = BPhilippopolis腓利波波利斯
	f.腓利波波利斯Philippopolis.SetParent(f)

	f.普罗德摩斯Prodromos = BProdromos普罗德摩斯
	f.普罗德摩斯Prodromos.SetParent(f)

	f.斯忒尼马科斯Stenimachos = BStenimachos斯忒尼马科斯
	f.斯忒尼马科斯Stenimachos.SetParent(f)

}
