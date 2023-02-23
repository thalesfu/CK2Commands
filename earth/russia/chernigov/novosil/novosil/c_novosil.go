package novosil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NovosilCounty interface {
	feud.County
	BKhomutovo霍穆托沃() feud.Barony
	BKolpna科尔普纳() feud.Barony
	BLivny利夫内() feud.Barony
	BNovosil诺沃西利() feud.Barony
	BPokrovskoye波克罗夫斯科耶() feud.Barony
	BStrakhovka斯特拉霍夫卡() feud.Barony
	BVerkhovye韦尔霍夫耶() feud.Barony
}

type 诺沃西利NovosilCounty struct {
	feud.BaseCounty
	霍穆托沃Khomutovo      feud.Barony
	科尔普纳Kolpna         feud.Barony
	利夫内Livny           feud.Barony
	诺沃西利Novosil        feud.Barony
	波克罗夫斯科耶Pokrovskoye feud.Barony
	斯特拉霍夫卡Strakhovka   feud.Barony
	韦尔霍夫耶Verkhovye     feud.Barony
}

func (c *诺沃西利NovosilCounty) BKhomutovo霍穆托沃() feud.Barony {
	return c.霍穆托沃Khomutovo
}

func (c *诺沃西利NovosilCounty) BKolpna科尔普纳() feud.Barony {
	return c.科尔普纳Kolpna
}

func (c *诺沃西利NovosilCounty) BLivny利夫内() feud.Barony {
	return c.利夫内Livny
}

func (c *诺沃西利NovosilCounty) BNovosil诺沃西利() feud.Barony {
	return c.诺沃西利Novosil
}

func (c *诺沃西利NovosilCounty) BPokrovskoye波克罗夫斯科耶() feud.Barony {
	return c.波克罗夫斯科耶Pokrovskoye
}

func (c *诺沃西利NovosilCounty) BStrakhovka斯特拉霍夫卡() feud.Barony {
	return c.斯特拉霍夫卡Strakhovka
}

func (c *诺沃西利NovosilCounty) BVerkhovye韦尔霍夫耶() feud.Barony {
	return c.韦尔霍夫耶Verkhovye
}

var CNovosil诺沃西利 NovosilCounty = &诺沃西利NovosilCounty{}

func init() {
	f := CNovosil诺沃西利.(*诺沃西利NovosilCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1634",
		Title:     "novosil",
		TitleName: "诺沃西利",
		TitleCode: "c_novosil",
		Baronies:  map[string]feud.Barony{},
	}

	f.霍穆托沃Khomutovo = BKhomutovo霍穆托沃
	f.霍穆托沃Khomutovo.SetParent(f)

	f.科尔普纳Kolpna = BKolpna科尔普纳
	f.科尔普纳Kolpna.SetParent(f)

	f.利夫内Livny = BLivny利夫内
	f.利夫内Livny.SetParent(f)

	f.诺沃西利Novosil = BNovosil诺沃西利
	f.诺沃西利Novosil.SetParent(f)

	f.波克罗夫斯科耶Pokrovskoye = BPokrovskoye波克罗夫斯科耶
	f.波克罗夫斯科耶Pokrovskoye.SetParent(f)

	f.斯特拉霍夫卡Strakhovka = BStrakhovka斯特拉霍夫卡
	f.斯特拉霍夫卡Strakhovka.SetParent(f)

	f.韦尔霍夫耶Verkhovye = BVerkhovye韦尔霍夫耶
	f.韦尔霍夫耶Verkhovye.SetParent(f)

}
