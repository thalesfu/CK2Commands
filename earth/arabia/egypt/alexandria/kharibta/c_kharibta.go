package kharibta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KharibtaCounty interface {
    feud.County
    BAtfih艾特菲赫() 	feud.Barony
    BKharibta海里卜塔() 	feud.Barony
    BMerimda迈里穆达() 	feud.Barony
    BNaucratis_egypt瑙克拉提斯() 	feud.Barony
    BNikiou尼基乌() 	feud.Barony
    BScetis斯刻提斯() 	feud.Barony
    BShatnuf沙特努夫() 	feud.Barony
}

type 海里卜塔KharibtaCounty struct {
	feud.BaseCounty
	艾特菲赫Atfih 	feud.Barony
	海里卜塔Kharibta 	feud.Barony
	迈里穆达Merimda 	feud.Barony
	瑙克拉提斯Naucratis_egypt 	feud.Barony
	尼基乌Nikiou 	feud.Barony
	斯刻提斯Scetis 	feud.Barony
	沙特努夫Shatnuf 	feud.Barony
}

func (c *海里卜塔KharibtaCounty) BAtfih艾特菲赫() feud.Barony {
	return c.艾特菲赫Atfih
}
    
func (c *海里卜塔KharibtaCounty) BKharibta海里卜塔() feud.Barony {
	return c.海里卜塔Kharibta
}
    
func (c *海里卜塔KharibtaCounty) BMerimda迈里穆达() feud.Barony {
	return c.迈里穆达Merimda
}
    
func (c *海里卜塔KharibtaCounty) BNaucratis_egypt瑙克拉提斯() feud.Barony {
	return c.瑙克拉提斯Naucratis_egypt
}
    
func (c *海里卜塔KharibtaCounty) BNikiou尼基乌() feud.Barony {
	return c.尼基乌Nikiou
}
    
func (c *海里卜塔KharibtaCounty) BScetis斯刻提斯() feud.Barony {
	return c.斯刻提斯Scetis
}
    
func (c *海里卜塔KharibtaCounty) BShatnuf沙特努夫() feud.Barony {
	return c.沙特努夫Shatnuf
}
    
var CKharibta海里卜塔 KharibtaCounty = &海里卜塔KharibtaCounty{}

func init() {
	f := CKharibta海里卜塔.(*海里卜塔KharibtaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2004",
		Title:     "kharibta",
		TitleName: "海里卜塔",
		TitleCode: "c_kharibta",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾特菲赫Atfih = BAtfih艾特菲赫
	f.艾特菲赫Atfih.SetParent(f)
	
	f.海里卜塔Kharibta = BKharibta海里卜塔
	f.海里卜塔Kharibta.SetParent(f)
	
	f.迈里穆达Merimda = BMerimda迈里穆达
	f.迈里穆达Merimda.SetParent(f)
	
	f.瑙克拉提斯Naucratis_egypt = BNaucratis_egypt瑙克拉提斯
	f.瑙克拉提斯Naucratis_egypt.SetParent(f)
	
	f.尼基乌Nikiou = BNikiou尼基乌
	f.尼基乌Nikiou.SetParent(f)
	
	f.斯刻提斯Scetis = BScetis斯刻提斯
	f.斯刻提斯Scetis.SetParent(f)
	
	f.沙特努夫Shatnuf = BShatnuf沙特努夫
	f.沙特努夫Shatnuf.SetParent(f)
	
}
