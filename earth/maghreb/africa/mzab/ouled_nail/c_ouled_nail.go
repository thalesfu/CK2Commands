package ouled_nail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Ouled_nailCounty interface {
    feud.County
    BAinzia艾因兹() 	feud.Barony
    BAl_aghwat艾格瓦特() 	feud.Barony
    BBeneldir本迪尔() 	feud.Barony
    BBordjelcaid布尔吉采德() 	feud.Barony
    BMsila姆西拉() 	feud.Barony
    BTilghemt提勒盖姆特() 	feud.Barony
    BTubirett布依拉() 	feud.Barony
}

type 艾格瓦特Ouled_nailCounty struct {
	feud.BaseCounty
	艾因兹Ainzia 	feud.Barony
	艾格瓦特Al_aghwat 	feud.Barony
	本迪尔Beneldir 	feud.Barony
	布尔吉采德Bordjelcaid 	feud.Barony
	姆西拉Msila 	feud.Barony
	提勒盖姆特Tilghemt 	feud.Barony
	布依拉Tubirett 	feud.Barony
}

func (c *艾格瓦特Ouled_nailCounty) BAinzia艾因兹() feud.Barony {
	return c.艾因兹Ainzia
}
    
func (c *艾格瓦特Ouled_nailCounty) BAl_aghwat艾格瓦特() feud.Barony {
	return c.艾格瓦特Al_aghwat
}
    
func (c *艾格瓦特Ouled_nailCounty) BBeneldir本迪尔() feud.Barony {
	return c.本迪尔Beneldir
}
    
func (c *艾格瓦特Ouled_nailCounty) BBordjelcaid布尔吉采德() feud.Barony {
	return c.布尔吉采德Bordjelcaid
}
    
func (c *艾格瓦特Ouled_nailCounty) BMsila姆西拉() feud.Barony {
	return c.姆西拉Msila
}
    
func (c *艾格瓦特Ouled_nailCounty) BTilghemt提勒盖姆特() feud.Barony {
	return c.提勒盖姆特Tilghemt
}
    
func (c *艾格瓦特Ouled_nailCounty) BTubirett布依拉() feud.Barony {
	return c.布依拉Tubirett
}
    
var COuled_nail艾格瓦特 Ouled_nailCounty = &艾格瓦特Ouled_nailCounty{}

func init() {
	f := COuled_nail艾格瓦特.(*艾格瓦特Ouled_nailCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "828",
		Title:     "ouled_nail",
		TitleName: "艾格瓦特",
		TitleCode: "c_ouled_nail",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾因兹Ainzia = BAinzia艾因兹
	f.艾因兹Ainzia.SetParent(f)
	
	f.艾格瓦特Al_aghwat = BAl_aghwat艾格瓦特
	f.艾格瓦特Al_aghwat.SetParent(f)
	
	f.本迪尔Beneldir = BBeneldir本迪尔
	f.本迪尔Beneldir.SetParent(f)
	
	f.布尔吉采德Bordjelcaid = BBordjelcaid布尔吉采德
	f.布尔吉采德Bordjelcaid.SetParent(f)
	
	f.姆西拉Msila = BMsila姆西拉
	f.姆西拉Msila.SetParent(f)
	
	f.提勒盖姆特Tilghemt = BTilghemt提勒盖姆特
	f.提勒盖姆特Tilghemt.SetParent(f)
	
	f.布依拉Tubirett = BTubirett布依拉
	f.布依拉Tubirett.SetParent(f)
	
}
