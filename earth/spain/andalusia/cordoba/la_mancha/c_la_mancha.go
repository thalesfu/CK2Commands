package la_mancha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type La_manchaCounty interface {
    feud.County
    BAlarcon阿拉扎() 	feud.Barony
    BBarrax巴拉克斯() 	feud.Barony
    BJorquera霍尔克拉() 	feud.Barony
    BLagineta拉希内塔() 	feud.Barony
    BLaroda拉罗达() 	feud.Barony
    BMunera穆内拉() 	feud.Barony
    BQuintanardelrey金塔纳尔德尔雷伊() 	feud.Barony
    BTarazona塔拉索纳() 	feud.Barony
}

type 拉曼查La_manchaCounty struct {
	feud.BaseCounty
	阿拉扎Alarcon 	feud.Barony
	巴拉克斯Barrax 	feud.Barony
	霍尔克拉Jorquera 	feud.Barony
	拉希内塔Lagineta 	feud.Barony
	拉罗达Laroda 	feud.Barony
	穆内拉Munera 	feud.Barony
	金塔纳尔德尔雷伊Quintanardelrey 	feud.Barony
	塔拉索纳Tarazona 	feud.Barony
}

func (c *拉曼查La_manchaCounty) BAlarcon阿拉扎() feud.Barony {
	return c.阿拉扎Alarcon
}
    
func (c *拉曼查La_manchaCounty) BBarrax巴拉克斯() feud.Barony {
	return c.巴拉克斯Barrax
}
    
func (c *拉曼查La_manchaCounty) BJorquera霍尔克拉() feud.Barony {
	return c.霍尔克拉Jorquera
}
    
func (c *拉曼查La_manchaCounty) BLagineta拉希内塔() feud.Barony {
	return c.拉希内塔Lagineta
}
    
func (c *拉曼查La_manchaCounty) BLaroda拉罗达() feud.Barony {
	return c.拉罗达Laroda
}
    
func (c *拉曼查La_manchaCounty) BMunera穆内拉() feud.Barony {
	return c.穆内拉Munera
}
    
func (c *拉曼查La_manchaCounty) BQuintanardelrey金塔纳尔德尔雷伊() feud.Barony {
	return c.金塔纳尔德尔雷伊Quintanardelrey
}
    
func (c *拉曼查La_manchaCounty) BTarazona塔拉索纳() feud.Barony {
	return c.塔拉索纳Tarazona
}
    
var CLa_mancha拉曼查 La_manchaCounty = &拉曼查La_manchaCounty{}

func init() {
	f := CLa_mancha拉曼查.(*拉曼查La_manchaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "178",
		Title:     "la_mancha",
		TitleName: "拉曼查",
		TitleCode: "c_la_mancha",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉扎Alarcon = BAlarcon阿拉扎
	f.阿拉扎Alarcon.SetParent(f)
	
	f.巴拉克斯Barrax = BBarrax巴拉克斯
	f.巴拉克斯Barrax.SetParent(f)
	
	f.霍尔克拉Jorquera = BJorquera霍尔克拉
	f.霍尔克拉Jorquera.SetParent(f)
	
	f.拉希内塔Lagineta = BLagineta拉希内塔
	f.拉希内塔Lagineta.SetParent(f)
	
	f.拉罗达Laroda = BLaroda拉罗达
	f.拉罗达Laroda.SetParent(f)
	
	f.穆内拉Munera = BMunera穆内拉
	f.穆内拉Munera.SetParent(f)
	
	f.金塔纳尔德尔雷伊Quintanardelrey = BQuintanardelrey金塔纳尔德尔雷伊
	f.金塔纳尔德尔雷伊Quintanardelrey.SetParent(f)
	
	f.塔拉索纳Tarazona = BTarazona塔拉索纳
	f.塔拉索纳Tarazona.SetParent(f)
	
}
