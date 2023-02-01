package main

import (
	"context"
	"example/myco-api/ent"
	"example/myco-api/ent/grainjar"
	"example/myco-api/ent/sporesyringe"
	"log"
	"strconv"

	"example/myco-api/keys"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	db   *ent.Client
	http *gin.Engine
}

var svr Server

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, 400
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, 500
	}
	if !check {
		return http.StatusBadRequest, 400
	}

	return http.StatusOK, 200
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context, httpCode, errCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func initDatabase() {
	keys := keys.GetKeys()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", keys.User, keys.Password, keys.Host, keys.Port, keys.Database)
	client, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	//defer client.Close()

	svr.db = client

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func runHttpServer() {
	router := gin.Default()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	svr.http = router

	router.GET("/api/grainjars", getGrainJars)
	router.GET("/api/grainjar/:id", getGrainJar)
	router.POST("/api/grainjar", postGrainJar)

	router.POST("/api/sporesyringe", postSporeSyringe)
	_ = router.Run("localhost:8080")
}

func main() {
	initDatabase()
	runHttpServer()
}

func getGrainJars(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
	grainJar, err := svr.db.GrainJar.Query().All(context.Background())
	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "Get grain jars failed: "+err.Error(), nil)
		return
	}
	c.IndentedJSON(http.StatusOK, grainJar)
}

func getGrainJar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	grainJar, err := svr.db.GrainJar.Query().
		Where(grainjar.ID(id)).
		First(context.Background())
	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "Get grain jars failed: "+err.Error(), nil)
		return
	}
	c.IndentedJSON(http.StatusOK, grainJar)
}

func postGrainJar(c *gin.Context) {
	type PostParam struct {
		Grain        string `form:"grain" json:"grain" valid:"Required; MaxSize(50)"`
		SporeSyringe string `form:"spore_syringe" json:"spore_syringe" valid:"Required"`
	}
	var form PostParam

	httpCode, errCode := BindAndValid(c, &form)
	if errCode != 200 {
		ResponseJSON(c, httpCode, errCode, "invalid param", nil)
		return
	}

	id, _ := strconv.Atoi(form.SporeSyringe)
	sporeSyringe, err := svr.db.SporeSyringe.Query().
		Where(sporesyringe.ID(id)).
		First(context.Background())
	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "Couldn't find spore syringe: "+err.Error(), nil)
		return
	}

	grainJar, err := svr.db.GrainJar.
		Create().
		SetGrain(form.Grain).
		Save(context.Background())
	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "create grain jar failed: "+err.Error(), nil)
		return
	}

	_, newerr := svr.db.SporeSyringe.
		Update().
		AddGrainJar(grainJar).
		Where(sporesyringe.ID(sporeSyringe.ID)).
		Save(context.Background())
	if newerr != nil {
		ResponseJSON(c, http.StatusOK, 500, "create grain jar failed: "+newerr.Error(), nil)
		return
	}

	type ResponseData struct {
		ID               uint64 `json:"id"`
		Grain            string `json:"grain"`
		InnoculationDate string `json:"innoculation_date"`
	}
	var resp ResponseData
	resp.Grain = grainJar.Grain
	resp.ID = uint64(grainJar.ID)
	resp.InnoculationDate = grainJar.InnoculationDate.GoString()

	ResponseJSON(c, http.StatusOK, 200, "", resp)
}

func postSporeSyringe(c *gin.Context) {
	type PostParam struct {
		Species  string `form:"species" json:"species" valid:"Required; MaxSize(100)"`
		Supplier string `form:"supplier" json:"supplier" valid:"Required; MaxSize(100)"`
	}
	var form PostParam

	httpCode, errCode := BindAndValid(c, &form)
	if errCode != 200 {
		ResponseJSON(c, httpCode, errCode, "invalid param", nil)
		return
	}

	sporeSyringe, err := svr.db.SporeSyringe.
		Create().
		SetSpecies(form.Species).
		SetSupplier(form.Supplier).
		Save(context.Background())
	if err != nil {
		ResponseJSON(c, http.StatusOK, 500, "create spore syringe failed: "+err.Error(), nil)
		return
	}

	type ResponseData struct {
		ID       uint64 `json:"id"`
		Species  string `json:"species"`
		Supplier string `json:"supplier"`
	}
	var resp ResponseData
	resp.Species = sporeSyringe.Species
	resp.ID = uint64(sporeSyringe.ID)
	resp.Supplier = sporeSyringe.Supplier

	ResponseJSON(c, http.StatusOK, 200, "", resp)
}
