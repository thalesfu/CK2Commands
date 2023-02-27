package wielunska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WielunskaCounty interface {
    feud.County
    BCzarnozyly恰尔诺日维() 	feud.Barony
    BGrabow格拉布夫() 	feud.Barony
    BOsjakow奥夏库夫() 	feud.Barony
    BOstrowek奥斯特鲁韦克() 	feud.Barony
    BOstrzeszow奥斯切舒夫() 	feud.Barony
    BWielun维隆() 	feud.Barony
    BWierzchlas维日赫拉斯() 	feud.Barony
}

type 维隆WielunskaCounty struct {
	feud.BaseCounty
	恰尔诺日维Czarnozyly 	feud.Barony
	格拉布夫Grabow 	feud.Barony
	奥夏库夫Osjakow 	feud.Barony
	奥斯特鲁韦克Ostrowek 	feud.Barony
	奥斯切舒夫Ostrzeszow 	feud.Barony
	维隆Wielun 	feud.Barony
	维日赫拉斯Wierzchlas 	feud.Barony
}

func (c *维隆WielunskaCounty) BCzarnozyly恰尔诺日维() feud.Barony {
	return c.恰尔诺日维Czarnozyly
}
    
func (c *维隆WielunskaCounty) BGrabow格拉布夫() feud.Barony {
	return c.格拉布夫Grabow
}
    
func (c *维隆WielunskaCounty) BOsjakow奥夏库夫() feud.Barony {
	return c.奥夏库夫Osjakow
}
    
func (c *维隆WielunskaCounty) BOstrowek奥斯特鲁韦克() feud.Barony {
	return c.奥斯特鲁韦克Ostrowek
}
    
func (c *维隆WielunskaCounty) BOstrzeszow奥斯切舒夫() feud.Barony {
	return c.奥斯切舒夫Ostrzeszow
}
    
func (c *维隆WielunskaCounty) BWielun维隆() feud.Barony {
	return c.维隆Wielun
}
    
func (c *维隆WielunskaCounty) BWierzchlas维日赫拉斯() feud.Barony {
	return c.维日赫拉斯Wierzchlas
}
    
var CWielunska维隆 WielunskaCounty = &维隆WielunskaCounty{}

func init() {
	f := CWielunska维隆.(*维隆WielunskaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1586",
		Title:     "wielunska",
		TitleName: "维隆",
		TitleCode: "c_wielunska",
		Baronies:  map[string]feud.Barony{},
	}

	f.恰尔诺日维Czarnozyly = BCzarnozyly恰尔诺日维
	f.恰尔诺日维Czarnozyly.SetParent(f)
	
	f.格拉布夫Grabow = BGrabow格拉布夫
	f.格拉布夫Grabow.SetParent(f)
	
	f.奥夏库夫Osjakow = BOsjakow奥夏库夫
	f.奥夏库夫Osjakow.SetParent(f)
	
	f.奥斯特鲁韦克Ostrowek = BOstrowek奥斯特鲁韦克
	f.奥斯特鲁韦克Ostrowek.SetParent(f)
	
	f.奥斯切舒夫Ostrzeszow = BOstrzeszow奥斯切舒夫
	f.奥斯切舒夫Ostrzeszow.SetParent(f)
	
	f.维隆Wielun = BWielun维隆
	f.维隆Wielun.SetParent(f)
	
	f.维日赫拉斯Wierzchlas = BWierzchlas维日赫拉斯
	f.维日赫拉斯Wierzchlas.SetParent(f)
	
}
