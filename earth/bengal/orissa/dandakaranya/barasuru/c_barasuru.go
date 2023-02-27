package barasuru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BarasuruCounty interface {
    feud.County
    BBarasuru婆罗须卢() 	feud.Barony
    BBarsoor柏索尔() 	feud.Barony
    BDantewada达恩泰瓦达() 	feud.Barony
    BKonokjora拘诺迦殊罗() 	feud.Barony
    BKorandh拘兰陀() 	feud.Barony
    BKothaguda拘多瞿荼() 	feud.Barony
    BManghech曼俱侯支() 	feud.Barony
}

type 婆罗须卢BarasuruCounty struct {
	feud.BaseCounty
	婆罗须卢Barasuru 	feud.Barony
	柏索尔Barsoor 	feud.Barony
	达恩泰瓦达Dantewada 	feud.Barony
	拘诺迦殊罗Konokjora 	feud.Barony
	拘兰陀Korandh 	feud.Barony
	拘多瞿荼Kothaguda 	feud.Barony
	曼俱侯支Manghech 	feud.Barony
}

func (c *婆罗须卢BarasuruCounty) BBarasuru婆罗须卢() feud.Barony {
	return c.婆罗须卢Barasuru
}
    
func (c *婆罗须卢BarasuruCounty) BBarsoor柏索尔() feud.Barony {
	return c.柏索尔Barsoor
}
    
func (c *婆罗须卢BarasuruCounty) BDantewada达恩泰瓦达() feud.Barony {
	return c.达恩泰瓦达Dantewada
}
    
func (c *婆罗须卢BarasuruCounty) BKonokjora拘诺迦殊罗() feud.Barony {
	return c.拘诺迦殊罗Konokjora
}
    
func (c *婆罗须卢BarasuruCounty) BKorandh拘兰陀() feud.Barony {
	return c.拘兰陀Korandh
}
    
func (c *婆罗须卢BarasuruCounty) BKothaguda拘多瞿荼() feud.Barony {
	return c.拘多瞿荼Kothaguda
}
    
func (c *婆罗须卢BarasuruCounty) BManghech曼俱侯支() feud.Barony {
	return c.曼俱侯支Manghech
}
    
var CBarasuru婆罗须卢 BarasuruCounty = &婆罗须卢BarasuruCounty{}

func init() {
	f := CBarasuru婆罗须卢.(*婆罗须卢BarasuruCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1252",
		Title:     "barasuru",
		TitleName: "婆罗须卢",
		TitleCode: "c_barasuru",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆罗须卢Barasuru = BBarasuru婆罗须卢
	f.婆罗须卢Barasuru.SetParent(f)
	
	f.柏索尔Barsoor = BBarsoor柏索尔
	f.柏索尔Barsoor.SetParent(f)
	
	f.达恩泰瓦达Dantewada = BDantewada达恩泰瓦达
	f.达恩泰瓦达Dantewada.SetParent(f)
	
	f.拘诺迦殊罗Konokjora = BKonokjora拘诺迦殊罗
	f.拘诺迦殊罗Konokjora.SetParent(f)
	
	f.拘兰陀Korandh = BKorandh拘兰陀
	f.拘兰陀Korandh.SetParent(f)
	
	f.拘多瞿荼Kothaguda = BKothaguda拘多瞿荼
	f.拘多瞿荼Kothaguda.SetParent(f)
	
	f.曼俱侯支Manghech = BManghech曼俱侯支
	f.曼俱侯支Manghech.SetParent(f)
	
}
