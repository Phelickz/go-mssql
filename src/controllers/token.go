package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mailgun/mailgun-go/v4"
)

func SaveTokenInDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		// channel := make(chan string)

		//getting query from request
		pin, _ := c.GetQuery("pin")
		email, _ := c.GetQuery("email")
		phoneNumber, _ := c.GetQuery("phoneNumber")
		token, _ := c.GetQuery("token")

		//creating context
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		//saving token in database
		_, err := dbInstance.QueryContext(ctx, "generateToken", sql.Named("pin", pin), sql.Named("email", email), sql.Named("phoneNumber", phoneNumber), sql.Named("token", token))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
			return
		}

		//sending sms with another go routine
		//I should make a channel to listen to the response from the go routine
		go sendSmsToken(phoneNumber, token, c)

		//sending email and waiting for the process to complete
		tokenErr := sendEmailToken(email, token, c)

		if tokenErr != nil {
			fmt.Println(tokenErr)
		}

		c.JSON(http.StatusOK, gin.H{"Message": "Sms token has been sent"})

	}
}

var sAPI string = "http://ngn.rmlconnect.net:8080/bulksms/bulksms"

func sendSmsToken(phoneNumber string, token string, c *gin.Context) error {

	// sending sms to phoneNumber

	//init client
	client := &http.Client{}

	// getting request object
	req, err := http.NewRequest(http.MethodGet, sAPI, nil)

	// message to send to user
	message := "Your OTP validation code is " + token

	if err != nil {
		log.Fatal(err)
		return err
	}

	//adding query
	q := req.URL.Query()
	q.Add("username", "oakpensions")
	q.Add("password", "sbnXHwMt")
	q.Add("type", "0&dlr=1")
	q.Add("destination", phoneNumber)
	q.Add("source", "OAKPENSIONS")
	q.Add("message", message)

	req.URL.RawQuery = q.Encode()

	//making request
	response, reqErr := client.Do(req)

	if reqErr != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"Error": "Token could not be sent"})
		fmt.Println(reqErr)
		return reqErr
	}

	defer response.Body.Close()

	//reading response
	responseBody, bodyErr := ioutil.ReadAll(response.Body)

	if bodyErr != nil {
		fmt.Println(bodyErr)
	}

	fmt.Println(string(responseBody))

	return nil

}

func sendEmailToken(email string, token string, c *gin.Context) error {

	//loading env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
	}

	//email content to send
	message := fmt.Sprintf("Dear Oak Pensions Client, Thank you for staying with us till date. Generated token to confirm your authentication is %v.\n\n\n\nThank you for choosing Oak Pensions Ltd.\nWe encourage you to log on to our website on www.oakpensions.com", token)

	//getting config files from env
	apiKey := os.Getenv("API_KEY")
	domain := os.Getenv("DOMAIN")

	//params for mailgun
	sender := "Oak Pensions Ltd <no-reply@oakpensions.com>"
	subject := "OAK PENSIONS"
	body := message
	recipient := email

	//instance of mailgun
	mg := mailgun.NewMailgun(domain, apiKey)

	sendEmail := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, sendEmail)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	c.JSON(http.StatusOK, gin.H{"message": resp})
	return nil

}
