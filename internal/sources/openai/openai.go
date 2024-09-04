package openai

import (
	"fmt"

	"github.com/medali/ramo-server/internal/app"
)



func HelperFunc(infos app.HelperInput) string{



	fmt.Println(infos.BusinessName)

	return ""
	
}
