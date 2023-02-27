package kolguyev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolguyevCounty interface {
    feud.County
    BBugryanka布格良卡() 	feud.Barony
    BGubistaya古比斯塔亚() 	feud.Barony
    BGusinaya古西纳亚() 	feud.Barony
    BKolguyev科尔古耶夫() 	feud.Barony
    BKonkina孔基纳() 	feud.Barony
    BPervaya佩尔瓦亚() 	feud.Barony
    BVelikaya韦利卡亚() 	feud.Barony
}

type 科尔古耶夫KolguyevCounty struct {
	feud.BaseCounty
	布格良卡Bugryanka 	feud.Barony
	古比斯塔亚Gubistaya 	feud.Barony
	古西纳亚Gusinaya 	feud.Barony
	科尔古耶夫Kolguyev 	feud.Barony
	孔基纳Konkina 	feud.Barony
	佩尔瓦亚Pervaya 	feud.Barony
	韦利卡亚Velikaya 	feud.Barony
}

func (c *科尔古耶夫KolguyevCounty) BBugryanka布格良卡() feud.Barony {
	return c.布格良卡Bugryanka
}
    
func (c *科尔古耶夫KolguyevCounty) BGubistaya古比斯塔亚() feud.Barony {
	return c.古比斯塔亚Gubistaya
}
    
func (c *科尔古耶夫KolguyevCounty) BGusinaya古西纳亚() feud.Barony {
	return c.古西纳亚Gusinaya
}
    
func (c *科尔古耶夫KolguyevCounty) BKolguyev科尔古耶夫() feud.Barony {
	return c.科尔古耶夫Kolguyev
}
    
func (c *科尔古耶夫KolguyevCounty) BKonkina孔基纳() feud.Barony {
	return c.孔基纳Konkina
}
    
func (c *科尔古耶夫KolguyevCounty) BPervaya佩尔瓦亚() feud.Barony {
	return c.佩尔瓦亚Pervaya
}
    
func (c *科尔古耶夫KolguyevCounty) BVelikaya韦利卡亚() feud.Barony {
	return c.韦利卡亚Velikaya
}
    
var CKolguyev科尔古耶夫 KolguyevCounty = &科尔古耶夫KolguyevCounty{}

func init() {
	f := CKolguyev科尔古耶夫.(*科尔古耶夫KolguyevCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1829",
		Title:     "kolguyev",
		TitleName: "科尔古耶夫",
		TitleCode: "c_kolguyev",
		Baronies:  map[string]feud.Barony{},
	}

	f.布格良卡Bugryanka = BBugryanka布格良卡
	f.布格良卡Bugryanka.SetParent(f)
	
	f.古比斯塔亚Gubistaya = BGubistaya古比斯塔亚
	f.古比斯塔亚Gubistaya.SetParent(f)
	
	f.古西纳亚Gusinaya = BGusinaya古西纳亚
	f.古西纳亚Gusinaya.SetParent(f)
	
	f.科尔古耶夫Kolguyev = BKolguyev科尔古耶夫
	f.科尔古耶夫Kolguyev.SetParent(f)
	
	f.孔基纳Konkina = BKonkina孔基纳
	f.孔基纳Konkina.SetParent(f)
	
	f.佩尔瓦亚Pervaya = BPervaya佩尔瓦亚
	f.佩尔瓦亚Pervaya.SetParent(f)
	
	f.韦利卡亚Velikaya = BVelikaya韦利卡亚
	f.韦利卡亚Velikaya.SetParent(f)
	
}
