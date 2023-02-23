package urbino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UrbinoCounty interface {
	feud.County
	BCittadicastello卡斯泰洛城() feud.Barony
	BFano法诺() feud.Barony
	BFossombrone福松布罗内() feud.Barony
	BMontefeltro蒙特费尔特罗() feud.Barony
	BSanmarino圣马力诺() feud.Barony
	BSanseverino圣塞韦里诺() feud.Barony
	BUrbino乌尔比诺() feud.Barony
}

type 乌尔比诺UrbinoCounty struct {
	feud.BaseCounty
	卡斯泰洛城Cittadicastello feud.Barony
	法诺Fano               feud.Barony
	福松布罗内Fossombrone     feud.Barony
	蒙特费尔特罗Montefeltro    feud.Barony
	圣马力诺Sanmarino        feud.Barony
	圣塞韦里诺Sanseverino     feud.Barony
	乌尔比诺Urbino           feud.Barony
}

func (c *乌尔比诺UrbinoCounty) BCittadicastello卡斯泰洛城() feud.Barony {
	return c.卡斯泰洛城Cittadicastello
}

func (c *乌尔比诺UrbinoCounty) BFano法诺() feud.Barony {
	return c.法诺Fano
}

func (c *乌尔比诺UrbinoCounty) BFossombrone福松布罗内() feud.Barony {
	return c.福松布罗内Fossombrone
}

func (c *乌尔比诺UrbinoCounty) BMontefeltro蒙特费尔特罗() feud.Barony {
	return c.蒙特费尔特罗Montefeltro
}

func (c *乌尔比诺UrbinoCounty) BSanmarino圣马力诺() feud.Barony {
	return c.圣马力诺Sanmarino
}

func (c *乌尔比诺UrbinoCounty) BSanseverino圣塞韦里诺() feud.Barony {
	return c.圣塞韦里诺Sanseverino
}

func (c *乌尔比诺UrbinoCounty) BUrbino乌尔比诺() feud.Barony {
	return c.乌尔比诺Urbino
}

var CUrbino乌尔比诺 UrbinoCounty = &乌尔比诺UrbinoCounty{}

func init() {
	f := CUrbino乌尔比诺.(*乌尔比诺UrbinoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "329",
		Title:     "urbino",
		TitleName: "乌尔比诺",
		TitleCode: "c_urbino",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡斯泰洛城Cittadicastello = BCittadicastello卡斯泰洛城
	f.卡斯泰洛城Cittadicastello.SetParent(f)

	f.法诺Fano = BFano法诺
	f.法诺Fano.SetParent(f)

	f.福松布罗内Fossombrone = BFossombrone福松布罗内
	f.福松布罗内Fossombrone.SetParent(f)

	f.蒙特费尔特罗Montefeltro = BMontefeltro蒙特费尔特罗
	f.蒙特费尔特罗Montefeltro.SetParent(f)

	f.圣马力诺Sanmarino = BSanmarino圣马力诺
	f.圣马力诺Sanmarino.SetParent(f)

	f.圣塞韦里诺Sanseverino = BSanseverino圣塞韦里诺
	f.圣塞韦里诺Sanseverino.SetParent(f)

	f.乌尔比诺Urbino = BUrbino乌尔比诺
	f.乌尔比诺Urbino.SetParent(f)

}
