package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendEmail(c *gin.Context) {
	apiKey := "tu-clave-api" // Reemplaza con tu clave de API de SendGrid
	from := "go.study.bo@gmail.com"
	to := "leohermoso18@gmail.com"
	subject := "Prueba de correo"
	body := "Prueba de correo desde Go"

	// Crear el cuerpo del correo electrónico en formato JSON
	data := map[string]interface{}{
		"personalizations": []map[string]interface{}{
			{
				"to": []map[string]interface{}{
					{
						"email": to,
					},
				},
			},
		},
		"from": map[string]interface{}{
			"email": from,
		},
		"subject": subject,
		"content": []map[string]interface{}{
			{
				"type":  "text/plain",
				"value": body,
			},
		},
	}

	jsonData, _ := json.Marshal(data)

	// Realizar la solicitud POST a la API de SendGrid
	url := "https://api.sendgrid.com/v3/mail/send"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(400, gin.H{"error": "failed to send email"})
		return
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "failed to send email"})
		return
	}
	defer resp.Body.Close()

	c.JSON(200, gin.H{"message": "email sent successfully"})
}
