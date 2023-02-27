package bengal

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/gauda"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/nadia"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/suhma"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/vanga"
	"github.com/thalesfu/CK2Commands/earth/bengal/bengal/varendra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BengalKingdom interface {
    feud.Kingdom
    DGauda乔荼() 	gauda.GaudaDuke
    DNadia那地耶() 	nadia.NadiaDuke
    DSuhma苏吸摩() 	suhma.SuhmaDuke
    DVanga旁伽() 	vanga.VangaDuke
    DVarendra婆连陀罗() 	varendra.VarendraDuke
}

type 旁伽罗BengalKingdom struct {
	feud.BaseKingdom
	乔荼Gauda 	gauda.GaudaDuke
	那地耶Nadia 	nadia.NadiaDuke
	苏吸摩Suhma 	suhma.SuhmaDuke
	旁伽Vanga 	vanga.VangaDuke
	婆连陀罗Varendra 	varendra.VarendraDuke
}

func (k *旁伽罗BengalKingdom) DGauda乔荼() gauda.GaudaDuke {
	return k.乔荼Gauda
}
    
func (k *旁伽罗BengalKingdom) DNadia那地耶() nadia.NadiaDuke {
	return k.那地耶Nadia
}
    
func (k *旁伽罗BengalKingdom) DSuhma苏吸摩() suhma.SuhmaDuke {
	return k.苏吸摩Suhma
}
    
func (k *旁伽罗BengalKingdom) DVanga旁伽() vanga.VangaDuke {
	return k.旁伽Vanga
}
    
func (k *旁伽罗BengalKingdom) DVarendra婆连陀罗() varendra.VarendraDuke {
	return k.婆连陀罗Varendra
}
    
var KBengal旁伽罗 BengalKingdom = &旁伽罗BengalKingdom{}

func init() {
	f := KBengal旁伽罗.(*旁伽罗BengalKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "bengal",
		TitleName: "旁伽罗",
		TitleCode: "k_bengal",
		Dukes:  map[string]feud.Duke{},
	}

	f.乔荼Gauda = gauda.DGauda乔荼
	f.乔荼Gauda.SetParent(f)
	
	f.那地耶Nadia = nadia.DNadia那地耶
	f.那地耶Nadia.SetParent(f)
	
	f.苏吸摩Suhma = suhma.DSuhma苏吸摩
	f.苏吸摩Suhma.SetParent(f)
	
	f.旁伽Vanga = vanga.DVanga旁伽
	f.旁伽Vanga.SetParent(f)
	
	f.婆连陀罗Varendra = varendra.DVarendra婆连陀罗
	f.婆连陀罗Varendra.SetParent(f)
	
}
