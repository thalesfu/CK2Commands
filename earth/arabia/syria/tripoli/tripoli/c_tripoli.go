package tripoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TripoliCounty interface {
	feud.County
	BAlqadmus卡德穆斯() feud.Barony
	BAlqulayat古莱阿特() feud.Barony
	BArqah阿拉赫() feud.Barony
	BBesmedin比梅丁() feud.Barony
	BBoutron拜特龙() feud.Barony
	BGibelet季比莱特() feud.Barony
	BNephin涅芬() feud.Barony
	BSyrtripoli的黎波里() feud.Barony
}

type 的黎波里TripoliCounty struct {
	feud.BaseCounty
	卡德穆斯Alqadmus   feud.Barony
	古莱阿特Alqulayat  feud.Barony
	阿拉赫Arqah       feud.Barony
	比梅丁Besmedin    feud.Barony
	拜特龙Boutron     feud.Barony
	季比莱特Gibelet    feud.Barony
	涅芬Nephin       feud.Barony
	的黎波里Syrtripoli feud.Barony
}

func (c *的黎波里TripoliCounty) BAlqadmus卡德穆斯() feud.Barony {
	return c.卡德穆斯Alqadmus
}

func (c *的黎波里TripoliCounty) BAlqulayat古莱阿特() feud.Barony {
	return c.古莱阿特Alqulayat
}

func (c *的黎波里TripoliCounty) BArqah阿拉赫() feud.Barony {
	return c.阿拉赫Arqah
}

func (c *的黎波里TripoliCounty) BBesmedin比梅丁() feud.Barony {
	return c.比梅丁Besmedin
}

func (c *的黎波里TripoliCounty) BBoutron拜特龙() feud.Barony {
	return c.拜特龙Boutron
}

func (c *的黎波里TripoliCounty) BGibelet季比莱特() feud.Barony {
	return c.季比莱特Gibelet
}

func (c *的黎波里TripoliCounty) BNephin涅芬() feud.Barony {
	return c.涅芬Nephin
}

func (c *的黎波里TripoliCounty) BSyrtripoli的黎波里() feud.Barony {
	return c.的黎波里Syrtripoli
}

var CTripoli的黎波里 TripoliCounty = &的黎波里TripoliCounty{}

func init() {
	f := CTripoli的黎波里.(*的黎波里TripoliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "767",
		Title:     "tripoli",
		TitleName: "的黎波里",
		TitleCode: "c_tripoli",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡德穆斯Alqadmus = BAlqadmus卡德穆斯
	f.卡德穆斯Alqadmus.SetParent(f)

	f.古莱阿特Alqulayat = BAlqulayat古莱阿特
	f.古莱阿特Alqulayat.SetParent(f)

	f.阿拉赫Arqah = BArqah阿拉赫
	f.阿拉赫Arqah.SetParent(f)

	f.比梅丁Besmedin = BBesmedin比梅丁
	f.比梅丁Besmedin.SetParent(f)

	f.拜特龙Boutron = BBoutron拜特龙
	f.拜特龙Boutron.SetParent(f)

	f.季比莱特Gibelet = BGibelet季比莱特
	f.季比莱特Gibelet.SetParent(f)

	f.涅芬Nephin = BNephin涅芬
	f.涅芬Nephin.SetParent(f)

	f.的黎波里Syrtripoli = BSyrtripoli的黎波里
	f.的黎波里Syrtripoli.SetParent(f)

}
