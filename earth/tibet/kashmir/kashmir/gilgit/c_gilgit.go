package gilgit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GilgitCounty interface {
    feud.County
    BBarjangle巴尔姜戈() 	feud.Barony
    BChatorkhand查陀尔坎德() 	feud.Barony
    BChitral吉德拉尔() 	feud.Barony
    BDanyor丹约尔() 	feud.Barony
    BGilgit吉尔吉特() 	feud.Barony
    BMinawar弥那伐尔() 	feud.Barony
    BPahot波呼特() 	feud.Barony
}

type 吉尔吉特GilgitCounty struct {
	feud.BaseCounty
	巴尔姜戈Barjangle 	feud.Barony
	查陀尔坎德Chatorkhand 	feud.Barony
	吉德拉尔Chitral 	feud.Barony
	丹约尔Danyor 	feud.Barony
	吉尔吉特Gilgit 	feud.Barony
	弥那伐尔Minawar 	feud.Barony
	波呼特Pahot 	feud.Barony
}

func (c *吉尔吉特GilgitCounty) BBarjangle巴尔姜戈() feud.Barony {
	return c.巴尔姜戈Barjangle
}
    
func (c *吉尔吉特GilgitCounty) BChatorkhand查陀尔坎德() feud.Barony {
	return c.查陀尔坎德Chatorkhand
}
    
func (c *吉尔吉特GilgitCounty) BChitral吉德拉尔() feud.Barony {
	return c.吉德拉尔Chitral
}
    
func (c *吉尔吉特GilgitCounty) BDanyor丹约尔() feud.Barony {
	return c.丹约尔Danyor
}
    
func (c *吉尔吉特GilgitCounty) BGilgit吉尔吉特() feud.Barony {
	return c.吉尔吉特Gilgit
}
    
func (c *吉尔吉特GilgitCounty) BMinawar弥那伐尔() feud.Barony {
	return c.弥那伐尔Minawar
}
    
func (c *吉尔吉特GilgitCounty) BPahot波呼特() feud.Barony {
	return c.波呼特Pahot
}
    
var CGilgit吉尔吉特 GilgitCounty = &吉尔吉特GilgitCounty{}

func init() {
	f := CGilgit吉尔吉特.(*吉尔吉特GilgitCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1464",
		Title:     "gilgit",
		TitleName: "吉尔吉特",
		TitleCode: "c_gilgit",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔姜戈Barjangle = BBarjangle巴尔姜戈
	f.巴尔姜戈Barjangle.SetParent(f)
	
	f.查陀尔坎德Chatorkhand = BChatorkhand查陀尔坎德
	f.查陀尔坎德Chatorkhand.SetParent(f)
	
	f.吉德拉尔Chitral = BChitral吉德拉尔
	f.吉德拉尔Chitral.SetParent(f)
	
	f.丹约尔Danyor = BDanyor丹约尔
	f.丹约尔Danyor.SetParent(f)
	
	f.吉尔吉特Gilgit = BGilgit吉尔吉特
	f.吉尔吉特Gilgit.SetParent(f)
	
	f.弥那伐尔Minawar = BMinawar弥那伐尔
	f.弥那伐尔Minawar.SetParent(f)
	
	f.波呼特Pahot = BPahot波呼特
	f.波呼特Pahot.SetParent(f)
	
}
