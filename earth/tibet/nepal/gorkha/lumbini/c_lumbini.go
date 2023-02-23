package lumbini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LumbiniCounty interface {
	feud.County
	BBhairahawa陪胪诃伐() feud.Barony
	BLumbini岚毗尼() feud.Barony
	BMusikot牟斯拘吒() feud.Barony
	BNigrodharama尼拘陀园() feud.Barony
	BSandhikharka桑提迦伽() feud.Barony
	BTaulihawa道利赫瓦() feud.Barony
	BTulsipur都罗私补罗() feud.Barony
}

type 岚毗尼LumbiniCounty struct {
	feud.BaseCounty
	陪胪诃伐Bhairahawa   feud.Barony
	岚毗尼Lumbini       feud.Barony
	牟斯拘吒Musikot      feud.Barony
	尼拘陀园Nigrodharama feud.Barony
	桑提迦伽Sandhikharka feud.Barony
	道利赫瓦Taulihawa    feud.Barony
	都罗私补罗Tulsipur    feud.Barony
}

func (c *岚毗尼LumbiniCounty) BBhairahawa陪胪诃伐() feud.Barony {
	return c.陪胪诃伐Bhairahawa
}

func (c *岚毗尼LumbiniCounty) BLumbini岚毗尼() feud.Barony {
	return c.岚毗尼Lumbini
}

func (c *岚毗尼LumbiniCounty) BMusikot牟斯拘吒() feud.Barony {
	return c.牟斯拘吒Musikot
}

func (c *岚毗尼LumbiniCounty) BNigrodharama尼拘陀园() feud.Barony {
	return c.尼拘陀园Nigrodharama
}

func (c *岚毗尼LumbiniCounty) BSandhikharka桑提迦伽() feud.Barony {
	return c.桑提迦伽Sandhikharka
}

func (c *岚毗尼LumbiniCounty) BTaulihawa道利赫瓦() feud.Barony {
	return c.道利赫瓦Taulihawa
}

func (c *岚毗尼LumbiniCounty) BTulsipur都罗私补罗() feud.Barony {
	return c.都罗私补罗Tulsipur
}

var CLumbini岚毗尼 LumbiniCounty = &岚毗尼LumbiniCounty{}

func init() {
	f := CLumbini岚毗尼.(*岚毗尼LumbiniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1474",
		Title:     "lumbini",
		TitleName: "岚毗尼",
		TitleCode: "c_lumbini",
		Baronies:  map[string]feud.Barony{},
	}

	f.陪胪诃伐Bhairahawa = BBhairahawa陪胪诃伐
	f.陪胪诃伐Bhairahawa.SetParent(f)

	f.岚毗尼Lumbini = BLumbini岚毗尼
	f.岚毗尼Lumbini.SetParent(f)

	f.牟斯拘吒Musikot = BMusikot牟斯拘吒
	f.牟斯拘吒Musikot.SetParent(f)

	f.尼拘陀园Nigrodharama = BNigrodharama尼拘陀园
	f.尼拘陀园Nigrodharama.SetParent(f)

	f.桑提迦伽Sandhikharka = BSandhikharka桑提迦伽
	f.桑提迦伽Sandhikharka.SetParent(f)

	f.道利赫瓦Taulihawa = BTaulihawa道利赫瓦
	f.道利赫瓦Taulihawa.SetParent(f)

	f.都罗私补罗Tulsipur = BTulsipur都罗私补罗
	f.都罗私补罗Tulsipur.SetParent(f)

}
