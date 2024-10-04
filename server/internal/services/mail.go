package services

import (
	// "bytes"
	// "fmt"
	// "mime/multipart"
	// "net/smtp"
	// "net/textproto"
	"gopkg.in/gomail.v2"
	
)

// func SendEmailWithAttachment(subject, to, from, smtpServer, smtpPort, username, password string) error {

// 	body := `
// <html>
// <head></head>
// <body>
// <p>Hello there,</p>
// <p>Here your is final inspection report.  </p>
// </p>https://drive.google.com/file/d/1YztrIGLJlx-LzV1o0vzYED2uMxKIbSCI/view?usp=sharing</p>
// </p>Thank you!</p>
// </body>
// </html>
// `

// 	var buf bytes.Buffer
// 	writer := multipart.NewWriter(&buf)

// 	header := make(map[string][]string)
// 	header["From"] = []string{from}
// 	header["To"] = []string{to}
// 	header["Subject"] = []string{subject}
// 	header["Content-Type"] = []string{"multipart/mixed; boundary=" + writer.Boundary()}

// 	for key, values := range header {
// 		for _, value := range values {
// 			buf.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
// 		}
// 	}
// 	buf.WriteString("\r\n")

// 	htmlPart, _ := writer.CreatePart(textproto.MIMEHeader{
// 		"Content-Type": {"text/html; charset=UTF-8"},
// 	})
// 	_, err := htmlPart.Write([]byte(body))
// 	if err != nil {
// 		return fmt.Errorf("failed to write HTML part: %v", err)
// 	}

// 	// Create the PDF attachment part
// 	// attachmentPart, _ := writer.CreatePart(textproto.MIMEHeader{
// 	//     "Content-Type":              {"application/pdf"},
// 	//     "Content-Disposition":       {"attachment; filename=\"" + filepath.Base(pdfFilePath) + "\""},
// 	//     "Content-Transfer-Encoding": {"base64"},
// 	// })
// 	// _, err = attachmentPart.Write(pdfData)
// 	// if err != nil {
// 	//     return fmt.Errorf("failed to write PDF part: %v", err)
// 	// }

// 	writer.Close()

// 	auth := smtp.PlainAuth("", username, password, smtpServer)

// 	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, buf.Bytes())
// 	if err != nil {
// 		return fmt.Errorf("failed to send email: %v", err)
// 	}

// 	return nil
// }
func SendEmail(to string, subject string, body string,  smtpServer string, smtpPort int, smtpUser string, smtpPassword string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Setup the SMTP dialer
	 // use "587" for most SMTP servers
	d := gomail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}