package bern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BernCounty interface {
    feud.County
    BBern伯尔尼() 	feud.Barony
    BBiel比尔() 	feud.Barony
    BFribourg弗里堡() 	feud.Barony
    BLuzern卢塞恩() 	feud.Barony
    BMurten穆尔滕() 	feud.Barony
    BSursee苏尔塞() 	feud.Barony
    BThun图恩() 	feud.Barony
    BUnterseen下森() 	feud.Barony
}

type 伯尔尼BernCounty struct {
	feud.BaseCounty
	伯尔尼Bern 	feud.Barony
	比尔Biel 	feud.Barony
	弗里堡Fribourg 	feud.Barony
	卢塞恩Luzern 	feud.Barony
	穆尔滕Murten 	feud.Barony
	苏尔塞Sursee 	feud.Barony
	图恩Thun 	feud.Barony
	下森Unterseen 	feud.Barony
}

func (c *伯尔尼BernCounty) BBern伯尔尼() feud.Barony {
	return c.伯尔尼Bern
}
    
func (c *伯尔尼BernCounty) BBiel比尔() feud.Barony {
	return c.比尔Biel
}
    
func (c *伯尔尼BernCounty) BFribourg弗里堡() feud.Barony {
	return c.弗里堡Fribourg
}
    
func (c *伯尔尼BernCounty) BLuzern卢塞恩() feud.Barony {
	return c.卢塞恩Luzern
}
    
func (c *伯尔尼BernCounty) BMurten穆尔滕() feud.Barony {
	return c.穆尔滕Murten
}
    
func (c *伯尔尼BernCounty) BSursee苏尔塞() feud.Barony {
	return c.苏尔塞Sursee
}
    
func (c *伯尔尼BernCounty) BThun图恩() feud.Barony {
	return c.图恩Thun
}
    
func (c *伯尔尼BernCounty) BUnterseen下森() feud.Barony {
	return c.下森Unterseen
}
    
var CBern伯尔尼 BernCounty = &伯尔尼BernCounty{}

func init() {
	f := CBern伯尔尼.(*伯尔尼BernCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "244",
		Title:     "bern",
		TitleName: "伯尔尼",
		TitleCode: "c_bern",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯尔尼Bern = BBern伯尔尼
	f.伯尔尼Bern.SetParent(f)
	
	f.比尔Biel = BBiel比尔
	f.比尔Biel.SetParent(f)
	
	f.弗里堡Fribourg = BFribourg弗里堡
	f.弗里堡Fribourg.SetParent(f)
	
	f.卢塞恩Luzern = BLuzern卢塞恩
	f.卢塞恩Luzern.SetParent(f)
	
	f.穆尔滕Murten = BMurten穆尔滕
	f.穆尔滕Murten.SetParent(f)
	
	f.苏尔塞Sursee = BSursee苏尔塞
	f.苏尔塞Sursee.SetParent(f)
	
	f.图恩Thun = BThun图恩
	f.图恩Thun.SetParent(f)
	
	f.下森Unterseen = BUnterseen下森
	f.下森Unterseen.SetParent(f)
	
}
