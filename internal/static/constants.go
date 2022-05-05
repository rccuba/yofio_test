package static

const (
	//static
	CONFIG_FILE_NAME string = "config"
	CONFIG_FILE_TYPE string = "json"
	LOG_FILE         string = "info.log"
	APP_PORT         string = "APP_PORT"
	APP_HOST         string = "APP_HOST"
	MONGO_HOST       string = "MONGO_HOST"
	MONGO_DATABASE   string = "MONGO_DATABASE"
	MONGO_PORT       string = "MONGO_PORT"
	MONGO_USER       string = "MONGO_USER"
	MONGO_PASSWORD   string = "MONGO_PASSWORD"
	//messages
	MsgResponseStartProcess       string = "Start process -->   "
	MsgResponseStartApplication   string = "Starting the application...."
	MsgResponseStartError         string = "ERROR LEYENDO ARCHIVO DE CONFIGURACION"
	MsgResponseStartingNow        string = "STARTING API..."
	MsgResponseStartingEndpoints  string = "STARTING ENDPOINTS..."
	MsgResponseStartingRoutes     string = "STARTING ROUTES..."
	MsgResponseStartingSwagger    string = "STARTING SWAGGER..."
	MsgResponseConnectedMongoDB   string = "Connected to MongoDB!"
	MsgResponseObjectExists       string = "OBJETO EXISTENTE"
	MsgResponseServerErrorWrongID string = "ERROR DE SERVIDOR: ID CON FORMATO DE VALOR INCORRECTO"
	MsgResponseServerErrorNoID    string = "ERROR DE SERVIDOR: ID INEXISTENTE"
	MsgResponseServerErrorNoData  string = "ERROR DE SERVIDOR: SIN DATOS EXISTENTES"
	MsgApiRestTitle               string = "YoFio - Backend Golang - Prueba Técnica"
	MsgApiRestDescription         string = "Calcular las posibles cantidades de créditos de\n$300, $500 y $700 que podemos otorgar con el \ntotal de la inversión. Si existe más de una\nopción podrías seleccionar cualquiera de ellas.\n\nAutor: Roberto C. Gonzalez Reyes"
	MsgApiRestVersion1            string = "1.0"
	MsgTestEXPECTED               string = "EXPECTED"
	MsgSuccessfully               string = "OPERACION SATISFACTORIA"
	MsgErrorOperation             string = "ERROR EN OPERACION"
	MsgUndeliveredAmount          string = "NO ES POSIBLE ENTREGAR EL CREDITO PARA EL MONTO"
	//message middleware
	MsgInvalidInvestment        string = "EL VALOR PROPORCIONADO ES INVALIDO"
	MsgUnauthorizatedInvestment string = "EL VALOR PROPORCIONADO DEBE SER UN MULTIPLO DE 100"
	//message test
	MsgTestValidValueInvestment     string = "PROBANDO VALOR VALIDO DE INVERSION"
	MsgTestUndeliveredAmount        string = "PROBANDO VALOR INVALIDO (MENOR QUE 300) DE INVERSION"
	MsgTestUnauthorizatedInvestment string = "PROBANDO VALOR NO AUTORIZADO (NO MULTIPLO DE 100) DE INVERSION"
	MsgTestUnsupportedHTTPMethod    string = "PROBANDO METODO HTTP NO SOPORTADO"
	MsgTestStatistics               string = "PROBANDO ESTADISTICAS"
	//URLs
	URLStartingNow      string = "/"
	URLApi              string = "/api"
	URLCreditAssignment string = "/credit-assignment"
	URLStatistics       string = "/statistics"
	//types responses
	ERROR   string = "ERROR"
	SUCCESS string = "SUCCESS"
	TEST    string = "TEST"
	OPTIONS string = "OPTIONS"
	//keyvals
	KeyType    string = "type"
	KeyURL     string = "URL"
	KeyMessage string = "message"
	KeyTs      string = "ts"
	KeyCaller  string = "caller"
	//schemas
	SchemaHttp string = "http"
	//misc
	ValueEmpty string = ""
	//collectios
	CollectionCreditAssignment string = "credit_assignment"
	//fields
	FieldInvestment string = "investment"
	//sorts
	SortAsc string = ""
)
