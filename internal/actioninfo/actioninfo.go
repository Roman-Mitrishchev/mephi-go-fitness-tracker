package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, datastring := range dataset {
		if err := dp.Parse(datastring); err != nil {
			log.Printf("ошибка разбора %q: %v", datastring, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка обработки %q: %v", datastring, err)
			continue
		}
		fmt.Println(info)
	}
}
