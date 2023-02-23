package valencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ValenciaCounty interface {
	feud.County
	BAlacuas阿拉夸斯() feud.Barony
	BAlberique阿尔韦里克() feud.Barony
	BChiva奇瓦() feud.Barony
	BCuartdepoblet夸尔特德波布莱特() feud.Barony
	BGandia甘迪亚() feud.Barony
	BJativa沙蒂瓦() feud.Barony
	BTorrent托伦特() feud.Barony
	BValencia瓦伦西亚() feud.Barony
}

type 瓦伦西亚ValenciaCounty struct {
	feud.BaseCounty
	阿拉夸斯Alacuas           feud.Barony
	阿尔韦里克Alberique        feud.Barony
	奇瓦Chiva               feud.Barony
	夸尔特德波布莱特Cuartdepoblet feud.Barony
	甘迪亚Gandia             feud.Barony
	沙蒂瓦Jativa             feud.Barony
	托伦特Torrent            feud.Barony
	瓦伦西亚Valencia          feud.Barony
}

func (c *瓦伦西亚ValenciaCounty) BAlacuas阿拉夸斯() feud.Barony {
	return c.阿拉夸斯Alacuas
}

func (c *瓦伦西亚ValenciaCounty) BAlberique阿尔韦里克() feud.Barony {
	return c.阿尔韦里克Alberique
}

func (c *瓦伦西亚ValenciaCounty) BChiva奇瓦() feud.Barony {
	return c.奇瓦Chiva
}

func (c *瓦伦西亚ValenciaCounty) BCuartdepoblet夸尔特德波布莱特() feud.Barony {
	return c.夸尔特德波布莱特Cuartdepoblet
}

func (c *瓦伦西亚ValenciaCounty) BGandia甘迪亚() feud.Barony {
	return c.甘迪亚Gandia
}

func (c *瓦伦西亚ValenciaCounty) BJativa沙蒂瓦() feud.Barony {
	return c.沙蒂瓦Jativa
}

func (c *瓦伦西亚ValenciaCounty) BTorrent托伦特() feud.Barony {
	return c.托伦特Torrent
}

func (c *瓦伦西亚ValenciaCounty) BValencia瓦伦西亚() feud.Barony {
	return c.瓦伦西亚Valencia
}

var CValencia瓦伦西亚 ValenciaCounty = &瓦伦西亚ValenciaCounty{}

func init() {
	f := CValencia瓦伦西亚.(*瓦伦西亚ValenciaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "171",
		Title:     "valencia",
		TitleName: "瓦伦西亚",
		TitleCode: "c_valencia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉夸斯Alacuas = BAlacuas阿拉夸斯
	f.阿拉夸斯Alacuas.SetParent(f)

	f.阿尔韦里克Alberique = BAlberique阿尔韦里克
	f.阿尔韦里克Alberique.SetParent(f)

	f.奇瓦Chiva = BChiva奇瓦
	f.奇瓦Chiva.SetParent(f)

	f.夸尔特德波布莱特Cuartdepoblet = BCuartdepoblet夸尔特德波布莱特
	f.夸尔特德波布莱特Cuartdepoblet.SetParent(f)

	f.甘迪亚Gandia = BGandia甘迪亚
	f.甘迪亚Gandia.SetParent(f)

	f.沙蒂瓦Jativa = BJativa沙蒂瓦
	f.沙蒂瓦Jativa.SetParent(f)

	f.托伦特Torrent = BTorrent托伦特
	f.托伦特Torrent.SetParent(f)

	f.瓦伦西亚Valencia = BValencia瓦伦西亚
	f.瓦伦西亚Valencia.SetParent(f)

}
