package farafra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FarafraCounty interface {
    feud.County
    BAin_bishay艾因比射() 	feud.Barony
    BAin_della艾因德拉() 	feud.Barony
    BFarafra法拉弗拉() 	feud.Barony
    BMut穆特() 	feud.Barony
    BQasrfarfra费拉菲拉堡() 	feud.Barony
    BWadi_hinnis欣尼斯谷() 	feud.Barony
    BWhite_desert白色沙漠() 	feud.Barony
}

type c_farafraFarafraCounty struct {
	feud.BaseCounty
	艾因比射Ain_bishay 	feud.Barony
	艾因德拉Ain_della 	feud.Barony
	法拉弗拉Farafra 	feud.Barony
	穆特Mut 	feud.Barony
	费拉菲拉堡Qasrfarfra 	feud.Barony
	欣尼斯谷Wadi_hinnis 	feud.Barony
	白色沙漠White_desert 	feud.Barony
}

func (c *c_farafraFarafraCounty) BAin_bishay艾因比射() feud.Barony {
	return c.艾因比射Ain_bishay
}
    
func (c *c_farafraFarafraCounty) BAin_della艾因德拉() feud.Barony {
	return c.艾因德拉Ain_della
}
    
func (c *c_farafraFarafraCounty) BFarafra法拉弗拉() feud.Barony {
	return c.法拉弗拉Farafra
}
    
func (c *c_farafraFarafraCounty) BMut穆特() feud.Barony {
	return c.穆特Mut
}
    
func (c *c_farafraFarafraCounty) BQasrfarfra费拉菲拉堡() feud.Barony {
	return c.费拉菲拉堡Qasrfarfra
}
    
func (c *c_farafraFarafraCounty) BWadi_hinnis欣尼斯谷() feud.Barony {
	return c.欣尼斯谷Wadi_hinnis
}
    
func (c *c_farafraFarafraCounty) BWhite_desert白色沙漠() feud.Barony {
	return c.白色沙漠White_desert
}
    
var CFarafrac_farafra FarafraCounty = &c_farafraFarafraCounty{}

func init() {
	f := CFarafrac_farafra.(*c_farafraFarafraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2007",
		Title:     "farafra",
		TitleName: "c_farafra",
		TitleCode: "c_farafra",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾因比射Ain_bishay = BAin_bishay艾因比射
	f.艾因比射Ain_bishay.SetParent(f)
	
	f.艾因德拉Ain_della = BAin_della艾因德拉
	f.艾因德拉Ain_della.SetParent(f)
	
	f.法拉弗拉Farafra = BFarafra法拉弗拉
	f.法拉弗拉Farafra.SetParent(f)
	
	f.穆特Mut = BMut穆特
	f.穆特Mut.SetParent(f)
	
	f.费拉菲拉堡Qasrfarfra = BQasrfarfra费拉菲拉堡
	f.费拉菲拉堡Qasrfarfra.SetParent(f)
	
	f.欣尼斯谷Wadi_hinnis = BWadi_hinnis欣尼斯谷
	f.欣尼斯谷Wadi_hinnis.SetParent(f)
	
	f.白色沙漠White_desert = BWhite_desert白色沙漠
	f.白色沙漠White_desert.SetParent(f)
	
}
