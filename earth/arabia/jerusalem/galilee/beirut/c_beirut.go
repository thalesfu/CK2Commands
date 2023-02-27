package beirut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BeirutCounty interface {
    feud.County
    BAbualhasan阿布_哈桑() 	feud.Barony
    BBeirut贝鲁特() 	feud.Barony
    BBeitkfeya比克费亚() 	feud.Barony
    BCavedetyron泰隆窟() 	feud.Barony
    BJournie尤里涅() 	feud.Barony
    BMashgarah迈什盖拉() 	feud.Barony
    BSarepta撒勒法() 	feud.Barony
    BSidon西顿() 	feud.Barony
}

type 贝鲁特BeirutCounty struct {
	feud.BaseCounty
	阿布_哈桑Abualhasan 	feud.Barony
	贝鲁特Beirut 	feud.Barony
	比克费亚Beitkfeya 	feud.Barony
	泰隆窟Cavedetyron 	feud.Barony
	尤里涅Journie 	feud.Barony
	迈什盖拉Mashgarah 	feud.Barony
	撒勒法Sarepta 	feud.Barony
	西顿Sidon 	feud.Barony
}

func (c *贝鲁特BeirutCounty) BAbualhasan阿布_哈桑() feud.Barony {
	return c.阿布_哈桑Abualhasan
}
    
func (c *贝鲁特BeirutCounty) BBeirut贝鲁特() feud.Barony {
	return c.贝鲁特Beirut
}
    
func (c *贝鲁特BeirutCounty) BBeitkfeya比克费亚() feud.Barony {
	return c.比克费亚Beitkfeya
}
    
func (c *贝鲁特BeirutCounty) BCavedetyron泰隆窟() feud.Barony {
	return c.泰隆窟Cavedetyron
}
    
func (c *贝鲁特BeirutCounty) BJournie尤里涅() feud.Barony {
	return c.尤里涅Journie
}
    
func (c *贝鲁特BeirutCounty) BMashgarah迈什盖拉() feud.Barony {
	return c.迈什盖拉Mashgarah
}
    
func (c *贝鲁特BeirutCounty) BSarepta撒勒法() feud.Barony {
	return c.撒勒法Sarepta
}
    
func (c *贝鲁特BeirutCounty) BSidon西顿() feud.Barony {
	return c.西顿Sidon
}
    
var CBeirut贝鲁特 BeirutCounty = &贝鲁特BeirutCounty{}

func init() {
	f := CBeirut贝鲁特.(*贝鲁特BeirutCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "770",
		Title:     "beirut",
		TitleName: "贝鲁特",
		TitleCode: "c_beirut",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布_哈桑Abualhasan = BAbualhasan阿布_哈桑
	f.阿布_哈桑Abualhasan.SetParent(f)
	
	f.贝鲁特Beirut = BBeirut贝鲁特
	f.贝鲁特Beirut.SetParent(f)
	
	f.比克费亚Beitkfeya = BBeitkfeya比克费亚
	f.比克费亚Beitkfeya.SetParent(f)
	
	f.泰隆窟Cavedetyron = BCavedetyron泰隆窟
	f.泰隆窟Cavedetyron.SetParent(f)
	
	f.尤里涅Journie = BJournie尤里涅
	f.尤里涅Journie.SetParent(f)
	
	f.迈什盖拉Mashgarah = BMashgarah迈什盖拉
	f.迈什盖拉Mashgarah.SetParent(f)
	
	f.撒勒法Sarepta = BSarepta撒勒法
	f.撒勒法Sarepta.SetParent(f)
	
	f.西顿Sidon = BSidon西顿
	f.西顿Sidon.SetParent(f)
	
}
