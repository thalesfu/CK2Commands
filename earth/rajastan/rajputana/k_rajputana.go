package rajputana

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/ajmer"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/jangladesh"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/maru"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/medapata"
	"github.com/thalesfu/CK2Commands/earth/rajastan/rajputana/stravani"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RajputanaKingdom interface {
    feud.Kingdom
    DAjmer阿阇耶迷楼() 	ajmer.AjmerDuke
    DJangladesh常伽罗提舍() 	jangladesh.JangladeshDuke
    DMaru摩楼() 	maru.MaruDuke
    DMedapata迷陀波吒() 	medapata.MedapataDuke
    DStravani萨怛罗婆尼() 	stravani.StravaniDuke
}

type 罗阇弗多那RajputanaKingdom struct {
	feud.BaseKingdom
	阿阇耶迷楼Ajmer 	ajmer.AjmerDuke
	常伽罗提舍Jangladesh 	jangladesh.JangladeshDuke
	摩楼Maru 	maru.MaruDuke
	迷陀波吒Medapata 	medapata.MedapataDuke
	萨怛罗婆尼Stravani 	stravani.StravaniDuke
}

func (k *罗阇弗多那RajputanaKingdom) DAjmer阿阇耶迷楼() ajmer.AjmerDuke {
	return k.阿阇耶迷楼Ajmer
}
    
func (k *罗阇弗多那RajputanaKingdom) DJangladesh常伽罗提舍() jangladesh.JangladeshDuke {
	return k.常伽罗提舍Jangladesh
}
    
func (k *罗阇弗多那RajputanaKingdom) DMaru摩楼() maru.MaruDuke {
	return k.摩楼Maru
}
    
func (k *罗阇弗多那RajputanaKingdom) DMedapata迷陀波吒() medapata.MedapataDuke {
	return k.迷陀波吒Medapata
}
    
func (k *罗阇弗多那RajputanaKingdom) DStravani萨怛罗婆尼() stravani.StravaniDuke {
	return k.萨怛罗婆尼Stravani
}
    
var KRajputana罗阇弗多那 RajputanaKingdom = &罗阇弗多那RajputanaKingdom{}

func init() {
	f := KRajputana罗阇弗多那.(*罗阇弗多那RajputanaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "rajputana",
		TitleName: "罗阇弗多那",
		TitleCode: "k_rajputana",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿阇耶迷楼Ajmer = ajmer.DAjmer阿阇耶迷楼
	f.阿阇耶迷楼Ajmer.SetParent(f)
	
	f.常伽罗提舍Jangladesh = jangladesh.DJangladesh常伽罗提舍
	f.常伽罗提舍Jangladesh.SetParent(f)
	
	f.摩楼Maru = maru.DMaru摩楼
	f.摩楼Maru.SetParent(f)
	
	f.迷陀波吒Medapata = medapata.DMedapata迷陀波吒
	f.迷陀波吒Medapata.SetParent(f)
	
	f.萨怛罗婆尼Stravani = stravani.DStravani萨怛罗婆尼
	f.萨怛罗婆尼Stravani.SetParent(f)
	
}
