package diskit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DiskitCounty interface {
    feud.County
    BCharassa查拉萨() 	feud.Barony
    BDiskit德吉() 	feud.Barony
    BHundar匈达() 	feud.Barony
    BPanamik帕纳米克() 	feud.Barony
    BThirit地里特() 	feud.Barony
    BThoise托易斯() 	feud.Barony
    BUnmaru乌玛鲁() 	feud.Barony
}

type 德吉DiskitCounty struct {
	feud.BaseCounty
	查拉萨Charassa 	feud.Barony
	德吉Diskit 	feud.Barony
	匈达Hundar 	feud.Barony
	帕纳米克Panamik 	feud.Barony
	地里特Thirit 	feud.Barony
	托易斯Thoise 	feud.Barony
	乌玛鲁Unmaru 	feud.Barony
}

func (c *德吉DiskitCounty) BCharassa查拉萨() feud.Barony {
	return c.查拉萨Charassa
}
    
func (c *德吉DiskitCounty) BDiskit德吉() feud.Barony {
	return c.德吉Diskit
}
    
func (c *德吉DiskitCounty) BHundar匈达() feud.Barony {
	return c.匈达Hundar
}
    
func (c *德吉DiskitCounty) BPanamik帕纳米克() feud.Barony {
	return c.帕纳米克Panamik
}
    
func (c *德吉DiskitCounty) BThirit地里特() feud.Barony {
	return c.地里特Thirit
}
    
func (c *德吉DiskitCounty) BThoise托易斯() feud.Barony {
	return c.托易斯Thoise
}
    
func (c *德吉DiskitCounty) BUnmaru乌玛鲁() feud.Barony {
	return c.乌玛鲁Unmaru
}
    
var CDiskit德吉 DiskitCounty = &德吉DiskitCounty{}

func init() {
	f := CDiskit德吉.(*德吉DiskitCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1467",
		Title:     "diskit",
		TitleName: "德吉",
		TitleCode: "c_diskit",
		Baronies:  map[string]feud.Barony{},
	}

	f.查拉萨Charassa = BCharassa查拉萨
	f.查拉萨Charassa.SetParent(f)
	
	f.德吉Diskit = BDiskit德吉
	f.德吉Diskit.SetParent(f)
	
	f.匈达Hundar = BHundar匈达
	f.匈达Hundar.SetParent(f)
	
	f.帕纳米克Panamik = BPanamik帕纳米克
	f.帕纳米克Panamik.SetParent(f)
	
	f.地里特Thirit = BThirit地里特
	f.地里特Thirit.SetParent(f)
	
	f.托易斯Thoise = BThoise托易斯
	f.托易斯Thoise.SetParent(f)
	
	f.乌玛鲁Unmaru = BUnmaru乌玛鲁
	f.乌玛鲁Unmaru.SetParent(f)
	
}
