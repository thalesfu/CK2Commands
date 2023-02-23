package kaliopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KaliopolisCounty interface {
	feud.County
	BGallipoli加里波利() feud.Barony
	BHeraclea赫拉克利亚() feud.Barony
	BLysimachia利西马其亚() feud.Barony
	BMadyta马蒂塔() feud.Barony
	BPanidos帕尼多斯() feud.Barony
	BRhaidestos莱迪斯托斯() feud.Barony
	BSelymbria瑟利姆布里亚() feud.Barony
	BSestus希斯图斯() feud.Barony
}

type 卡利奥波利斯KaliopolisCounty struct {
	feud.BaseCounty
	加里波利Gallipoli   feud.Barony
	赫拉克利亚Heraclea   feud.Barony
	利西马其亚Lysimachia feud.Barony
	马蒂塔Madyta       feud.Barony
	帕尼多斯Panidos     feud.Barony
	莱迪斯托斯Rhaidestos feud.Barony
	瑟利姆布里亚Selymbria feud.Barony
	希斯图斯Sestus      feud.Barony
}

func (c *卡利奥波利斯KaliopolisCounty) BGallipoli加里波利() feud.Barony {
	return c.加里波利Gallipoli
}

func (c *卡利奥波利斯KaliopolisCounty) BHeraclea赫拉克利亚() feud.Barony {
	return c.赫拉克利亚Heraclea
}

func (c *卡利奥波利斯KaliopolisCounty) BLysimachia利西马其亚() feud.Barony {
	return c.利西马其亚Lysimachia
}

func (c *卡利奥波利斯KaliopolisCounty) BMadyta马蒂塔() feud.Barony {
	return c.马蒂塔Madyta
}

func (c *卡利奥波利斯KaliopolisCounty) BPanidos帕尼多斯() feud.Barony {
	return c.帕尼多斯Panidos
}

func (c *卡利奥波利斯KaliopolisCounty) BRhaidestos莱迪斯托斯() feud.Barony {
	return c.莱迪斯托斯Rhaidestos
}

func (c *卡利奥波利斯KaliopolisCounty) BSelymbria瑟利姆布里亚() feud.Barony {
	return c.瑟利姆布里亚Selymbria
}

func (c *卡利奥波利斯KaliopolisCounty) BSestus希斯图斯() feud.Barony {
	return c.希斯图斯Sestus
}

var CKaliopolis卡利奥波利斯 KaliopolisCounty = &卡利奥波利斯KaliopolisCounty{}

func init() {
	f := CKaliopolis卡利奥波利斯.(*卡利奥波利斯KaliopolisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "495",
		Title:     "kaliopolis",
		TitleName: "卡利奥波利斯",
		TitleCode: "c_kaliopolis",
		Baronies:  map[string]feud.Barony{},
	}

	f.加里波利Gallipoli = BGallipoli加里波利
	f.加里波利Gallipoli.SetParent(f)

	f.赫拉克利亚Heraclea = BHeraclea赫拉克利亚
	f.赫拉克利亚Heraclea.SetParent(f)

	f.利西马其亚Lysimachia = BLysimachia利西马其亚
	f.利西马其亚Lysimachia.SetParent(f)

	f.马蒂塔Madyta = BMadyta马蒂塔
	f.马蒂塔Madyta.SetParent(f)

	f.帕尼多斯Panidos = BPanidos帕尼多斯
	f.帕尼多斯Panidos.SetParent(f)

	f.莱迪斯托斯Rhaidestos = BRhaidestos莱迪斯托斯
	f.莱迪斯托斯Rhaidestos.SetParent(f)

	f.瑟利姆布里亚Selymbria = BSelymbria瑟利姆布里亚
	f.瑟利姆布里亚Selymbria.SetParent(f)

	f.希斯图斯Sestus = BSestus希斯图斯
	f.希斯图斯Sestus.SetParent(f)

}
