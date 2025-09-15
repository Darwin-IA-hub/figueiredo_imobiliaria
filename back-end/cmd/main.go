package main

import (
	"back-end/controllers"
	"back-end/database"
	"back-end/repository"
	"back-end/routes"
	"back-end/usecases"
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func CleanJSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			body, _ := io.ReadAll(c.Request.Body)
			body = cleanJSONBody(body) // (a mesma função que você já tem)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		c.Next() // segue para o handler real
	}
}

func cleanJSONBody(body []byte) []byte {
	// Remove caracteres invisíveis problemáticos
	removals := [][]byte{
		[]byte{0xC2, 0xA0},       // NO-BREAK SPACE
		[]byte{0xE2, 0x80, 0x8B}, // Zero Width Space
		[]byte{0xE2, 0x80, 0x8C}, // Zero Width Non-Joiner
		[]byte{0xE2, 0x80, 0x8D}, // Zero Width Joiner
		[]byte{0xE2, 0x80, 0x8E}, // Left-to-Right Mark
		[]byte{0xE2, 0x80, 0x8F}, // Right-to-Left Mark
		[]byte{0xEF, 0xBB, 0xBF}, // BOM UTF-8
	}
	for _, pattern := range removals {
		body = bytes.ReplaceAll(body, pattern, []byte{})
	}
	// Remove ASCII control chars: \r, \t, \v, \f (se necessário)
	body = bytes.ReplaceAll(body, []byte{0x0D}, []byte{}) // \r
	body = bytes.ReplaceAll(body, []byte{0x09}, []byte{}) // \t
	body = bytes.ReplaceAll(body, []byte{0x0B}, []byte{}) // \v
	body = bytes.ReplaceAll(body, []byte{0x0C}, []byte{}) // \f
	return body
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	dbConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	server.Use(CORSMiddleware())
	server.Use(CleanJSONMiddleware())

	routes.SetMessageRoutes(server, controllers.NewMessageController(usecases.NewMessageUseCases(repository.NewMessageRepository(dbConnection))))
	routes.SetUsuarioRoutes(server, controllers.NewUsuarioController(usecases.NewUsuarioUseCases(repository.NewUsuarioRepository(dbConnection))))
	routes.SetupClienteRoutes(server, controllers.NewClienteController(usecases.NewClienteUseCases(repository.NewClienteRepository(dbConnection))))
	routes.SetupContactRoutes(server, controllers.NewContactController(usecases.NewContactUseCases(repository.NewContactRepository(dbConnection))))
	routes.SetupLeadRoutes(server, controllers.NewLeadController(usecases.NewLeadUseCases(repository.NewLeadRepository(dbConnection))))

	routes.SetupFinanciamentoRoutes(server, controllers.NewFinanciamentoController(usecases.NewFinanciamentoUseCases(repository.NewFinanciamentoRepository(dbConnection))))
	routes.SetupFotoRoutes(server, controllers.NewFotoController(usecases.NewFotoUseCases(repository.NewFotoRepository(dbConnection))))
	routes.SetupImovelRoutes(server, controllers.NewImovelController(usecases.NewImovelUseCases(repository.NewImovelRepository(dbConnection))))
	routes.SetupConjugeRoutes(server, controllers.NewConjugeController(usecases.NewConjugeUseCases(repository.NewConjugeRepository(dbConnection))))
	routes.SetupLancamentoRoutes(server, controllers.NewLancamentoController(usecases.NewLancamentoUseCases(repository.NewLancamentoRepository(dbConnection))))
	routes.SetupInteresseRoutes(server, controllers.NewInteresseController(usecases.NewInteresseUseCases(repository.NewInteresseRepository(dbConnection))))

	port := server.Run(":3034")
	fmt.Println("Servidor rodando na porta: ", port)
}
