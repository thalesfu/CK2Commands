package hayya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HayyaCounty interface {
	feud.County
	BAliab阿里阿布() feud.Barony
	BGebeit杰拜特() feud.Barony
	BHayya海亚() feud.Barony
	BKataigarsa卡塔伽萨() feud.Barony
	BMusmer穆斯马尔() feud.Barony
	BSinkat辛卡特() feud.Barony
	BTaris塔里斯() feud.Barony
}

type 海亚HayyaCounty struct {
	feud.BaseCounty
	阿里阿布Aliab      feud.Barony
	杰拜特Gebeit      feud.Barony
	海亚Hayya        feud.Barony
	卡塔伽萨Kataigarsa feud.Barony
	穆斯马尔Musmer     feud.Barony
	辛卡特Sinkat      feud.Barony
	塔里斯Taris       feud.Barony
}

func (c *海亚HayyaCounty) BAliab阿里阿布() feud.Barony {
	return c.阿里阿布Aliab
}

func (c *海亚HayyaCounty) BGebeit杰拜特() feud.Barony {
	return c.杰拜特Gebeit
}

func (c *海亚HayyaCounty) BHayya海亚() feud.Barony {
	return c.海亚Hayya
}

func (c *海亚HayyaCounty) BKataigarsa卡塔伽萨() feud.Barony {
	return c.卡塔伽萨Kataigarsa
}

func (c *海亚HayyaCounty) BMusmer穆斯马尔() feud.Barony {
	return c.穆斯马尔Musmer
}

func (c *海亚HayyaCounty) BSinkat辛卡特() feud.Barony {
	return c.辛卡特Sinkat
}

func (c *海亚HayyaCounty) BTaris塔里斯() feud.Barony {
	return c.塔里斯Taris
}

var CHayya海亚 HayyaCounty = &海亚HayyaCounty{}

func init() {
	f := CHayya海亚.(*海亚HayyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "878",
		Title:     "hayya",
		TitleName: "海亚",
		TitleCode: "c_hayya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿里阿布Aliab = BAliab阿里阿布
	f.阿里阿布Aliab.SetParent(f)

	f.杰拜特Gebeit = BGebeit杰拜特
	f.杰拜特Gebeit.SetParent(f)

	f.海亚Hayya = BHayya海亚
	f.海亚Hayya.SetParent(f)

	f.卡塔伽萨Kataigarsa = BKataigarsa卡塔伽萨
	f.卡塔伽萨Kataigarsa.SetParent(f)

	f.穆斯马尔Musmer = BMusmer穆斯马尔
	f.穆斯马尔Musmer.SetParent(f)

	f.辛卡特Sinkat = BSinkat辛卡特
	f.辛卡特Sinkat.SetParent(f)

	f.塔里斯Taris = BTaris塔里斯
	f.塔里斯Taris.SetParent(f)

}
