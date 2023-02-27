package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BanavasiCounty interface {
    feud.County
    BBalligavi巴里加维() 	feud.Barony
    BCandragutti旃陀罗古提() 	feud.Barony
    BHangal汉加尔() 	feud.Barony
    BIkkeri伊凯里() 	feud.Barony
    BKuruvatti俱卢伐底() 	feud.Barony
    BLakkundi拉昆迪() 	feud.Barony
    BSringeri斯灵盖里() 	feud.Barony
    BVaijayanti吠阇演帝() 	feud.Barony
    BVankapura房迦补罗() 	feud.Barony
}

type 巴南巴西BanavasiCounty struct {
	feud.BaseCounty
	巴里加维Balligavi 	feud.Barony
	旃陀罗古提Candragutti 	feud.Barony
	汉加尔Hangal 	feud.Barony
	伊凯里Ikkeri 	feud.Barony
	俱卢伐底Kuruvatti 	feud.Barony
	拉昆迪Lakkundi 	feud.Barony
	斯灵盖里Sringeri 	feud.Barony
	吠阇演帝Vaijayanti 	feud.Barony
	房迦补罗Vankapura 	feud.Barony
}

func (c *巴南巴西BanavasiCounty) BBalligavi巴里加维() feud.Barony {
	return c.巴里加维Balligavi
}
    
func (c *巴南巴西BanavasiCounty) BCandragutti旃陀罗古提() feud.Barony {
	return c.旃陀罗古提Candragutti
}
    
func (c *巴南巴西BanavasiCounty) BHangal汉加尔() feud.Barony {
	return c.汉加尔Hangal
}
    
func (c *巴南巴西BanavasiCounty) BIkkeri伊凯里() feud.Barony {
	return c.伊凯里Ikkeri
}
    
func (c *巴南巴西BanavasiCounty) BKuruvatti俱卢伐底() feud.Barony {
	return c.俱卢伐底Kuruvatti
}
    
func (c *巴南巴西BanavasiCounty) BLakkundi拉昆迪() feud.Barony {
	return c.拉昆迪Lakkundi
}
    
func (c *巴南巴西BanavasiCounty) BSringeri斯灵盖里() feud.Barony {
	return c.斯灵盖里Sringeri
}
    
func (c *巴南巴西BanavasiCounty) BVaijayanti吠阇演帝() feud.Barony {
	return c.吠阇演帝Vaijayanti
}
    
func (c *巴南巴西BanavasiCounty) BVankapura房迦补罗() feud.Barony {
	return c.房迦补罗Vankapura
}
    
var CBanavasi巴南巴西 BanavasiCounty = &巴南巴西BanavasiCounty{}

func init() {
	f := CBanavasi巴南巴西.(*巴南巴西BanavasiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1140",
		Title:     "banavasi",
		TitleName: "巴南巴西",
		TitleCode: "c_banavasi",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴里加维Balligavi = BBalligavi巴里加维
	f.巴里加维Balligavi.SetParent(f)
	
	f.旃陀罗古提Candragutti = BCandragutti旃陀罗古提
	f.旃陀罗古提Candragutti.SetParent(f)
	
	f.汉加尔Hangal = BHangal汉加尔
	f.汉加尔Hangal.SetParent(f)
	
	f.伊凯里Ikkeri = BIkkeri伊凯里
	f.伊凯里Ikkeri.SetParent(f)
	
	f.俱卢伐底Kuruvatti = BKuruvatti俱卢伐底
	f.俱卢伐底Kuruvatti.SetParent(f)
	
	f.拉昆迪Lakkundi = BLakkundi拉昆迪
	f.拉昆迪Lakkundi.SetParent(f)
	
	f.斯灵盖里Sringeri = BSringeri斯灵盖里
	f.斯灵盖里Sringeri.SetParent(f)
	
	f.吠阇演帝Vaijayanti = BVaijayanti吠阇演帝
	f.吠阇演帝Vaijayanti.SetParent(f)
	
	f.房迦补罗Vankapura = BVankapura房迦补罗
	f.房迦补罗Vankapura.SetParent(f)
	
}
