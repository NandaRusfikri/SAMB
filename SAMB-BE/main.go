package main

import (
	controller "SAMB-BE/controllers"
	"SAMB-BE/docs"
	"SAMB-BE/driver"
	"SAMB-BE/pkg"
	penerimaan_barangRepository "SAMB-BE/repositorys/barang_masuk"
	productRepository "SAMB-BE/repositorys/master_product"
	supplierRepository "SAMB-BE/repositorys/master_supplier"
	warehouseRepository "SAMB-BE/repositorys/master_warehouse"
	"SAMB-BE/schemas"
	penerimaan_barangService "SAMB-BE/services/barang_masuk"
	productService "SAMB-BE/services/master_product"
	supplierService "SAMB-BE/services/master_supplier"
	warehouseService "SAMB-BE/services/master_warehouse"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
)

var ConfigEnv schemas.SchemaEnvironment

func init() {
	//os.Setenv("TZ", os.Getenv("TIMEZONE"))
}

func main() {
	//loc, err := time.LoadLocation(os.Getenv("TIMEZONE"))
	//if err != nil {
	//	log.Fatal("main() - time - error: ", err)
	//}
	//time.Local = loc
	ConfigEnv = pkg.Environment(".env")
	RESTPort, err := strconv.Atoi(ConfigEnv.REST_PORT)
	if err != nil {
		log.Errorln("REST_PORT is not valid ", err.Error())
	}

	DbSQL, err := driver.SetupDBSQL(ConfigEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := SetupRouter(ConfigEnv)

	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "SAMB-BE"
	docs.SwaggerInfo.Description = "SAMB-BE"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ConfigEnv.SWAGGER_HOST
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// @title SAMB-BE
	// @version 1.0
	// @description This is a sample server celler server.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email nandarusfikri@gmail.com

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @query.collection.format multi

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization
	// @x-extension-openapi {"example": "value on a json format"}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	BarangMasukRepo := penerimaan_barangRepository.InitPenerimaanBarangRepository(DbSQL)
	BarangMasukServ := penerimaan_barangService.InitServicePenerimaanBarang(BarangMasukRepo)

	productRepo := productRepository.InitMasterProductRepository(DbSQL)
	productServ := productService.InitServiceMasterProduct(productRepo)

	supplierRepo := supplierRepository.InitMasterSupplierRepository(DbSQL)
	supplierServ := supplierService.InitServiceMasterSupplier(supplierRepo)

	warehouseRepo := warehouseRepository.InitMasterWareHouseRepository(DbSQL)
	warehouseServ := warehouseService.InitServiceMasterWareHouse(warehouseRepo)

	controller.InitControllerPenerimaanBarang(app, BarangMasukServ)
	controller.InitControllerMasterProduct(app, productServ)
	controller.InitControllerMasterSupplier(app, supplierServ)
	controller.InitControllerMasterWareHouse(app, warehouseServ)
	controller.InitControllerDefault(app)

	app.Run(fmt.Sprintf(":%v", RESTPort))

}

func SetupRouter(config schemas.SchemaEnvironment) *gin.Engine {

	app := gin.Default()

	if config.GO_ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.GO_ENV == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)

	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	//app.Use(helmet.Default())
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	return app
}
