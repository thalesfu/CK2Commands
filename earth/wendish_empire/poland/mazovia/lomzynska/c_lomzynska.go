package lomzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LomzynskaCounty interface {
    feud.County
    BKolno科尔诺() 	feud.Barony
    BLomza沃姆扎() 	feud.Barony
    BNur努尔() 	feud.Barony
    BOstroleka奥斯特罗文卡() 	feud.Barony
    BRadzilow拉济武夫() 	feud.Barony
    BSwieck希温茨克() 	feud.Barony
    BWizna维兹纳() 	feud.Barony
}

type 沃姆扎LomzynskaCounty struct {
	feud.BaseCounty
	科尔诺Kolno 	feud.Barony
	沃姆扎Lomza 	feud.Barony
	努尔Nur 	feud.Barony
	奥斯特罗文卡Ostroleka 	feud.Barony
	拉济武夫Radzilow 	feud.Barony
	希温茨克Swieck 	feud.Barony
	维兹纳Wizna 	feud.Barony
}

func (c *沃姆扎LomzynskaCounty) BKolno科尔诺() feud.Barony {
	return c.科尔诺Kolno
}
    
func (c *沃姆扎LomzynskaCounty) BLomza沃姆扎() feud.Barony {
	return c.沃姆扎Lomza
}
    
func (c *沃姆扎LomzynskaCounty) BNur努尔() feud.Barony {
	return c.努尔Nur
}
    
func (c *沃姆扎LomzynskaCounty) BOstroleka奥斯特罗文卡() feud.Barony {
	return c.奥斯特罗文卡Ostroleka
}
    
func (c *沃姆扎LomzynskaCounty) BRadzilow拉济武夫() feud.Barony {
	return c.拉济武夫Radzilow
}
    
func (c *沃姆扎LomzynskaCounty) BSwieck希温茨克() feud.Barony {
	return c.希温茨克Swieck
}
    
func (c *沃姆扎LomzynskaCounty) BWizna维兹纳() feud.Barony {
	return c.维兹纳Wizna
}
    
var CLomzynska沃姆扎 LomzynskaCounty = &沃姆扎LomzynskaCounty{}

func init() {
	f := CLomzynska沃姆扎.(*沃姆扎LomzynskaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1583",
		Title:     "lomzynska",
		TitleName: "沃姆扎",
		TitleCode: "c_lomzynska",
		Baronies:  map[string]feud.Barony{},
	}

	f.科尔诺Kolno = BKolno科尔诺
	f.科尔诺Kolno.SetParent(f)
	
	f.沃姆扎Lomza = BLomza沃姆扎
	f.沃姆扎Lomza.SetParent(f)
	
	f.努尔Nur = BNur努尔
	f.努尔Nur.SetParent(f)
	
	f.奥斯特罗文卡Ostroleka = BOstroleka奥斯特罗文卡
	f.奥斯特罗文卡Ostroleka.SetParent(f)
	
	f.拉济武夫Radzilow = BRadzilow拉济武夫
	f.拉济武夫Radzilow.SetParent(f)
	
	f.希温茨克Swieck = BSwieck希温茨克
	f.希温茨克Swieck.SetParent(f)
	
	f.维兹纳Wizna = BWizna维兹纳
	f.维兹纳Wizna.SetParent(f)
	
}
