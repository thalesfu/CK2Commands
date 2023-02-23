package dhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DharaCounty interface {
	feud.County
	BBetma博摩() feud.Barony
	BDepalpur德波补罗() feud.Barony
	BDhara陀罗() feud.Barony
	BDharampuri诃愣补罗() feud.Barony
	BIndore印多尔() feud.Barony
	BJobat焦婆德() feud.Barony
	BSailana萨罗那() feud.Barony
}

type 陀罗DharaCounty struct {
	feud.BaseCounty
	博摩Betma        feud.Barony
	德波补罗Depalpur   feud.Barony
	陀罗Dhara        feud.Barony
	诃愣补罗Dharampuri feud.Barony
	印多尔Indore      feud.Barony
	焦婆德Jobat       feud.Barony
	萨罗那Sailana     feud.Barony
}

func (c *陀罗DharaCounty) BBetma博摩() feud.Barony {
	return c.博摩Betma
}

func (c *陀罗DharaCounty) BDepalpur德波补罗() feud.Barony {
	return c.德波补罗Depalpur
}

func (c *陀罗DharaCounty) BDhara陀罗() feud.Barony {
	return c.陀罗Dhara
}

func (c *陀罗DharaCounty) BDharampuri诃愣补罗() feud.Barony {
	return c.诃愣补罗Dharampuri
}

func (c *陀罗DharaCounty) BIndore印多尔() feud.Barony {
	return c.印多尔Indore
}

func (c *陀罗DharaCounty) BJobat焦婆德() feud.Barony {
	return c.焦婆德Jobat
}

func (c *陀罗DharaCounty) BSailana萨罗那() feud.Barony {
	return c.萨罗那Sailana
}

var CDhara陀罗 DharaCounty = &陀罗DharaCounty{}

func init() {
	f := CDhara陀罗.(*陀罗DharaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1149",
		Title:     "dhara",
		TitleName: "陀罗",
		TitleCode: "c_dhara",
		Baronies:  map[string]feud.Barony{},
	}

	f.博摩Betma = BBetma博摩
	f.博摩Betma.SetParent(f)

	f.德波补罗Depalpur = BDepalpur德波补罗
	f.德波补罗Depalpur.SetParent(f)

	f.陀罗Dhara = BDhara陀罗
	f.陀罗Dhara.SetParent(f)

	f.诃愣补罗Dharampuri = BDharampuri诃愣补罗
	f.诃愣补罗Dharampuri.SetParent(f)

	f.印多尔Indore = BIndore印多尔
	f.印多尔Indore.SetParent(f)

	f.焦婆德Jobat = BJobat焦婆德
	f.焦婆德Jobat.SetParent(f)

	f.萨罗那Sailana = BSailana萨罗那
	f.萨罗那Sailana.SetParent(f)

}
