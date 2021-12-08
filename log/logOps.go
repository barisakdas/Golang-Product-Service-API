package log

import "log"

func LogJson(logType, controllerName, functionName, errorDetail string) {
	errorMessage := `{
						"Type":` + logType + `,
						"Controller_Name":` + controllerName + `
						"Function_Name":` + functionName + `,
						"Error_Detail":` + errorDetail + `
					}`
	log.Fatalln(errorMessage)
}
