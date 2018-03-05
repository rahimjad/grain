package price

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"../../db"
	"../../entities"
	"github.com/gin-gonic/gin"
)

type combinedData struct {
	Price     entities.Price
	Commodity entities.Commodity
	Parity    entities.Parity
	Location  entities.Location
}

//HandlePost creates a price and its associated relationships
func HandlePost(c *gin.Context) {
	r := c.Request
	fmt.Print("MAKING A POST\n")
	decoder := json.NewDecoder(r.Body)
	var cd combinedData
	err := decoder.Decode(&cd)
	if err != nil {
		fmt.Printf("%v", err)
	}

	pri := cd.Price
	pri.ID = generateUUID()
	com := cd.Commodity
	com.ID = generateUUID()
	loc := cd.Location
	loc.ID = generateUUID()
	par := cd.Parity
	par.ID = generateUUID()

	pri.CommodityID = com.ID
	com.ParityID = par.ID
	par.LocationID = loc.ID

	db.InsertModel(&pri)
	db.InsertModel(&com)
	db.InsertModel(&par)
	db.InsertModel(&loc)

	response, _ := json.Marshal(&cd)

	c.JSON(http.StatusOK, response)
}

func generateUUID() string {
	uuid, _ := exec.Command("uuidgen").Output()
	return fmt.Sprintf("%s", uuid)
}
