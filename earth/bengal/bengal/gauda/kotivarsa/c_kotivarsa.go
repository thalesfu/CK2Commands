package kotivarsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KotivarsaCounty interface {
    feud.County
    BBangarh槃姞利呬() 	feud.Barony
    BDevkot提婆拘吒() 	feud.Barony
    BJagaddala阇伽陀罗() 	feud.Barony
    BKorat拘剌吒() 	feud.Barony
    BPurena富梨那() 	feud.Barony
    BRamavati罗摩伐底() 	feud.Barony
    BSiliguri斯梨古利() 	feud.Barony
}

type 俱知勃里沙KotivarsaCounty struct {
	feud.BaseCounty
	槃姞利呬Bangarh 	feud.Barony
	提婆拘吒Devkot 	feud.Barony
	阇伽陀罗Jagaddala 	feud.Barony
	拘剌吒Korat 	feud.Barony
	富梨那Purena 	feud.Barony
	罗摩伐底Ramavati 	feud.Barony
	斯梨古利Siliguri 	feud.Barony
}

func (c *俱知勃里沙KotivarsaCounty) BBangarh槃姞利呬() feud.Barony {
	return c.槃姞利呬Bangarh
}
    
func (c *俱知勃里沙KotivarsaCounty) BDevkot提婆拘吒() feud.Barony {
	return c.提婆拘吒Devkot
}
    
func (c *俱知勃里沙KotivarsaCounty) BJagaddala阇伽陀罗() feud.Barony {
	return c.阇伽陀罗Jagaddala
}
    
func (c *俱知勃里沙KotivarsaCounty) BKorat拘剌吒() feud.Barony {
	return c.拘剌吒Korat
}
    
func (c *俱知勃里沙KotivarsaCounty) BPurena富梨那() feud.Barony {
	return c.富梨那Purena
}
    
func (c *俱知勃里沙KotivarsaCounty) BRamavati罗摩伐底() feud.Barony {
	return c.罗摩伐底Ramavati
}
    
func (c *俱知勃里沙KotivarsaCounty) BSiliguri斯梨古利() feud.Barony {
	return c.斯梨古利Siliguri
}
    
var CKotivarsa俱知勃里沙 KotivarsaCounty = &俱知勃里沙KotivarsaCounty{}

func init() {
	f := CKotivarsa俱知勃里沙.(*俱知勃里沙KotivarsaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1153",
		Title:     "kotivarsa",
		TitleName: "俱知勃里沙",
		TitleCode: "c_kotivarsa",
		Baronies:  map[string]feud.Barony{},
	}

	f.槃姞利呬Bangarh = BBangarh槃姞利呬
	f.槃姞利呬Bangarh.SetParent(f)
	
	f.提婆拘吒Devkot = BDevkot提婆拘吒
	f.提婆拘吒Devkot.SetParent(f)
	
	f.阇伽陀罗Jagaddala = BJagaddala阇伽陀罗
	f.阇伽陀罗Jagaddala.SetParent(f)
	
	f.拘剌吒Korat = BKorat拘剌吒
	f.拘剌吒Korat.SetParent(f)
	
	f.富梨那Purena = BPurena富梨那
	f.富梨那Purena.SetParent(f)
	
	f.罗摩伐底Ramavati = BRamavati罗摩伐底
	f.罗摩伐底Ramavati.SetParent(f)
	
	f.斯梨古利Siliguri = BSiliguri斯梨古利
	f.斯梨古利Siliguri.SetParent(f)
	
}
