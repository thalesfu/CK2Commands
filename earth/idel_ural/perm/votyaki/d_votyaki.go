package votyaki

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/votyaki/keltma"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/votyaki/votyaki"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VotyakiDuke interface {
    feud.Duke
    CKeltma克利特马() 	keltma.KeltmaCounty
    CVotyaki沃佳克() 	votyaki.VotyakiCounty
}

type 沃佳克VotyakiDuke struct {
	feud.BaseDuke
	克利特马Keltma 	keltma.KeltmaCounty
	沃佳克Votyaki 	votyaki.VotyakiCounty
}

func (d *沃佳克VotyakiDuke) CKeltma克利特马() keltma.KeltmaCounty {
	return d.克利特马Keltma
}
    
func (d *沃佳克VotyakiDuke) CVotyaki沃佳克() votyaki.VotyakiCounty {
	return d.沃佳克Votyaki
}
    
var DVotyaki沃佳克 VotyakiDuke = &沃佳克VotyakiDuke{}

func init() {
	f := DVotyaki沃佳克.(*沃佳克VotyakiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "votyaki",
		TitleName: "沃佳克",
		TitleCode: "d_votyaki",
		Counties:  map[string]feud.County{},
	}

	f.克利特马Keltma = keltma.CKeltma克利特马
	f.克利特马Keltma.SetParent(f)
	
	f.沃佳克Votyaki = votyaki.CVotyaki沃佳克
	f.沃佳克Votyaki.SetParent(f)
	
}
