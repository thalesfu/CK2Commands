package aquileia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AquileiaCounty interface {
	feud.County
	BAquileia阿奎莱亚() feud.Barony
	BConcordia康科迪亚() feud.Barony
	BGrado格拉多() feud.Barony
	BPortoguaro波尔托格鲁阿罗() feud.Barony
	BSacile萨奇莱() feud.Barony
	BSistiana西斯蒂亚那() feud.Barony
	BTrieste的里雅斯特() feud.Barony
}

type 阿奎莱亚AquileiaCounty struct {
	feud.BaseCounty
	阿奎莱亚Aquileia      feud.Barony
	康科迪亚Concordia     feud.Barony
	格拉多Grado          feud.Barony
	波尔托格鲁阿罗Portoguaro feud.Barony
	萨奇莱Sacile         feud.Barony
	西斯蒂亚那Sistiana     feud.Barony
	的里雅斯特Trieste      feud.Barony
}

func (c *阿奎莱亚AquileiaCounty) BAquileia阿奎莱亚() feud.Barony {
	return c.阿奎莱亚Aquileia
}

func (c *阿奎莱亚AquileiaCounty) BConcordia康科迪亚() feud.Barony {
	return c.康科迪亚Concordia
}

func (c *阿奎莱亚AquileiaCounty) BGrado格拉多() feud.Barony {
	return c.格拉多Grado
}

func (c *阿奎莱亚AquileiaCounty) BPortoguaro波尔托格鲁阿罗() feud.Barony {
	return c.波尔托格鲁阿罗Portoguaro
}

func (c *阿奎莱亚AquileiaCounty) BSacile萨奇莱() feud.Barony {
	return c.萨奇莱Sacile
}

func (c *阿奎莱亚AquileiaCounty) BSistiana西斯蒂亚那() feud.Barony {
	return c.西斯蒂亚那Sistiana
}

func (c *阿奎莱亚AquileiaCounty) BTrieste的里雅斯特() feud.Barony {
	return c.的里雅斯特Trieste
}

var CAquileia阿奎莱亚 AquileiaCounty = &阿奎莱亚AquileiaCounty{}

func init() {
	f := CAquileia阿奎莱亚.(*阿奎莱亚AquileiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "358",
		Title:     "aquileia",
		TitleName: "阿奎莱亚",
		TitleCode: "c_aquileia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿奎莱亚Aquileia = BAquileia阿奎莱亚
	f.阿奎莱亚Aquileia.SetParent(f)

	f.康科迪亚Concordia = BConcordia康科迪亚
	f.康科迪亚Concordia.SetParent(f)

	f.格拉多Grado = BGrado格拉多
	f.格拉多Grado.SetParent(f)

	f.波尔托格鲁阿罗Portoguaro = BPortoguaro波尔托格鲁阿罗
	f.波尔托格鲁阿罗Portoguaro.SetParent(f)

	f.萨奇莱Sacile = BSacile萨奇莱
	f.萨奇莱Sacile.SetParent(f)

	f.西斯蒂亚那Sistiana = BSistiana西斯蒂亚那
	f.西斯蒂亚那Sistiana.SetParent(f)

	f.的里雅斯特Trieste = BTrieste的里雅斯特
	f.的里雅斯特Trieste.SetParent(f)

}
