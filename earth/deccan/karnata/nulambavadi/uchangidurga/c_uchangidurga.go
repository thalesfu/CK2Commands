package uchangidurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UchangidurgaCounty interface {
    feud.County
    BChitradurga质多罗突伽() 	feud.Barony
    BHarihar赫里赫尔() 	feud.Barony
    BHemavati醯摩伐底() 	feud.Barony
    BNidugallu尼杜加卢() 	feud.Barony
    BParivi帕里维() 	feud.Barony
    BSira尸罗() 	feud.Barony
    BUchangidurga郁昌耆突伽() 	feud.Barony
}

type 郁昌耆突伽UchangidurgaCounty struct {
	feud.BaseCounty
	质多罗突伽Chitradurga 	feud.Barony
	赫里赫尔Harihar 	feud.Barony
	醯摩伐底Hemavati 	feud.Barony
	尼杜加卢Nidugallu 	feud.Barony
	帕里维Parivi 	feud.Barony
	尸罗Sira 	feud.Barony
	郁昌耆突伽Uchangidurga 	feud.Barony
}

func (c *郁昌耆突伽UchangidurgaCounty) BChitradurga质多罗突伽() feud.Barony {
	return c.质多罗突伽Chitradurga
}
    
func (c *郁昌耆突伽UchangidurgaCounty) BHarihar赫里赫尔() feud.Barony {
	return c.赫里赫尔Harihar
}
    
func (c *郁昌耆突伽UchangidurgaCounty) BHemavati醯摩伐底() feud.Barony {
	return c.醯摩伐底Hemavati
}
    
func (c *郁昌耆突伽UchangidurgaCounty) BNidugallu尼杜加卢() feud.Barony {
	return c.尼杜加卢Nidugallu
}
    
func (c *郁昌耆突伽UchangidurgaCounty) BParivi帕里维() feud.Barony {
	return c.帕里维Parivi
}
    
func (c *郁昌耆突伽UchangidurgaCounty) BSira尸罗() feud.Barony {
	return c.尸罗Sira
}
    
func (c *郁昌耆突伽UchangidurgaCounty) BUchangidurga郁昌耆突伽() feud.Barony {
	return c.郁昌耆突伽Uchangidurga
}
    
var CUchangidurga郁昌耆突伽 UchangidurgaCounty = &郁昌耆突伽UchangidurgaCounty{}

func init() {
	f := CUchangidurga郁昌耆突伽.(*郁昌耆突伽UchangidurgaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1121",
		Title:     "uchangidurga",
		TitleName: "郁昌耆突伽",
		TitleCode: "c_uchangidurga",
		Baronies:  map[string]feud.Barony{},
	}

	f.质多罗突伽Chitradurga = BChitradurga质多罗突伽
	f.质多罗突伽Chitradurga.SetParent(f)
	
	f.赫里赫尔Harihar = BHarihar赫里赫尔
	f.赫里赫尔Harihar.SetParent(f)
	
	f.醯摩伐底Hemavati = BHemavati醯摩伐底
	f.醯摩伐底Hemavati.SetParent(f)
	
	f.尼杜加卢Nidugallu = BNidugallu尼杜加卢
	f.尼杜加卢Nidugallu.SetParent(f)
	
	f.帕里维Parivi = BParivi帕里维
	f.帕里维Parivi.SetParent(f)
	
	f.尸罗Sira = BSira尸罗
	f.尸罗Sira.SetParent(f)
	
	f.郁昌耆突伽Uchangidurga = BUchangidurga郁昌耆突伽
	f.郁昌耆突伽Uchangidurga.SetParent(f)
	
}
