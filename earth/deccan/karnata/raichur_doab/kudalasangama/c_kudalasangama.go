package kudalasangama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KudalasangamaCounty interface {
    feud.County
    BAnegundi阿内古尼提() 	feud.Barony
    BBagali巴加利() 	feud.Barony
    BDambal敦巴尔() 	feud.Barony
    BKoppam科潘() 	feud.Barony
    BKudalasangama俱荼罗僧伽摩() 	feud.Barony
    BKuknur鸠努尔() 	feud.Barony
    BMusangi穆桑戈() 	feud.Barony
}

type 俱荼罗僧伽摩KudalasangamaCounty struct {
	feud.BaseCounty
	阿内古尼提Anegundi 	feud.Barony
	巴加利Bagali 	feud.Barony
	敦巴尔Dambal 	feud.Barony
	科潘Koppam 	feud.Barony
	俱荼罗僧伽摩Kudalasangama 	feud.Barony
	鸠努尔Kuknur 	feud.Barony
	穆桑戈Musangi 	feud.Barony
}

func (c *俱荼罗僧伽摩KudalasangamaCounty) BAnegundi阿内古尼提() feud.Barony {
	return c.阿内古尼提Anegundi
}
    
func (c *俱荼罗僧伽摩KudalasangamaCounty) BBagali巴加利() feud.Barony {
	return c.巴加利Bagali
}
    
func (c *俱荼罗僧伽摩KudalasangamaCounty) BDambal敦巴尔() feud.Barony {
	return c.敦巴尔Dambal
}
    
func (c *俱荼罗僧伽摩KudalasangamaCounty) BKoppam科潘() feud.Barony {
	return c.科潘Koppam
}
    
func (c *俱荼罗僧伽摩KudalasangamaCounty) BKudalasangama俱荼罗僧伽摩() feud.Barony {
	return c.俱荼罗僧伽摩Kudalasangama
}
    
func (c *俱荼罗僧伽摩KudalasangamaCounty) BKuknur鸠努尔() feud.Barony {
	return c.鸠努尔Kuknur
}
    
func (c *俱荼罗僧伽摩KudalasangamaCounty) BMusangi穆桑戈() feud.Barony {
	return c.穆桑戈Musangi
}
    
var CKudalasangama俱荼罗僧伽摩 KudalasangamaCounty = &俱荼罗僧伽摩KudalasangamaCounty{}

func init() {
	f := CKudalasangama俱荼罗僧伽摩.(*俱荼罗僧伽摩KudalasangamaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1202",
		Title:     "kudalasangama",
		TitleName: "俱荼罗僧伽摩",
		TitleCode: "c_kudalasangama",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿内古尼提Anegundi = BAnegundi阿内古尼提
	f.阿内古尼提Anegundi.SetParent(f)
	
	f.巴加利Bagali = BBagali巴加利
	f.巴加利Bagali.SetParent(f)
	
	f.敦巴尔Dambal = BDambal敦巴尔
	f.敦巴尔Dambal.SetParent(f)
	
	f.科潘Koppam = BKoppam科潘
	f.科潘Koppam.SetParent(f)
	
	f.俱荼罗僧伽摩Kudalasangama = BKudalasangama俱荼罗僧伽摩
	f.俱荼罗僧伽摩Kudalasangama.SetParent(f)
	
	f.鸠努尔Kuknur = BKuknur鸠努尔
	f.鸠努尔Kuknur.SetParent(f)
	
	f.穆桑戈Musangi = BMusangi穆桑戈
	f.穆桑戈Musangi.SetParent(f)
	
}
