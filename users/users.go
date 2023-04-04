package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Users map[int]User = make(map[int]User)

func GetUserById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	if id, _ := strconv.ParseInt(id, 10, 64); id > 0 {
		id := int(id)
		if user, ok := Users[id]; ok {
			c.JSON(http.StatusOK, gin.H{
				"userInfo": user,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"userInfo": fmt.Sprint("Error: userId ", id, " is not exist"),
			})
		}
	}
}

func PostUserById(c *gin.Context) {
	body := c.Request.Body
	b, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"userInfo": "Wrong approach",
		})
	}

	var u User
	json.Unmarshal([]byte(b), &u)
	if _, ok := Users[u.Id]; !ok {
		Users[u.Id] = u
		c.JSON(http.StatusCreated, gin.H{
			"userInfo": Users[u.Id],
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"userInfo": "userId is exist",
		})
	}
}

func PutUserById(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"userInfo": "Wrong approach",
		})
	}

	var u User
	json.Unmarshal([]byte(body), &u)
	if user, ok := Users[u.Id]; ok {
		if user == u {
			c.JSON(http.StatusNoContent, gin.H{
				"userInfo": "Content is not changed",
			})
		} else if user.Password == u.Password {
			if u.ChangedPassword != "" {
				u.Password = u.ChangedPassword
				u.ChangedPassword = ""
			}
			Users[u.Id] = u
			c.JSON(http.StatusCreated, gin.H{
				"userInfo": Users[u.Id],
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"userInfo": "Wrong approach",
			})
		}
	}
}

func DeleteUserById(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"userInfo": "Wrong approach",
		})
	}

	var u User
	json.Unmarshal([]byte(body), &u)

	if user, ok := Users[u.Id]; ok && user.Password == u.Password {
		delete(Users, u.Id)
		c.JSON(http.StatusNoContent, gin.H{
			"userInfo": "user is deleted",
		})
	}
}
