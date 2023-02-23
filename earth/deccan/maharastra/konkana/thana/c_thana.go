package thana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThanaCounty interface {
	feud.County
	BBassein巴森() feud.Barony
	BChaul招罗() feud.Barony
	BDabhol达布尔() feud.Barony
	BJanjira缮恃罗() feud.Barony
	BKanhagiri建诃耆厘() feud.Barony
	BMahim马哈音() feud.Barony
	BThana都奴何() feud.Barony
}

type 都奴何ThanaCounty struct {
	feud.BaseCounty
	巴森Bassein     feud.Barony
	招罗Chaul       feud.Barony
	达布尔Dabhol     feud.Barony
	缮恃罗Janjira    feud.Barony
	建诃耆厘Kanhagiri feud.Barony
	马哈音Mahim      feud.Barony
	都奴何Thana      feud.Barony
}

func (c *都奴何ThanaCounty) BBassein巴森() feud.Barony {
	return c.巴森Bassein
}

func (c *都奴何ThanaCounty) BChaul招罗() feud.Barony {
	return c.招罗Chaul
}

func (c *都奴何ThanaCounty) BDabhol达布尔() feud.Barony {
	return c.达布尔Dabhol
}

func (c *都奴何ThanaCounty) BJanjira缮恃罗() feud.Barony {
	return c.缮恃罗Janjira
}

func (c *都奴何ThanaCounty) BKanhagiri建诃耆厘() feud.Barony {
	return c.建诃耆厘Kanhagiri
}

func (c *都奴何ThanaCounty) BMahim马哈音() feud.Barony {
	return c.马哈音Mahim
}

func (c *都奴何ThanaCounty) BThana都奴何() feud.Barony {
	return c.都奴何Thana
}

var CThana都奴何 ThanaCounty = &都奴何ThanaCounty{}

func init() {
	f := CThana都奴何.(*都奴何ThanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1125",
		Title:     "thana",
		TitleName: "都奴何",
		TitleCode: "c_thana",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴森Bassein = BBassein巴森
	f.巴森Bassein.SetParent(f)

	f.招罗Chaul = BChaul招罗
	f.招罗Chaul.SetParent(f)

	f.达布尔Dabhol = BDabhol达布尔
	f.达布尔Dabhol.SetParent(f)

	f.缮恃罗Janjira = BJanjira缮恃罗
	f.缮恃罗Janjira.SetParent(f)

	f.建诃耆厘Kanhagiri = BKanhagiri建诃耆厘
	f.建诃耆厘Kanhagiri.SetParent(f)

	f.马哈音Mahim = BMahim马哈音
	f.马哈音Mahim.SetParent(f)

	f.都奴何Thana = BThana都奴何
	f.都奴何Thana.SetParent(f)

}
