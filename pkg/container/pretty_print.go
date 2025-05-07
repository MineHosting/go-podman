// Go-podman/pkg/container/pretty_print.go
package container

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func PrettyPrint(containerData any) {
	val := reflect.ValueOf(containerData)

	switch val.Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i).Interface()
			printSingleContainer(item)
		}
	case reflect.Struct:
		printSingleContainer(containerData)

	default:
		fmt.Printf("[PrintContainerInfo]: Tipo não suportado: %T\n", containerData)
	}
}

func printSingleContainer(data any) {
	container, ok := data.(Container)
	if !ok {
		fmt.Printf("[printSingleContainer]: Erro ao fazer type assertion: %T\n", data)
		return
	}

	jsonBytes, err := json.MarshalIndent(container, "", "  ")
	if err != nil {
		fmt.Printf("[printSingleContainer]: Falha ao serializar container: %v\n", err)
		return
	}

	fmt.Println(string(jsonBytes))
}
