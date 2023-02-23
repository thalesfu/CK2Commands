package faro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FaroCounty interface {
	feud.County
	BAlcoutim阿尔科廷() feud.Barony
	BCastromarim马林堡() feud.Barony
	BFaro法鲁() feud.Barony
	BLoule洛莱() feud.Barony
	BOlhao奥良() feud.Barony
	BSaobrasdealportel圣布拉什迪阿尔波特尔() feud.Barony
	BTavira塔维拉() feud.Barony
}

type 法鲁FaroCounty struct {
	feud.BaseCounty
	阿尔科廷Alcoutim                feud.Barony
	马林堡Castromarim              feud.Barony
	法鲁Faro                      feud.Barony
	洛莱Loule                     feud.Barony
	奥良Olhao                     feud.Barony
	圣布拉什迪阿尔波特尔Saobrasdealportel feud.Barony
	塔维拉Tavira                   feud.Barony
}

func (c *法鲁FaroCounty) BAlcoutim阿尔科廷() feud.Barony {
	return c.阿尔科廷Alcoutim
}

func (c *法鲁FaroCounty) BCastromarim马林堡() feud.Barony {
	return c.马林堡Castromarim
}

func (c *法鲁FaroCounty) BFaro法鲁() feud.Barony {
	return c.法鲁Faro
}

func (c *法鲁FaroCounty) BLoule洛莱() feud.Barony {
	return c.洛莱Loule
}

func (c *法鲁FaroCounty) BOlhao奥良() feud.Barony {
	return c.奥良Olhao
}

func (c *法鲁FaroCounty) BSaobrasdealportel圣布拉什迪阿尔波特尔() feud.Barony {
	return c.圣布拉什迪阿尔波特尔Saobrasdealportel
}

func (c *法鲁FaroCounty) BTavira塔维拉() feud.Barony {
	return c.塔维拉Tavira
}

var CFaro法鲁 FaroCounty = &法鲁FaroCounty{}

func init() {
	f := CFaro法鲁.(*法鲁FaroCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "163",
		Title:     "faro",
		TitleName: "法鲁",
		TitleCode: "c_faro",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔科廷Alcoutim = BAlcoutim阿尔科廷
	f.阿尔科廷Alcoutim.SetParent(f)

	f.马林堡Castromarim = BCastromarim马林堡
	f.马林堡Castromarim.SetParent(f)

	f.法鲁Faro = BFaro法鲁
	f.法鲁Faro.SetParent(f)

	f.洛莱Loule = BLoule洛莱
	f.洛莱Loule.SetParent(f)

	f.奥良Olhao = BOlhao奥良
	f.奥良Olhao.SetParent(f)

	f.圣布拉什迪阿尔波特尔Saobrasdealportel = BSaobrasdealportel圣布拉什迪阿尔波特尔
	f.圣布拉什迪阿尔波特尔Saobrasdealportel.SetParent(f)

	f.塔维拉Tavira = BTavira塔维拉
	f.塔维拉Tavira.SetParent(f)

}
