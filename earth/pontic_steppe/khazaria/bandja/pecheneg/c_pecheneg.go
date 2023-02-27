package pecheneg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PechenegCounty interface {
    feud.County
    BBandja班贾() 	feud.Barony
    BKargala卡尔加拉() 	feud.Barony
    BKhutorka胡托尔卡() 	feud.Barony
    BKichkas基奇卡斯() 	feud.Barony
    BSadovy萨多维() 	feud.Barony
    BSaraktash萨拉克塔什() 	feud.Barony
    BTolkachi托尔卡奇() 	feud.Barony
}

type 班贾PechenegCounty struct {
	feud.BaseCounty
	班贾Bandja 	feud.Barony
	卡尔加拉Kargala 	feud.Barony
	胡托尔卡Khutorka 	feud.Barony
	基奇卡斯Kichkas 	feud.Barony
	萨多维Sadovy 	feud.Barony
	萨拉克塔什Saraktash 	feud.Barony
	托尔卡奇Tolkachi 	feud.Barony
}

func (c *班贾PechenegCounty) BBandja班贾() feud.Barony {
	return c.班贾Bandja
}
    
func (c *班贾PechenegCounty) BKargala卡尔加拉() feud.Barony {
	return c.卡尔加拉Kargala
}
    
func (c *班贾PechenegCounty) BKhutorka胡托尔卡() feud.Barony {
	return c.胡托尔卡Khutorka
}
    
func (c *班贾PechenegCounty) BKichkas基奇卡斯() feud.Barony {
	return c.基奇卡斯Kichkas
}
    
func (c *班贾PechenegCounty) BSadovy萨多维() feud.Barony {
	return c.萨多维Sadovy
}
    
func (c *班贾PechenegCounty) BSaraktash萨拉克塔什() feud.Barony {
	return c.萨拉克塔什Saraktash
}
    
func (c *班贾PechenegCounty) BTolkachi托尔卡奇() feud.Barony {
	return c.托尔卡奇Tolkachi
}
    
var CPecheneg班贾 PechenegCounty = &班贾PechenegCounty{}

func init() {
	f := CPecheneg班贾.(*班贾PechenegCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "616",
		Title:     "pecheneg",
		TitleName: "班贾",
		TitleCode: "c_pecheneg",
		Baronies:  map[string]feud.Barony{},
	}

	f.班贾Bandja = BBandja班贾
	f.班贾Bandja.SetParent(f)
	
	f.卡尔加拉Kargala = BKargala卡尔加拉
	f.卡尔加拉Kargala.SetParent(f)
	
	f.胡托尔卡Khutorka = BKhutorka胡托尔卡
	f.胡托尔卡Khutorka.SetParent(f)
	
	f.基奇卡斯Kichkas = BKichkas基奇卡斯
	f.基奇卡斯Kichkas.SetParent(f)
	
	f.萨多维Sadovy = BSadovy萨多维
	f.萨多维Sadovy.SetParent(f)
	
	f.萨拉克塔什Saraktash = BSaraktash萨拉克塔什
	f.萨拉克塔什Saraktash.SetParent(f)
	
	f.托尔卡奇Tolkachi = BTolkachi托尔卡奇
	f.托尔卡奇Tolkachi.SetParent(f)
	
}
