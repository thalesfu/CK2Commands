package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrnoCounty interface {
    feud.County
    BBoskovice博斯科维采() 	feud.Barony
    BBrno布尔诺() 	feud.Barony
    BHodonin霍多宁() 	feud.Barony
    BMikulcice米库尔奇采() 	feud.Barony
    BPernstejn佩恩什泰因() 	feud.Barony
    BRajhrad赖赫拉德() 	feud.Barony
    BSvitavy斯维塔维() 	feud.Barony
    BVeligrad韦利格勒() 	feud.Barony
    BZdar日贾尔() 	feud.Barony
}

type 布尔诺BrnoCounty struct {
	feud.BaseCounty
	博斯科维采Boskovice 	feud.Barony
	布尔诺Brno 	feud.Barony
	霍多宁Hodonin 	feud.Barony
	米库尔奇采Mikulcice 	feud.Barony
	佩恩什泰因Pernstejn 	feud.Barony
	赖赫拉德Rajhrad 	feud.Barony
	斯维塔维Svitavy 	feud.Barony
	韦利格勒Veligrad 	feud.Barony
	日贾尔Zdar 	feud.Barony
}

func (c *布尔诺BrnoCounty) BBoskovice博斯科维采() feud.Barony {
	return c.博斯科维采Boskovice
}
    
func (c *布尔诺BrnoCounty) BBrno布尔诺() feud.Barony {
	return c.布尔诺Brno
}
    
func (c *布尔诺BrnoCounty) BHodonin霍多宁() feud.Barony {
	return c.霍多宁Hodonin
}
    
func (c *布尔诺BrnoCounty) BMikulcice米库尔奇采() feud.Barony {
	return c.米库尔奇采Mikulcice
}
    
func (c *布尔诺BrnoCounty) BPernstejn佩恩什泰因() feud.Barony {
	return c.佩恩什泰因Pernstejn
}
    
func (c *布尔诺BrnoCounty) BRajhrad赖赫拉德() feud.Barony {
	return c.赖赫拉德Rajhrad
}
    
func (c *布尔诺BrnoCounty) BSvitavy斯维塔维() feud.Barony {
	return c.斯维塔维Svitavy
}
    
func (c *布尔诺BrnoCounty) BVeligrad韦利格勒() feud.Barony {
	return c.韦利格勒Veligrad
}
    
func (c *布尔诺BrnoCounty) BZdar日贾尔() feud.Barony {
	return c.日贾尔Zdar
}
    
var CBrno布尔诺 BrnoCounty = &布尔诺BrnoCounty{}

func init() {
	f := CBrno布尔诺.(*布尔诺BrnoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "441",
		Title:     "brno",
		TitleName: "布尔诺",
		TitleCode: "c_brno",
		Baronies:  map[string]feud.Barony{},
	}

	f.博斯科维采Boskovice = BBoskovice博斯科维采
	f.博斯科维采Boskovice.SetParent(f)
	
	f.布尔诺Brno = BBrno布尔诺
	f.布尔诺Brno.SetParent(f)
	
	f.霍多宁Hodonin = BHodonin霍多宁
	f.霍多宁Hodonin.SetParent(f)
	
	f.米库尔奇采Mikulcice = BMikulcice米库尔奇采
	f.米库尔奇采Mikulcice.SetParent(f)
	
	f.佩恩什泰因Pernstejn = BPernstejn佩恩什泰因
	f.佩恩什泰因Pernstejn.SetParent(f)
	
	f.赖赫拉德Rajhrad = BRajhrad赖赫拉德
	f.赖赫拉德Rajhrad.SetParent(f)
	
	f.斯维塔维Svitavy = BSvitavy斯维塔维
	f.斯维塔维Svitavy.SetParent(f)
	
	f.韦利格勒Veligrad = BVeligrad韦利格勒
	f.韦利格勒Veligrad.SetParent(f)
	
	f.日贾尔Zdar = BZdar日贾尔
	f.日贾尔Zdar.SetParent(f)
	
}
