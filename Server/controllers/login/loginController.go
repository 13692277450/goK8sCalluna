package loginControllers

import (
	"database/sql"
	"fmt"
	"gok8s/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// LoginController
type LoginController struct{}

// InitDB initialize the database connection

func InitDB() (*sql.DB, error) {
	// MySQL configuration
	//dbConfig := "root:root123@tcp(142.171.228.250:3306)/k8s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConfig := "root:86868686mM@tcp(192.168.1.213:3306)/k8s?charset=utf8mb4&parseTime=True&loc=Local"

	// open database connection
	db, err := sql.Open("mysql", dbConfig)
	if err != nil {
		return nil, fmt.Errorf("Database connection error: %v", err)
	}
	// test the database connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error is : ", err)
		return nil, fmt.Errorf("Failure to connect to database: %v", err)
	}

	return db, nil
}

// Login handles user login
func (lc *LoginController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// get username and password from form
	db, err := InitDB()
	// initialize database
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"Error": "Database connection error, please try again later",
		})
		return
	}
	defer db.Close()

	// query user by username
	var dbPassword string
	err = db.QueryRow("SELECT password FROM userAccount WHERE username = ?", username).Scan(&dbPassword)

	// process query result
	switch {
	case err == sql.ErrNoRows:
		// user not found
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error": "Username or password is incorrect",
		})
	case err != nil:
		// database query error
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"Error": "Database query error, please try again later",
		})
	default:
		// verify password
		if password == dbPassword {
			// login success
			//c.Redirect(http.StatusSeeOther, "index.html")
			pods := models.GetPods()
			resources := models.GetResources()
			pvcs := models.GetPVC()

			c.HTML(200, "twocloumns.html", gin.H{
				"Podlist":       pods,
				"PodName":       pods[0].Name,
				"PodPhase":      pods[0].Status,
				"PodIP":         pods[0].PodIP,
				"NodeName":      pods[0].NodeName,
				"HostIP":        pods[0].HostIP,
				"StartTime":     pods[0].StartTime,
				"Namespace":     pods[0].Namespace,
				"ResourcesList": resources,
				"PvcList":       pvcs,
				"PvcName":       pvcs[0].Name,
				"PvcNameSpace":  pvcs[0].Namespace,
				"PvcStatus":     pvcs[0].Status,
			})
		} else {
			// password != dbPassword
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"Error": "Username or password is incorrect",
			})
		}
	}
}

// ShowLoginPage renders the login page
func (lc *LoginController) ShowLoginPage(c *gin.Context) {
	//c.HTML(http.StatusOK, "../../website/loginantd.html", gin.H{})
	c.HTML(http.StatusOK, "login.html", gin.H{})

}
