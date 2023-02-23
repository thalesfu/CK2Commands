package benghazi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BenghaziCounty interface {
	feud.County
	BAjdabiya艾季达比耶() feud.Barony
	BBarca巴尔切() feud.Barony
	BBenghazi班加西() feud.Barony
	BDaryanah代尔亚奈() feud.Barony
	BJardinah杰尔迪奈() feud.Barony
	BOriana奥利安() feud.Barony
	BTansulukh坦苏鲁赫() feud.Barony
	BTulmaytath图勒迈塞() feud.Barony
}

type 班加西BenghaziCounty struct {
	feud.BaseCounty
	艾季达比耶Ajdabiya  feud.Barony
	巴尔切Barca       feud.Barony
	班加西Benghazi    feud.Barony
	代尔亚奈Daryanah   feud.Barony
	杰尔迪奈Jardinah   feud.Barony
	奥利安Oriana      feud.Barony
	坦苏鲁赫Tansulukh  feud.Barony
	图勒迈塞Tulmaytath feud.Barony
}

func (c *班加西BenghaziCounty) BAjdabiya艾季达比耶() feud.Barony {
	return c.艾季达比耶Ajdabiya
}

func (c *班加西BenghaziCounty) BBarca巴尔切() feud.Barony {
	return c.巴尔切Barca
}

func (c *班加西BenghaziCounty) BBenghazi班加西() feud.Barony {
	return c.班加西Benghazi
}

func (c *班加西BenghaziCounty) BDaryanah代尔亚奈() feud.Barony {
	return c.代尔亚奈Daryanah
}

func (c *班加西BenghaziCounty) BJardinah杰尔迪奈() feud.Barony {
	return c.杰尔迪奈Jardinah
}

func (c *班加西BenghaziCounty) BOriana奥利安() feud.Barony {
	return c.奥利安Oriana
}

func (c *班加西BenghaziCounty) BTansulukh坦苏鲁赫() feud.Barony {
	return c.坦苏鲁赫Tansulukh
}

func (c *班加西BenghaziCounty) BTulmaytath图勒迈塞() feud.Barony {
	return c.图勒迈塞Tulmaytath
}

var CBenghazi班加西 BenghaziCounty = &班加西BenghaziCounty{}

func init() {
	f := CBenghazi班加西.(*班加西BenghaziCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "808",
		Title:     "benghazi",
		TitleName: "班加西",
		TitleCode: "c_benghazi",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾季达比耶Ajdabiya = BAjdabiya艾季达比耶
	f.艾季达比耶Ajdabiya.SetParent(f)

	f.巴尔切Barca = BBarca巴尔切
	f.巴尔切Barca.SetParent(f)

	f.班加西Benghazi = BBenghazi班加西
	f.班加西Benghazi.SetParent(f)

	f.代尔亚奈Daryanah = BDaryanah代尔亚奈
	f.代尔亚奈Daryanah.SetParent(f)

	f.杰尔迪奈Jardinah = BJardinah杰尔迪奈
	f.杰尔迪奈Jardinah.SetParent(f)

	f.奥利安Oriana = BOriana奥利安
	f.奥利安Oriana.SetParent(f)

	f.坦苏鲁赫Tansulukh = BTansulukh坦苏鲁赫
	f.坦苏鲁赫Tansulukh.SetParent(f)

	f.图勒迈塞Tulmaytath = BTulmaytath图勒迈塞
	f.图勒迈塞Tulmaytath.SetParent(f)

}
