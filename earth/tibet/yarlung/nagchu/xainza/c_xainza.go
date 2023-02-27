package xainza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type XainzaCounty interface {
    feud.County
    BBaingoin班戈() 	feud.Barony
    BJakhyung佳琼() 	feud.Barony
    BMaryo马跃() 	feud.Barony
    BPukpa普保() 	feud.Barony
    BSiling色林() 	feud.Barony
    BXainza申扎() 	feud.Barony
    BZhago下过() 	feud.Barony
}

type 申扎XainzaCounty struct {
	feud.BaseCounty
	班戈Baingoin 	feud.Barony
	佳琼Jakhyung 	feud.Barony
	马跃Maryo 	feud.Barony
	普保Pukpa 	feud.Barony
	色林Siling 	feud.Barony
	申扎Xainza 	feud.Barony
	下过Zhago 	feud.Barony
}

func (c *申扎XainzaCounty) BBaingoin班戈() feud.Barony {
	return c.班戈Baingoin
}
    
func (c *申扎XainzaCounty) BJakhyung佳琼() feud.Barony {
	return c.佳琼Jakhyung
}
    
func (c *申扎XainzaCounty) BMaryo马跃() feud.Barony {
	return c.马跃Maryo
}
    
func (c *申扎XainzaCounty) BPukpa普保() feud.Barony {
	return c.普保Pukpa
}
    
func (c *申扎XainzaCounty) BSiling色林() feud.Barony {
	return c.色林Siling
}
    
func (c *申扎XainzaCounty) BXainza申扎() feud.Barony {
	return c.申扎Xainza
}
    
func (c *申扎XainzaCounty) BZhago下过() feud.Barony {
	return c.下过Zhago
}
    
var CXainza申扎 XainzaCounty = &申扎XainzaCounty{}

func init() {
	f := CXainza申扎.(*申扎XainzaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1559",
		Title:     "xainza",
		TitleName: "申扎",
		TitleCode: "c_xainza",
		Baronies:  map[string]feud.Barony{},
	}

	f.班戈Baingoin = BBaingoin班戈
	f.班戈Baingoin.SetParent(f)
	
	f.佳琼Jakhyung = BJakhyung佳琼
	f.佳琼Jakhyung.SetParent(f)
	
	f.马跃Maryo = BMaryo马跃
	f.马跃Maryo.SetParent(f)
	
	f.普保Pukpa = BPukpa普保
	f.普保Pukpa.SetParent(f)
	
	f.色林Siling = BSiling色林
	f.色林Siling.SetParent(f)
	
	f.申扎Xainza = BXainza申扎
	f.申扎Xainza.SetParent(f)
	
	f.下过Zhago = BZhago下过
	f.下过Zhago.SetParent(f)
	
}
