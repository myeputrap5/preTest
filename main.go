package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Informasi struct {
	ID             string        `json:"id"`
	Language       string        `json:"language"`
	Appeared       int           `json:"appeared"`
	Created        []string      `json:"created"`
	Functional     bool          `json:"functional"`
	ObjectOriented bool          `json:"object-oriented"`
	Relation       *RelationData `json:"relation"`
}

type RelationData struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}

var relationDatas = RelationData{InfluencedBy: influencedBy, Influences: influences}
var influencedBy = []string{"B", "ALGOL 68", "Assembly", "FORTRAN"}
var influences = []string{"C++", "Objective-C", "C#", "Java", "Javascript", "PHP", "GO"}
var createdValue = []string{"Dennis Ritchie"}

var informasis = []Informasi{
	{ID: "1", Language: "C", Appeared: 1972, Created: createdValue, Functional: true, ObjectOriented: false, Relation: &relationDatas},
}

func getInformasis(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, informasis)
}

func getHello(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello Go developers"})
}

func isPalindromeUtil(inputan string) bool {
	balikan := false
	parameter := ""
	lenInputan := len(inputan)
	for a := lenInputan; a > 0; a-- {
		parameter = parameter + string(inputan[a-1])
	}
	if inputan == parameter {
		balikan = true
	}
	return balikan
}

func isPalindrome(context *gin.Context) {
	param := context.Param("param")
	balikan := isPalindromeUtil(param)
	if balikan {
		context.IndentedJSON(http.StatusAccepted, gin.H{"message": "Palindrome"})
	} else {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bukan Palindrome"})
	}

}

func addLanguage(context *gin.Context) {
	var newInformasi Informasi

	if err := context.BindJSON(&newInformasi); err != nil {
		return
	}
	if newInformasi.ID == "" {
		nomorID := len(informasis)
		newInformasi.ID = strconv.Itoa(nomorID + 1)
	}
	informasis = append(informasis, newInformasi)

	context.IndentedJSON(http.StatusCreated, newInformasi)
}

func getInformasiByID(id string) (*Informasi, error) {
	for i, t := range informasis {
		if t.ID == id {
			return &informasis[i], nil
		}
	}

	return nil, errors.New("Not Found")
}

func deleteByID(id string) (*Informasi, error) {
	for i, t := range informasis {
		if t.ID == id {
			informasis = append(informasis[:i], informasis[i+1:]...)
			return nil, nil
		}
	}

	return nil, errors.New("Not Found")

}

func getInformasi(context *gin.Context) {
	id := context.Param("id")
	info, err := getInformasiByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Language Not Found"})
	}

	context.IndentedJSON(http.StatusOK, info)
}

func deleteInformasi(context *gin.Context) {
	id := context.Param("id")
	_, err := deleteByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Language Not Found"})
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Language Deleted"})
}

func updateInformasi(context *gin.Context) {
	id := context.Param("id")
	info, err := getInformasiByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Language Not Found"})
	}

	var newInformasiPatch Informasi

	if err := context.BindJSON(&newInformasiPatch); err != nil {
		return
	}
	index, _ := strconv.Atoi(info.ID)
	if newInformasiPatch.ID == "" {
		newInformasiPatch.ID = strconv.Itoa(index)
	}
	informasis[index-1] = newInformasiPatch

	context.IndentedJSON(http.StatusOK, newInformasiPatch)
}

func main() {
	fmt.Println("hello world")
	router := gin.Default()
	router.GET("/language", getInformasis)
	router.GET("/language/:id", getInformasi)
	router.POST("/language", addLanguage)
	router.PATCH("/language/:id", updateInformasi)
	router.DELETE("/language/:id", deleteInformasi)
	router.GET("/palindrome/:param", isPalindrome)
	router.GET("", getHello)
	router.Run("localhost:8080")
}
